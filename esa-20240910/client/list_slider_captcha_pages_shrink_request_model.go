// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSliderCaptchaPagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSliderCaptchaPagesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSliderCaptchaPagesShrinkRequest
	GetPageSize() *int32
	SetQueryArgsShrink(v string) *ListSliderCaptchaPagesShrinkRequest
	GetQueryArgsShrink() *string
}

type ListSliderCaptchaPagesShrinkRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The filter parameters.
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListSliderCaptchaPagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSliderCaptchaPagesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSliderCaptchaPagesShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListSliderCaptchaPagesShrinkRequest) SetPageNumber(v int32) *ListSliderCaptchaPagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSliderCaptchaPagesShrinkRequest) SetPageSize(v int32) *ListSliderCaptchaPagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSliderCaptchaPagesShrinkRequest) SetQueryArgsShrink(v string) *ListSliderCaptchaPagesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListSliderCaptchaPagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
