// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *GetAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *GetAppInstanceGroupRequest
	GetProductType() *string
}

type GetAppInstanceGroupRequest struct {
	// The ID of the delivery group. You can call the [listAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
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

func (s GetAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *GetAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupRequest) SetProductType(v string) *GetAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *GetAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
