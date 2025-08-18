// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *DeleteAppInstancesRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceIds(v []*string) *DeleteAppInstancesRequest
	GetAppInstanceIds() []*string
	SetProductType(v string) *DeleteAppInstancesRequest
	GetProductType() *string
}

type DeleteAppInstancesRequest struct {
	// The ID of the delivery group. You can call the [listAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The IDs of application instances.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	AppInstanceIds []*string `json:"AppInstanceIds,omitempty" xml:"AppInstanceIds,omitempty" type:"Repeated"`
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

func (s DeleteAppInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DeleteAppInstancesRequest) GetAppInstanceIds() []*string {
	return s.AppInstanceIds
}

func (s *DeleteAppInstancesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DeleteAppInstancesRequest) SetAppInstanceGroupId(v string) *DeleteAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstancesRequest) SetAppInstanceIds(v []*string) *DeleteAppInstancesRequest {
	s.AppInstanceIds = v
	return s
}

func (s *DeleteAppInstancesRequest) SetProductType(v string) *DeleteAppInstancesRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteAppInstancesRequest) Validate() error {
	return dara.Validate(s)
}
