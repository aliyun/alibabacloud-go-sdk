// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProbeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProbeTaskResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateProbeTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProbeTaskResponseBody
	GetRequestId() *string
}

type UpdateProbeTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProbeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProbeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProbeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProbeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProbeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProbeTaskResponseBody) SetCode(v string) *UpdateProbeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProbeTaskResponseBody) SetMessage(v string) *UpdateProbeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProbeTaskResponseBody) SetRequestId(v string) *UpdateProbeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProbeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
