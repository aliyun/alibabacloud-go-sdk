// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpServerResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpServerResponseBody
	GetRequestId() *string
}

type UpdateMcpServerResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpServerResponseBody) SetCode(v string) *UpdateMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpServerResponseBody) SetMessage(v string) *UpdateMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpServerResponseBody) SetRequestId(v string) *UpdateMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
