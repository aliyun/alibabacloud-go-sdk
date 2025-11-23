// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteAuthorityTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteAuthorityTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteAuthorityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAuthorityTemplateResponseBody
	GetSuccess() *bool
	SetTid(v int64) *DeleteAuthorityTemplateResponseBody
	GetTid() *int64
}

type DeleteAuthorityTemplateResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteAuthorityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorityTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAuthorityTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAuthorityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuthorityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAuthorityTemplateResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteAuthorityTemplateResponseBody) SetErrorCode(v string) *DeleteAuthorityTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAuthorityTemplateResponseBody) SetErrorMessage(v string) *DeleteAuthorityTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAuthorityTemplateResponseBody) SetRequestId(v string) *DeleteAuthorityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuthorityTemplateResponseBody) SetSuccess(v bool) *DeleteAuthorityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAuthorityTemplateResponseBody) SetTid(v int64) *DeleteAuthorityTemplateResponseBody {
	s.Tid = &v
	return s
}

func (s *DeleteAuthorityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
