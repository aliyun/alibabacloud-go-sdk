// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedRewrite interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UnifiedRewrite
	GetEnabled() *bool
	SetTimeRange(v string) *UnifiedRewrite
	GetTimeRange() *string
}

type UnifiedRewrite struct {
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedRewrite) String() string {
	return dara.Prettify(s)
}

func (s UnifiedRewrite) GoString() string {
	return s.String()
}

func (s *UnifiedRewrite) GetEnabled() *bool {
	return s.Enabled
}

func (s *UnifiedRewrite) GetTimeRange() *string {
	return s.TimeRange
}

func (s *UnifiedRewrite) SetEnabled(v bool) *UnifiedRewrite {
	s.Enabled = &v
	return s
}

func (s *UnifiedRewrite) SetTimeRange(v string) *UnifiedRewrite {
	s.TimeRange = &v
	return s
}

func (s *UnifiedRewrite) Validate() error {
	return dara.Validate(s)
}
