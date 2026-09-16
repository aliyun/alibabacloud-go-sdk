// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeRecallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetKnowledgeRecallResponseBodyData) *GetKnowledgeRecallResponseBody
	GetData() *GetKnowledgeRecallResponseBodyData
	SetRequestId(v string) *GetKnowledgeRecallResponseBody
	GetRequestId() *string
}

type GetKnowledgeRecallResponseBody struct {
	// The returned data.
	Data *GetKnowledgeRecallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKnowledgeRecallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeRecallResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeRecallResponseBody) GetData() *GetKnowledgeRecallResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeRecallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeRecallResponseBody) SetData(v *GetKnowledgeRecallResponseBodyData) *GetKnowledgeRecallResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeRecallResponseBody) SetRequestId(v string) *GetKnowledgeRecallResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeRecallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeRecallResponseBodyData struct {
	// The total number of entries.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// recall 5 files
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The recall results.
	Results []map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The Tracing Analysis ID.
	//
	// example:
	//
	// qf_c41fc27697d3
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetKnowledgeRecallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeRecallResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeRecallResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetKnowledgeRecallResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeRecallResponseBodyData) GetResults() []map[string]interface{} {
	return s.Results
}

func (s *GetKnowledgeRecallResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *GetKnowledgeRecallResponseBodyData) SetCount(v int32) *GetKnowledgeRecallResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetKnowledgeRecallResponseBodyData) SetMessage(v string) *GetKnowledgeRecallResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetKnowledgeRecallResponseBodyData) SetResults(v []map[string]interface{}) *GetKnowledgeRecallResponseBodyData {
	s.Results = v
	return s
}

func (s *GetKnowledgeRecallResponseBodyData) SetTraceId(v string) *GetKnowledgeRecallResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *GetKnowledgeRecallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
