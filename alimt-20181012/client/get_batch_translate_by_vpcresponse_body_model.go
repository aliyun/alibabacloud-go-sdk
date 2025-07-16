// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateByVPCResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetBatchTranslateByVPCResponseBody
	GetCode() *int32
	SetMessage(v string) *GetBatchTranslateByVPCResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTranslateByVPCResponseBody
	GetRequestId() *string
	SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateByVPCResponseBody
	GetTranslatedList() []map[string]interface{}
}

type GetBatchTranslateByVPCResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Param checker failed! param:[sourceText]. reason: stringChecker not pass. Param length not less [0] and not greater[10000]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranslatedList []map[string]interface{} `json:"TranslatedList,omitempty" xml:"TranslatedList,omitempty" type:"Repeated"`
}

func (s GetBatchTranslateByVPCResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateByVPCResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetBatchTranslateByVPCResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTranslateByVPCResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTranslateByVPCResponseBody) GetTranslatedList() []map[string]interface{} {
	return s.TranslatedList
}

func (s *GetBatchTranslateByVPCResponseBody) SetCode(v int32) *GetBatchTranslateByVPCResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetMessage(v string) *GetBatchTranslateByVPCResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetRequestId(v string) *GetBatchTranslateByVPCResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateByVPCResponseBody {
	s.TranslatedList = v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) Validate() error {
	return dara.Validate(s)
}
