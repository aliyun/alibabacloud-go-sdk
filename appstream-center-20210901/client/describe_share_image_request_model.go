// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShareImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeShareImageRequest
	GetImageId() *string
}

type DescribeShareImageRequest struct {
	// The image ID.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeShareImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeShareImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeShareImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeShareImageRequest) SetImageId(v string) *DescribeShareImageRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeShareImageRequest) Validate() error {
	return dara.Validate(s)
}
