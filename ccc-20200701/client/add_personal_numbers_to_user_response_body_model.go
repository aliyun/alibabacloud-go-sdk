// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPersonalNumbersToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddPersonalNumbersToUserResponseBody
	GetCode() *string
	SetData(v []*string) *AddPersonalNumbersToUserResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *AddPersonalNumbersToUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddPersonalNumbersToUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPersonalNumbersToUserResponseBody
	GetRequestId() *string
}

type AddPersonalNumbersToUserResponseBody struct {
	// example:
	//
	// OK
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPersonalNumbersToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPersonalNumbersToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddPersonalNumbersToUserResponseBody) GetData() []*string {
	return s.Data
}

func (s *AddPersonalNumbersToUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddPersonalNumbersToUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPersonalNumbersToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPersonalNumbersToUserResponseBody) SetCode(v string) *AddPersonalNumbersToUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetData(v []*string) *AddPersonalNumbersToUserResponseBody {
	s.Data = v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetHttpStatusCode(v int32) *AddPersonalNumbersToUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetMessage(v string) *AddPersonalNumbersToUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetRequestId(v string) *AddPersonalNumbersToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) Validate() error {
	return dara.Validate(s)
}
