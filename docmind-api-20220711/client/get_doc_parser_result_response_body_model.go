// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParserResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocParserResultResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetDocParserResultResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetDocParserResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocParserResultResponseBody
	GetRequestId() *string
}

type GetDocParserResultResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDocParserResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocParserResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocParserResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocParserResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetDocParserResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocParserResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocParserResultResponseBody) SetCode(v string) *GetDocParserResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocParserResultResponseBody) SetData(v map[string]interface{}) *GetDocParserResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocParserResultResponseBody) SetMessage(v string) *GetDocParserResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocParserResultResponseBody) SetRequestId(v string) *GetDocParserResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocParserResultResponseBody) Validate() error {
	return dara.Validate(s)
}
