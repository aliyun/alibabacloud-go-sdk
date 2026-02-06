// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UntagResourcesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UntagResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UntagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UntagResourcesResponseBody
	GetSuccess() *bool
}

type UntagResourcesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UntagResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UntagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetHttpStatusCode(v int32) *UntagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
