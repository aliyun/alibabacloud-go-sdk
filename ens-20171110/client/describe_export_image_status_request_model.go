// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeExportImageStatusRequest
	GetImageId() *string
}

type DescribeExportImageStatusRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeExportImageStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeExportImageStatusRequest) SetImageId(v string) *DescribeExportImageStatusRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageStatusRequest) Validate() error {
	return dara.Validate(s)
}
