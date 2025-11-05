// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSchemaPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GrantSchemaPermissionResponseBody
	GetData() *bool
	SetErrorCode(v string) *GrantSchemaPermissionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GrantSchemaPermissionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GrantSchemaPermissionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GrantSchemaPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GrantSchemaPermissionResponseBody
	GetSuccess() *string
}

type GrantSchemaPermissionResponseBody struct {
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

func (s GrantSchemaPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantSchemaPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantSchemaPermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *GrantSchemaPermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GrantSchemaPermissionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GrantSchemaPermissionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GrantSchemaPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantSchemaPermissionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GrantSchemaPermissionResponseBody) SetData(v bool) *GrantSchemaPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) SetErrorCode(v string) *GrantSchemaPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) SetErrorMessage(v string) *GrantSchemaPermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) SetHttpStatusCode(v string) *GrantSchemaPermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) SetRequestId(v string) *GrantSchemaPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) SetSuccess(v string) *GrantSchemaPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantSchemaPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
