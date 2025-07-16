// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetBatchTranslateResponseBody
	GetCode() *int32
	SetMessage(v string) *GetBatchTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTranslateResponseBody
	GetRequestId() *string
	SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateResponseBody
	GetTranslatedList() []map[string]interface{}
}

type GetBatchTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranslatedList []map[string]interface{} `json:"TranslatedList,omitempty" xml:"TranslatedList,omitempty" type:"Repeated"`
}

func (s GetBatchTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetBatchTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTranslateResponseBody) GetTranslatedList() []map[string]interface{} {
	return s.TranslatedList
}

func (s *GetBatchTranslateResponseBody) SetCode(v int32) *GetBatchTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetMessage(v string) *GetBatchTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetRequestId(v string) *GetBatchTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateResponseBody {
	s.TranslatedList = v
	return s
}

func (s *GetBatchTranslateResponseBody) Validate() error {
	return dara.Validate(s)
}
