// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddToMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddToMetaCategoryResponseBody
	GetData() *bool
	SetErrorCode(v string) *AddToMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddToMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *AddToMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddToMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddToMetaCategoryResponseBody
	GetSuccess() *bool
}

type AddToMetaCategoryResponseBody struct {
	// Indicates whether the metatable was added to the specified category.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
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
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddToMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddToMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddToMetaCategoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddToMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddToMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddToMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddToMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddToMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddToMetaCategoryResponseBody) SetData(v bool) *AddToMetaCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) SetErrorCode(v string) *AddToMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) SetErrorMessage(v string) *AddToMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) SetHttpStatusCode(v int32) *AddToMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) SetRequestId(v string) *AddToMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) SetSuccess(v bool) *AddToMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddToMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
