// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTablePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GrantTablePermissionResponseBody
	GetData() *bool
	SetErrorCode(v string) *GrantTablePermissionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GrantTablePermissionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GrantTablePermissionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GrantTablePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GrantTablePermissionResponseBody
	GetSuccess() *string
}

type GrantTablePermissionResponseBody struct {
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

func (s GrantTablePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantTablePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantTablePermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *GrantTablePermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GrantTablePermissionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GrantTablePermissionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GrantTablePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantTablePermissionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GrantTablePermissionResponseBody) SetData(v bool) *GrantTablePermissionResponseBody {
	s.Data = &v
	return s
}

func (s *GrantTablePermissionResponseBody) SetErrorCode(v string) *GrantTablePermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantTablePermissionResponseBody) SetErrorMessage(v string) *GrantTablePermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantTablePermissionResponseBody) SetHttpStatusCode(v string) *GrantTablePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantTablePermissionResponseBody) SetRequestId(v string) *GrantTablePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantTablePermissionResponseBody) SetSuccess(v string) *GrantTablePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantTablePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
