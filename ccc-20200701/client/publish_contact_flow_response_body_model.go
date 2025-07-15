// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishContactFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishContactFlowResponseBody
	GetRequestId() *string
}

type PublishContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BFB6788F-20D4-5767-BC67-99EAAC28F789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishContactFlowResponseBody) SetCode(v string) *PublishContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *PublishContactFlowResponseBody) SetHttpStatusCode(v int32) *PublishContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishContactFlowResponseBody) SetMessage(v string) *PublishContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *PublishContactFlowResponseBody) SetRequestId(v string) *PublishContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
