// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDebugAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *GetDebugAppInstanceRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *GetDebugAppInstanceRequest
	GetProductType() *string
}

type GetDebugAppInstanceRequest struct {
	// The ID of the delivery group. You can call the `listAppInstanceGroup` operation to obtain the ID.
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

func (s GetDebugAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDebugAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetDebugAppInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetDebugAppInstanceRequest) SetAppInstanceGroupId(v string) *GetDebugAppInstanceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceRequest) SetProductType(v string) *GetDebugAppInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *GetDebugAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
