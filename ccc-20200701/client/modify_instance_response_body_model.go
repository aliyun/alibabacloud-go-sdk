// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyInstanceResponseBody
	GetCode() *string
	SetData(v string) *ModifyInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyInstanceResponseBody
	GetRequestId() *string
}

type ModifyInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 83TKE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceResponseBody) SetCode(v string) *ModifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetData(v string) *ModifyInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetMessage(v string) *ModifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
