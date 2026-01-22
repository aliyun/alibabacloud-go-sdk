// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentChatRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePolarAgentChatRecordsResponseBodyData) *DescribePolarAgentChatRecordsResponseBody
	GetData() []*DescribePolarAgentChatRecordsResponseBodyData
	SetRequestId(v string) *DescribePolarAgentChatRecordsResponseBody
	GetRequestId() *string
}

type DescribePolarAgentChatRecordsResponseBody struct {
	Data []*DescribePolarAgentChatRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarAgentChatRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentChatRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentChatRecordsResponseBody) GetData() []*DescribePolarAgentChatRecordsResponseBodyData {
	return s.Data
}

func (s *DescribePolarAgentChatRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarAgentChatRecordsResponseBody) SetData(v []*DescribePolarAgentChatRecordsResponseBodyData) *DescribePolarAgentChatRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBody) SetRequestId(v string) *DescribePolarAgentChatRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBody) Validate() error {
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

type DescribePolarAgentChatRecordsResponseBodyData struct {
	// example:
	//
	// xxx
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 0
	FeedbackType *string `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	// example:
	//
	// xxx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Query ID。
	//
	// example:
	//
	// sq202506261002hz8b24fe80067683
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// 44dcdf31-04cd-4a44-9bae-834dd6657e29
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribePolarAgentChatRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentChatRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) GetFeedbackType() *string {
	return s.FeedbackType
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) GetQuery() *string {
	return s.Query
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) SetAnswer(v string) *DescribePolarAgentChatRecordsResponseBodyData {
	s.Answer = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) SetFeedbackType(v string) *DescribePolarAgentChatRecordsResponseBodyData {
	s.FeedbackType = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) SetQuery(v string) *DescribePolarAgentChatRecordsResponseBodyData {
	s.Query = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) SetQueryId(v string) *DescribePolarAgentChatRecordsResponseBodyData {
	s.QueryId = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) SetSessionId(v string) *DescribePolarAgentChatRecordsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
