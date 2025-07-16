// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestShrink(v string) *GetNewestInnerGroupsShrinkRequest
	GetRequestShrink() *string
}

type GetNewestInnerGroupsShrinkRequest struct {
	// example:
	//
	// {}
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s GetNewestInnerGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *GetNewestInnerGroupsShrinkRequest) SetRequestShrink(v string) *GetNewestInnerGroupsShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *GetNewestInnerGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
