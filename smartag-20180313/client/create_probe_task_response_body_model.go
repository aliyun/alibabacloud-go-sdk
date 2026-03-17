// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProbeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProbeTaskResponseBody
	GetCode() *string
	SetData(v string) *CreateProbeTaskResponseBody
	GetData() *string
	SetMessage(v string) *CreateProbeTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProbeTaskResponseBody
	GetRequestId() *string
}

type CreateProbeTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the probe task.
	//
	// example:
	//
	// probe-****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateProbeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProbeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProbeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProbeTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateProbeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProbeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProbeTaskResponseBody) SetCode(v string) *CreateProbeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProbeTaskResponseBody) SetData(v string) *CreateProbeTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProbeTaskResponseBody) SetMessage(v string) *CreateProbeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProbeTaskResponseBody) SetRequestId(v string) *CreateProbeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProbeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
