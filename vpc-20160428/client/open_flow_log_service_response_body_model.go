// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlowLogServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenFlowLogServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenFlowLogServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenFlowLogServiceResponseBody
	GetRequestId() *string
}

type OpenFlowLogServiceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information returned after the flow log feature is enabled.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 28CF47AB-B6C0-5FA2-80C7-2B28826A92CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenFlowLogServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenFlowLogServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenFlowLogServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenFlowLogServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenFlowLogServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenFlowLogServiceResponseBody) SetCode(v string) *OpenFlowLogServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenFlowLogServiceResponseBody) SetMessage(v string) *OpenFlowLogServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenFlowLogServiceResponseBody) SetRequestId(v string) *OpenFlowLogServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenFlowLogServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
