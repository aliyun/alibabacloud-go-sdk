// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigImageRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ConfigImageRemarkRequest
	GetImageId() *string
	SetRemark(v string) *ConfigImageRemarkRequest
	GetRemark() *string
}

type ConfigImageRemarkRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// image-d79x4k11pmg19****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The description of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-****
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigImageRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigImageRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ConfigImageRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ConfigImageRemarkRequest) SetImageId(v string) *ConfigImageRemarkRequest {
	s.ImageId = &v
	return s
}

func (s *ConfigImageRemarkRequest) SetRemark(v string) *ConfigImageRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ConfigImageRemarkRequest) Validate() error {
	return dara.Validate(s)
}
