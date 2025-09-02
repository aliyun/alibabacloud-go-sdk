// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupGetOdpsRoleGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetData() []*string
	SetErrorCode(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUserGroupGetOdpsRoleGroupsResponseBody
	GetSuccess() *bool
}

type DsgUserGroupGetOdpsRoleGroupsResponseBody struct {
	// The data returned.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgUserGroupGetOdpsRoleGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupGetOdpsRoleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetData() []*string {
	return s.Data
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetData(v []*string) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetErrorCode(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetErrorMessage(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetHttpStatusCode(v int32) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetRequestId(v string) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) SetSuccess(v bool) *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
