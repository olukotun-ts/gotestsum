package main

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestNoSummaryValue_SetAndString(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		assert.Equal(t, newNoSummaryValue().String(), "none")
	})
	t.Run("one", func(t *testing.T) {
		value := newNoSummaryValue()
		assert.NilError(t, value.Set("output"))
		assert.Equal(t, value.String(), "output")
	})
	t.Run("some", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		value := newNoSummaryValue()
		assert.NilError(t, value.Set("errors,failed"))
		assert.Equal(t, value.String(), "failed,errors")
	})
	t.Run("bad value", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		value := newNoSummaryValue()
		assert.ErrorContains(t, value.Set("bogus"), "must be one or more of")
	})
}
