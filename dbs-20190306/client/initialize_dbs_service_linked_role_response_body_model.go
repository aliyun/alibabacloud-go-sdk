// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeDbsServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *InitializeDbsServiceLinkedRoleResponseBody
	GetData() *string
	SetErrMessage(v string) *InitializeDbsServiceLinkedRoleResponseBody
	GetErrMessage() *string
	SetErrorCode(v string) *InitializeDbsServiceLinkedRoleResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *InitializeDbsServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *InitializeDbsServiceLinkedRoleResponseBody
	GetSuccess() *string
}

type InitializeDbsServiceLinkedRoleResponseBody struct {
	// The value is null.
	//
	// example:
	//
	// null
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// EntityAlreadyExists.Role : The role already exists:AliyunServiceRoleForDBS\\r\\nRequestId : 73******-3B4D-501A-9505-FA8B9******
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// EntityAlreadyExists.Role
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitializeDbsServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeDbsServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) GetData() *string {
	return s.Data
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) SetData(v string) *InitializeDbsServiceLinkedRoleResponseBody {
	s.Data = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) SetErrMessage(v string) *InitializeDbsServiceLinkedRoleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) SetErrorCode(v string) *InitializeDbsServiceLinkedRoleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) SetRequestId(v string) *InitializeDbsServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) SetSuccess(v string) *InitializeDbsServiceLinkedRoleResponseBody {
	s.Success = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
