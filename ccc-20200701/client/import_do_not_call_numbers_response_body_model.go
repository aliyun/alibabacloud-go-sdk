// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDoNotCallNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportDoNotCallNumbersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportDoNotCallNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportDoNotCallNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ImportDoNotCallNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ImportDoNotCallNumbersResponseBody
	GetRequestId() *string
}

type ImportDoNotCallNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s ImportDoNotCallNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDoNotCallNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDoNotCallNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportDoNotCallNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportDoNotCallNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportDoNotCallNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ImportDoNotCallNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDoNotCallNumbersResponseBody) SetCode(v string) *ImportDoNotCallNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDoNotCallNumbersResponseBody) SetHttpStatusCode(v int32) *ImportDoNotCallNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportDoNotCallNumbersResponseBody) SetMessage(v string) *ImportDoNotCallNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDoNotCallNumbersResponseBody) SetParams(v []*string) *ImportDoNotCallNumbersResponseBody {
	s.Params = v
	return s
}

func (s *ImportDoNotCallNumbersResponseBody) SetRequestId(v string) *ImportDoNotCallNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDoNotCallNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
