// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFromMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteFromMetaCategoryResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteFromMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFromMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteFromMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteFromMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFromMetaCategoryResponseBody
	GetSuccess() *bool
}

type DeleteFromMetaCategoryResponseBody struct {
	// Indicates whether the table was removed from the specified category.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
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

func (s DeleteFromMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFromMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFromMetaCategoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteFromMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFromMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFromMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFromMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFromMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFromMetaCategoryResponseBody) SetData(v bool) *DeleteFromMetaCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) SetErrorCode(v string) *DeleteFromMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) SetErrorMessage(v string) *DeleteFromMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) SetHttpStatusCode(v int32) *DeleteFromMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) SetRequestId(v string) *DeleteFromMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) SetSuccess(v bool) *DeleteFromMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFromMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
