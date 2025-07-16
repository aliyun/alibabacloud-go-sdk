// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestShrink(v string) *ListCalendarsShrinkRequest
	GetRequestShrink() *string
}

type ListCalendarsShrinkRequest struct {
	// example:
	//
	// {}
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s ListCalendarsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarsShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *ListCalendarsShrinkRequest) SetRequestShrink(v string) *ListCalendarsShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListCalendarsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
