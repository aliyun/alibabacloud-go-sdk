// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *DeleteAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *DeleteAppInstanceGroupRequest
	GetProductType() *string
}

type DeleteAppInstanceGroupRequest struct {
	// The ID of the delivery group.
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

func (s DeleteAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DeleteAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DeleteAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *DeleteAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstanceGroupRequest) SetProductType(v string) *DeleteAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
