// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetUserRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterGetUserRolesResponseBodyData) *ModelRouterGetUserRolesResponseBody
	GetData() *ModelRouterGetUserRolesResponseBodyData
	SetErrCode(v string) *ModelRouterGetUserRolesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetUserRolesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetUserRolesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetUserRolesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetUserRolesResponseBody
	GetSuccess() *bool
}

type ModelRouterGetUserRolesResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterGetUserRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterGetUserRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetUserRolesResponseBody) GetData() *ModelRouterGetUserRolesResponseBodyData {
	return s.Data
}

func (s *ModelRouterGetUserRolesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetUserRolesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetUserRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetUserRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetUserRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetUserRolesResponseBody) SetData(v *ModelRouterGetUserRolesResponseBodyData) *ModelRouterGetUserRolesResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) SetErrCode(v string) *ModelRouterGetUserRolesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) SetErrMessage(v string) *ModelRouterGetUserRolesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetUserRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) SetRequestId(v string) *ModelRouterGetUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) SetSuccess(v bool) *ModelRouterGetUserRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetUserRolesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterGetUserRolesResponseBodyData struct {
	// The list of department role assignments.
	//
	// example:
	//
	// []
	DepartmentRoles []*UserDepartmentDTO `json:"departmentRoles,omitempty" xml:"departmentRoles,omitempty" type:"Repeated"`
}

func (s ModelRouterGetUserRolesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetUserRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterGetUserRolesResponseBodyData) GetDepartmentRoles() []*UserDepartmentDTO {
	return s.DepartmentRoles
}

func (s *ModelRouterGetUserRolesResponseBodyData) SetDepartmentRoles(v []*UserDepartmentDTO) *ModelRouterGetUserRolesResponseBodyData {
	s.DepartmentRoles = v
	return s
}

func (s *ModelRouterGetUserRolesResponseBodyData) Validate() error {
	if s.DepartmentRoles != nil {
		for _, item := range s.DepartmentRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
