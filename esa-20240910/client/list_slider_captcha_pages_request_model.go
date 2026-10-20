// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSliderCaptchaPagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSliderCaptchaPagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSliderCaptchaPagesRequest
	GetPageSize() *int32
	SetQueryArgs(v *ListSliderCaptchaPagesRequestQueryArgs) *ListSliderCaptchaPagesRequest
	GetQueryArgs() *ListSliderCaptchaPagesRequestQueryArgs
}

type ListSliderCaptchaPagesRequest struct {
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
	QueryArgs *ListSliderCaptchaPagesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListSliderCaptchaPagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesRequest) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSliderCaptchaPagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSliderCaptchaPagesRequest) GetQueryArgs() *ListSliderCaptchaPagesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListSliderCaptchaPagesRequest) SetPageNumber(v int32) *ListSliderCaptchaPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSliderCaptchaPagesRequest) SetPageSize(v int32) *ListSliderCaptchaPagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSliderCaptchaPagesRequest) SetQueryArgs(v *ListSliderCaptchaPagesRequestQueryArgs) *ListSliderCaptchaPagesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListSliderCaptchaPagesRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSliderCaptchaPagesRequestQueryArgs struct {
	// The fuzzy search keyword for the name.
	//
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
}

func (s ListSliderCaptchaPagesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesRequestQueryArgs) GetNameLike() *string {
	return s.NameLike
}

func (s *ListSliderCaptchaPagesRequestQueryArgs) SetNameLike(v string) *ListSliderCaptchaPagesRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListSliderCaptchaPagesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
