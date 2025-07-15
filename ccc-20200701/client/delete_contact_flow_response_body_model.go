// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContactFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactFlowResponseBody
	GetRequestId() *string
}

type DeleteContactFlowResponseBody struct {
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
	// 07511949-6DC4-5D0B-8FA8-FF8FA29B4217
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactFlowResponseBody) SetCode(v string) *DeleteContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactFlowResponseBody) SetHttpStatusCode(v int32) *DeleteContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteContactFlowResponseBody) SetMessage(v string) *DeleteContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactFlowResponseBody) SetRequestId(v string) *DeleteContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
