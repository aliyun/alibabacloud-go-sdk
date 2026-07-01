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
	SetData(v string) *UntagResourcesResponseBody
	GetData() *string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
}

type UntagResourcesResponseBody struct {
	// The request status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error code list](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result of deleting the tag. Valid values:
	//
	// - **true**: success.
	//
	// - **false**: failure.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UntagResourcesResponseBody) GetData() *string {
	return s.Data
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v string) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
