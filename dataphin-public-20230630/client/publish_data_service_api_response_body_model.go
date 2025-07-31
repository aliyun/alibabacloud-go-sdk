// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishDataServiceApiResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishDataServiceApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishDataServiceApiResponseBody
	GetSuccess() *bool
}

type PublishDataServiceApiResponseBody struct {
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

func (s PublishDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDataServiceApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishDataServiceApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishDataServiceApiResponseBody) SetCode(v string) *PublishDataServiceApiResponseBody {
	s.Code = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetHttpStatusCode(v int32) *PublishDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetMessage(v string) *PublishDataServiceApiResponseBody {
	s.Message = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetRequestId(v string) *PublishDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetSuccess(v bool) *PublishDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
