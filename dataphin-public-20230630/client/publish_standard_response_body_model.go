// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishStandardResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishStandardResponseBody
	GetSuccess() *bool
}

type PublishStandardResponseBody struct {
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
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishStandardResponseBody) GoString() string {
	return s.String()
}

func (s *PublishStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishStandardResponseBody) SetCode(v string) *PublishStandardResponseBody {
	s.Code = &v
	return s
}

func (s *PublishStandardResponseBody) SetHttpStatusCode(v int32) *PublishStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishStandardResponseBody) SetMessage(v string) *PublishStandardResponseBody {
	s.Message = &v
	return s
}

func (s *PublishStandardResponseBody) SetRequestId(v string) *PublishStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishStandardResponseBody) SetSuccess(v bool) *PublishStandardResponseBody {
	s.Success = &v
	return s
}

func (s *PublishStandardResponseBody) Validate() error {
	return dara.Validate(s)
}
