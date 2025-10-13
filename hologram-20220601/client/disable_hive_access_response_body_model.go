// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHiveAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableHiveAccessResponseBody
	GetData() *bool
	SetErrorCode(v string) *DisableHiveAccessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DisableHiveAccessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DisableHiveAccessResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DisableHiveAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableHiveAccessResponseBody
	GetSuccess() *bool
}

type DisableHiveAccessResponseBody struct {
	// The returned result.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 82B7A554-4D00-50DF-95D9-B59E7B4D5489
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableHiveAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableHiveAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableHiveAccessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableHiveAccessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableHiveAccessResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DisableHiveAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableHiveAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableHiveAccessResponseBody) SetData(v bool) *DisableHiveAccessResponseBody {
	s.Data = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetErrorCode(v string) *DisableHiveAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetErrorMessage(v string) *DisableHiveAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetHttpStatusCode(v string) *DisableHiveAccessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetRequestId(v string) *DisableHiveAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetSuccess(v bool) *DisableHiveAccessResponseBody {
	s.Success = &v
	return s
}

func (s *DisableHiveAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
