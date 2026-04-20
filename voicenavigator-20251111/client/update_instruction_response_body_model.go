// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstructionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstructionResponseBody
	GetCode() *string
	SetData(v string) *UpdateInstructionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateInstructionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstructionResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateInstructionResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateInstructionResponseBody
	GetRequestId() *string
}

type UpdateInstructionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Transfer00
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstructionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstructionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstructionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstructionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateInstructionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstructionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstructionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateInstructionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstructionResponseBody) SetCode(v string) *UpdateInstructionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstructionResponseBody) SetData(v string) *UpdateInstructionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstructionResponseBody) SetHttpStatusCode(v int32) *UpdateInstructionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstructionResponseBody) SetMessage(v string) *UpdateInstructionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstructionResponseBody) SetParams(v []*string) *UpdateInstructionResponseBody {
	s.Params = v
	return s
}

func (s *UpdateInstructionResponseBody) SetRequestId(v string) *UpdateInstructionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstructionResponseBody) Validate() error {
	return dara.Validate(s)
}
