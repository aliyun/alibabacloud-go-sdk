// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMetaCategoryResponseBodyData) *CreateMetaCategoryResponseBody
	GetData() *CreateMetaCategoryResponseBodyData
	SetErrorCode(v string) *CreateMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetaCategoryResponseBody
	GetSuccess() *bool
}

type CreateMetaCategoryResponseBody struct {
	// The information about the category.
	Data *CreateMetaCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryResponseBody) GetData() *CreateMetaCategoryResponseBodyData {
	return s.Data
}

func (s *CreateMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetaCategoryResponseBody) SetData(v *CreateMetaCategoryResponseBodyData) *CreateMetaCategoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetErrorCode(v string) *CreateMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetErrorMessage(v string) *CreateMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetHttpStatusCode(v int32) *CreateMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetRequestId(v string) *CreateMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetSuccess(v bool) *CreateMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMetaCategoryResponseBodyData struct {
	// The category ID.
	//
	// example:
	//
	// 223
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s CreateMetaCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryResponseBodyData) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateMetaCategoryResponseBodyData) SetCategoryId(v int64) *CreateMetaCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *CreateMetaCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
