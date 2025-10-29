// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *CreateBusinessResponseBody
	GetBusinessId() *int64
	SetErrorCode(v string) *CreateBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBusinessResponseBody
	GetSuccess() *bool
}

type CreateBusinessResponseBody struct {
	// The workflow ID.
	//
	// example:
	//
	// 100001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The error code.
	//
	// 	- Request succeeded: The ErrorCode field is not returned.
	//
	// 	- Request failed: The ErrorCode field is returned.
	//
	// For more information, see the error code section.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique ID of this request. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
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

func (s CreateBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBusinessResponseBody) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *CreateBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBusinessResponseBody) SetBusinessId(v int64) *CreateBusinessResponseBody {
	s.BusinessId = &v
	return s
}

func (s *CreateBusinessResponseBody) SetErrorCode(v string) *CreateBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBusinessResponseBody) SetErrorMessage(v string) *CreateBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBusinessResponseBody) SetHttpStatusCode(v int32) *CreateBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBusinessResponseBody) SetRequestId(v string) *CreateBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBusinessResponseBody) SetSuccess(v bool) *CreateBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}
