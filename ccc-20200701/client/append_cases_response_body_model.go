// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AppendCasesResponseBody
	GetCode() *string
	SetData(v []*AppendCasesResponseBodyData) *AppendCasesResponseBody
	GetData() []*AppendCasesResponseBodyData
	SetHttpStatusCode(v string) *AppendCasesResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *AppendCasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppendCasesResponseBody
	GetRequestId() *string
}

type AppendCasesResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*AppendCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AppendCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AppendCasesResponseBody) GetData() []*AppendCasesResponseBodyData {
	return s.Data
}

func (s *AppendCasesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *AppendCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppendCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppendCasesResponseBody) SetCode(v string) *AppendCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AppendCasesResponseBody) SetData(v []*AppendCasesResponseBodyData) *AppendCasesResponseBody {
	s.Data = v
	return s
}

func (s *AppendCasesResponseBody) SetHttpStatusCode(v string) *AppendCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AppendCasesResponseBody) SetMessage(v string) *AppendCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AppendCasesResponseBody) SetRequestId(v string) *AppendCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendCasesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppendCasesResponseBodyData struct {
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	PhoneNumber     *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ReferenceId     *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s AppendCasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppendCasesResponseBodyData) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *AppendCasesResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AppendCasesResponseBodyData) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *AppendCasesResponseBodyData) SetCustomVariables(v string) *AppendCasesResponseBodyData {
	s.CustomVariables = &v
	return s
}

func (s *AppendCasesResponseBodyData) SetPhoneNumber(v string) *AppendCasesResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *AppendCasesResponseBodyData) SetReferenceId(v string) *AppendCasesResponseBodyData {
	s.ReferenceId = &v
	return s
}

func (s *AppendCasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
