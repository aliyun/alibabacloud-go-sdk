// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyUserResponseBody
	GetCode() *string
	SetData(v string) *ModifyUserResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyUserResponseBody
	GetMessage() *string
	SetParams(v []*string) *ModifyUserResponseBody
	GetParams() []*string
	SetRequestId(v string) *ModifyUserResponseBody
	GetRequestId() *string
}

type ModifyUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyUserResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyUserResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ModifyUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserResponseBody) SetCode(v string) *ModifyUserResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserResponseBody) SetData(v string) *ModifyUserResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyUserResponseBody) SetHttpStatusCode(v int32) *ModifyUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserResponseBody) SetMessage(v string) *ModifyUserResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserResponseBody) SetParams(v []*string) *ModifyUserResponseBody {
	s.Params = v
	return s
}

func (s *ModifyUserResponseBody) SetRequestId(v string) *ModifyUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserResponseBody) Validate() error {
	return dara.Validate(s)
}
