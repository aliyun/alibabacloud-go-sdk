// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteTemplatesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteTemplatesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTemplatesResponseBody
	GetSuccess() *bool
}

type DeleteTemplatesResponseBody struct {
	// The information about the request denial..
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The modification result. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the error code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned. This parameter is not returned if the value of Code is `success`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTemplatesResponseBody) SetAccessDeniedDetail(v string) *DeleteTemplatesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteTemplatesResponseBody) SetCode(v string) *DeleteTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplatesResponseBody) SetHttpStatusCode(v int32) *DeleteTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTemplatesResponseBody) SetMessage(v string) *DeleteTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTemplatesResponseBody) SetRequestId(v string) *DeleteTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplatesResponseBody) SetSuccess(v bool) *DeleteTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
