// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWindowLimit interface {
	dara.Model
	String() string
	GoString() string
	SetDurationSecs(v int64) *WindowLimit
	GetDurationSecs() *int64
	SetLimit(v int64) *WindowLimit
	GetLimit() *int64
}

type WindowLimit struct {
	// 限流时间窗口的持续时间，单位为秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	DurationSecs *int64 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// 在指定时间窗口内允许的最大请求数量
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s WindowLimit) String() string {
	return dara.Prettify(s)
}

func (s WindowLimit) GoString() string {
	return s.String()
}

func (s *WindowLimit) GetDurationSecs() *int64 {
	return s.DurationSecs
}

func (s *WindowLimit) GetLimit() *int64 {
	return s.Limit
}

func (s *WindowLimit) SetDurationSecs(v int64) *WindowLimit {
	s.DurationSecs = &v
	return s
}

func (s *WindowLimit) SetLimit(v int64) *WindowLimit {
	s.Limit = &v
	return s
}

func (s *WindowLimit) Validate() error {
	return dara.Validate(s)
}
