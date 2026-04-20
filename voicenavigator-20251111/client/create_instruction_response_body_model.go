// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstructionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstructionResponseBody
	GetCode() *string
	SetData(v string) *CreateInstructionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateInstructionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstructionResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateInstructionResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateInstructionResponseBody
	GetRequestId() *string
}

type CreateInstructionResponseBody struct {
	// example:
	//
	// OK
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
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstructionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstructionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstructionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstructionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateInstructionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstructionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstructionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateInstructionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstructionResponseBody) SetCode(v string) *CreateInstructionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstructionResponseBody) SetData(v string) *CreateInstructionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstructionResponseBody) SetHttpStatusCode(v int32) *CreateInstructionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstructionResponseBody) SetMessage(v string) *CreateInstructionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstructionResponseBody) SetParams(v []*string) *CreateInstructionResponseBody {
	s.Params = v
	return s
}

func (s *CreateInstructionResponseBody) SetRequestId(v string) *CreateInstructionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstructionResponseBody) Validate() error {
	return dara.Validate(s)
}
