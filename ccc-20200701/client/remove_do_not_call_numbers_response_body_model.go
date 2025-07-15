// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDoNotCallNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveDoNotCallNumbersResponseBody
	GetCode() *string
	SetData(v string) *RemoveDoNotCallNumbersResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *RemoveDoNotCallNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveDoNotCallNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *RemoveDoNotCallNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *RemoveDoNotCallNumbersResponseBody
	GetRequestId() *string
}

type RemoveDoNotCallNumbersResponseBody struct {
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
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDoNotCallNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDoNotCallNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDoNotCallNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveDoNotCallNumbersResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveDoNotCallNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveDoNotCallNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveDoNotCallNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RemoveDoNotCallNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDoNotCallNumbersResponseBody) SetCode(v string) *RemoveDoNotCallNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) SetData(v string) *RemoveDoNotCallNumbersResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) SetHttpStatusCode(v int32) *RemoveDoNotCallNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) SetMessage(v string) *RemoveDoNotCallNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) SetParams(v []*string) *RemoveDoNotCallNumbersResponseBody {
	s.Params = v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) SetRequestId(v string) *RemoveDoNotCallNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
