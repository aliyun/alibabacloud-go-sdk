// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteMetaCategoryResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetaCategoryResponseBody
	GetSuccess() *bool
}

type DeleteMetaCategoryResponseBody struct {
	// The business data.
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

func (s DeleteMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaCategoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetaCategoryResponseBody) SetData(v bool) *DeleteMetaCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetErrorCode(v string) *DeleteMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetErrorMessage(v string) *DeleteMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetHttpStatusCode(v int32) *DeleteMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetRequestId(v string) *DeleteMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetSuccess(v bool) *DeleteMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
