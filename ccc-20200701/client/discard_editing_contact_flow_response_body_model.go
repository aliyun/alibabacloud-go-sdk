// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardEditingContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DiscardEditingContactFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DiscardEditingContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DiscardEditingContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DiscardEditingContactFlowResponseBody
	GetRequestId() *string
}

type DiscardEditingContactFlowResponseBody struct {
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
	// CB72B537-B531-598F-9617-A636FB8040C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DiscardEditingContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiscardEditingContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DiscardEditingContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DiscardEditingContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DiscardEditingContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DiscardEditingContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiscardEditingContactFlowResponseBody) SetCode(v string) *DiscardEditingContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DiscardEditingContactFlowResponseBody) SetHttpStatusCode(v int32) *DiscardEditingContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DiscardEditingContactFlowResponseBody) SetMessage(v string) *DiscardEditingContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DiscardEditingContactFlowResponseBody) SetRequestId(v string) *DiscardEditingContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiscardEditingContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
