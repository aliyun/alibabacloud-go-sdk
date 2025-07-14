// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationImageRequest
	GetAppId() *string
	SetImageUrl(v string) *DescribeApplicationImageRequest
	GetImageUrl() *string
}

type DescribeApplicationImageRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URL of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/demo/demo:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DescribeApplicationImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationImageRequest) SetAppId(v string) *DescribeApplicationImageRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationImageRequest) SetImageUrl(v string) *DescribeApplicationImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationImageRequest) Validate() error {
	return dara.Validate(s)
}
