// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceGroupImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageId(v string) *UpdateAppInstanceGroupImageRequest
	GetAppCenterImageId() *string
	SetAppInstanceGroupId(v string) *UpdateAppInstanceGroupImageRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *UpdateAppInstanceGroupImageRequest
	GetBizRegionId() *string
	SetProductType(v string) *UpdateAppInstanceGroupImageRequest
	GetProductType() *string
}

type UpdateAppInstanceGroupImageRequest struct {
	// The image ID of the application. You can obtain the ID from the Images page in the App Streaming console.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UpdateAppInstanceGroupImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageRequest) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *UpdateAppInstanceGroupImageRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *UpdateAppInstanceGroupImageRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *UpdateAppInstanceGroupImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppCenterImageId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppInstanceGroupId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetBizRegionId(v string) *UpdateAppInstanceGroupImageRequest {
	s.BizRegionId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetProductType(v string) *UpdateAppInstanceGroupImageRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) Validate() error {
	return dara.Validate(s)
}
