// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *UnbindRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *UnbindRequest
	GetAppInstanceId() *string
	SetAppInstancePersistentId(v string) *UnbindRequest
	GetAppInstancePersistentId() *string
	SetEndUserId(v string) *UnbindRequest
	GetEndUserId() *string
	SetProductType(v string) *UnbindRequest
	GetProductType() *string
}

type UnbindRequest struct {
	// The ID of the delivery group. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The session ID. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// example:
	//
	// ai-d297eyf83g5ni****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
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

func (s UnbindRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *UnbindRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *UnbindRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *UnbindRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *UnbindRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetAppInstancePersistentId(v string) *UnbindRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindRequest) Validate() error {
	return dara.Validate(s)
}
