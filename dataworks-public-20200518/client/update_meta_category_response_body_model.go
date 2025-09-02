// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateMetaCategoryResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMetaCategoryResponseBody
	GetSuccess() *bool
}

type UpdateMetaCategoryResponseBody struct {
	// Indicates whether the category is updated.
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

func (s UpdateMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaCategoryResponseBody) SetData(v bool) *UpdateMetaCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetErrorCode(v string) *UpdateMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetErrorMessage(v string) *UpdateMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetHttpStatusCode(v int32) *UpdateMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetRequestId(v string) *UpdateMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetSuccess(v bool) *UpdateMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
