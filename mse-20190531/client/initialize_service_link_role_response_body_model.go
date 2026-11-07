// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeServiceLinkRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitializeServiceLinkRoleResponseBody
	GetCode() *string
	SetData(v *InitializeServiceLinkRoleResponseBodyData) *InitializeServiceLinkRoleResponseBody
	GetData() *InitializeServiceLinkRoleResponseBodyData
	SetMessage(v string) *InitializeServiceLinkRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitializeServiceLinkRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitializeServiceLinkRoleResponseBody
	GetSuccess() *bool
}

type InitializeServiceLinkRoleResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *InitializeServiceLinkRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5E7D4E3-D30C-56C1-817F-F2B8CE6BXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitializeServiceLinkRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeServiceLinkRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitializeServiceLinkRoleResponseBody) GetData() *InitializeServiceLinkRoleResponseBodyData {
	return s.Data
}

func (s *InitializeServiceLinkRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitializeServiceLinkRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeServiceLinkRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitializeServiceLinkRoleResponseBody) SetCode(v string) *InitializeServiceLinkRoleResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBody) SetData(v *InitializeServiceLinkRoleResponseBodyData) *InitializeServiceLinkRoleResponseBody {
	s.Data = v
	return s
}

func (s *InitializeServiceLinkRoleResponseBody) SetMessage(v string) *InitializeServiceLinkRoleResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBody) SetRequestId(v string) *InitializeServiceLinkRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBody) SetSuccess(v bool) *InitializeServiceLinkRoleResponseBody {
	s.Success = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitializeServiceLinkRoleResponseBodyData struct {
	// The permission that is granted to the RAM user to call the API operation.
	//
	// example:
	//
	// ram:CreateServiceLinkedRole
	RequiredPermission *string `json:"RequiredPermission,omitempty" xml:"RequiredPermission,omitempty"`
	// The name of the created service-linked role.
	//
	// example:
	//
	// AliyunServiceRoleForMSE
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The service to which the service-linked role belongs.
	//
	// example:
	//
	// mse.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s InitializeServiceLinkRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitializeServiceLinkRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkRoleResponseBodyData) GetRequiredPermission() *string {
	return s.RequiredPermission
}

func (s *InitializeServiceLinkRoleResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *InitializeServiceLinkRoleResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *InitializeServiceLinkRoleResponseBodyData) SetRequiredPermission(v string) *InitializeServiceLinkRoleResponseBodyData {
	s.RequiredPermission = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBodyData) SetRoleName(v string) *InitializeServiceLinkRoleResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBodyData) SetServiceName(v string) *InitializeServiceLinkRoleResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *InitializeServiceLinkRoleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
