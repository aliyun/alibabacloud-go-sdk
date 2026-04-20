// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstructionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstructionResponseBody
	GetCode() *string
	SetData(v string) *DeleteInstructionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteInstructionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInstructionResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteInstructionResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteInstructionResponseBody
	GetRequestId() *string
}

type DeleteInstructionResponseBody struct {
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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstructionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstructionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstructionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstructionResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteInstructionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstructionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstructionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteInstructionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstructionResponseBody) SetCode(v string) *DeleteInstructionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstructionResponseBody) SetData(v string) *DeleteInstructionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstructionResponseBody) SetHttpStatusCode(v int32) *DeleteInstructionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstructionResponseBody) SetMessage(v string) *DeleteInstructionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstructionResponseBody) SetParams(v []*string) *DeleteInstructionResponseBody {
	s.Params = v
	return s
}

func (s *DeleteInstructionResponseBody) SetRequestId(v string) *DeleteInstructionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstructionResponseBody) Validate() error {
	return dara.Validate(s)
}
