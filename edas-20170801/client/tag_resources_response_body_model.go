// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TagResourcesResponseBody
	GetCode() *int32
	SetMessage(v string) *TagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
}

type TagResourcesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 000e5836-xxxx-xxxx-xxxx-0d6ab2ac4877
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetCode(v int32) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
