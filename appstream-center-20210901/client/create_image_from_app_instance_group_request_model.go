// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageFromAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageName(v string) *CreateImageFromAppInstanceGroupRequest
	GetAppCenterImageName() *string
	SetAppInstanceGroupId(v string) *CreateImageFromAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *CreateImageFromAppInstanceGroupRequest
	GetProductType() *string
}

type CreateImageFromAppInstanceGroupRequest struct {
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_name
	AppCenterImageName *string `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
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

func (s CreateImageFromAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupRequest) GetAppCenterImageName() *string {
	return s.AppCenterImageName
}

func (s *CreateImageFromAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *CreateImageFromAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppCenterImageName(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppCenterImageName = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetProductType(v string) *CreateImageFromAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
