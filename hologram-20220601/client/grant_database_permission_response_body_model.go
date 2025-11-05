// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDatabasePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GrantDatabasePermissionResponseBody
	GetData() *bool
	SetErrorCode(v string) *GrantDatabasePermissionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GrantDatabasePermissionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GrantDatabasePermissionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GrantDatabasePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GrantDatabasePermissionResponseBody
	GetSuccess() *string
}

type GrantDatabasePermissionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantDatabasePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantDatabasePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantDatabasePermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *GrantDatabasePermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GrantDatabasePermissionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GrantDatabasePermissionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GrantDatabasePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantDatabasePermissionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GrantDatabasePermissionResponseBody) SetData(v bool) *GrantDatabasePermissionResponseBody {
	s.Data = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) SetErrorCode(v string) *GrantDatabasePermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) SetErrorMessage(v string) *GrantDatabasePermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) SetHttpStatusCode(v string) *GrantDatabasePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) SetRequestId(v string) *GrantDatabasePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) SetSuccess(v string) *GrantDatabasePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantDatabasePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
