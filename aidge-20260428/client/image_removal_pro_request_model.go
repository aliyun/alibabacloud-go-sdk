// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *ImageRemovalProRequest
	GetAsync() *bool
	SetImageUrl(v string) *ImageRemovalProRequest
	GetImageUrl() *string
}

type ImageRemovalProRequest struct {
	// The call type. Valid values:
	//
	// - true: asynchronous.
	//
	// - false: synchronous.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The URL of the image to process.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://cbu01.alicdn.com/imgextra/i2/1067106875/O1CN013rvpXP20enxPQttb3_!!4611686018427380283-0-item_pic.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ImageRemovalProRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProRequest) GoString() string {
	return s.String()
}

func (s *ImageRemovalProRequest) GetAsync() *bool {
	return s.Async
}

func (s *ImageRemovalProRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemovalProRequest) SetAsync(v bool) *ImageRemovalProRequest {
	s.Async = &v
	return s
}

func (s *ImageRemovalProRequest) SetImageUrl(v string) *ImageRemovalProRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemovalProRequest) Validate() error {
	return dara.Validate(s)
}
