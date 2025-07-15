// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPrivacyCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartPrivacyCallResponseBody
	GetCode() *string
	SetData(v string) *StartPrivacyCallResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *StartPrivacyCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartPrivacyCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *StartPrivacyCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *StartPrivacyCallResponseBody
	GetRequestId() *string
}

type StartPrivacyCallResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// job-xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 0630E5DF-CEB0-445B-8626-D5C7481181C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPrivacyCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPrivacyCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartPrivacyCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartPrivacyCallResponseBody) GetData() *string {
	return s.Data
}

func (s *StartPrivacyCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartPrivacyCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartPrivacyCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *StartPrivacyCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPrivacyCallResponseBody) SetCode(v string) *StartPrivacyCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartPrivacyCallResponseBody) SetData(v string) *StartPrivacyCallResponseBody {
	s.Data = &v
	return s
}

func (s *StartPrivacyCallResponseBody) SetHttpStatusCode(v int32) *StartPrivacyCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPrivacyCallResponseBody) SetMessage(v string) *StartPrivacyCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartPrivacyCallResponseBody) SetParams(v []*string) *StartPrivacyCallResponseBody {
	s.Params = v
	return s
}

func (s *StartPrivacyCallResponseBody) SetRequestId(v string) *StartPrivacyCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPrivacyCallResponseBody) Validate() error {
	return dara.Validate(s)
}
