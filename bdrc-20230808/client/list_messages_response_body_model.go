// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody
	GetData() *ListMessagesResponseBodyData
	SetRequestId(v string) *ListMessagesResponseBody
	GetRequestId() *string
}

type ListMessagesResponseBody struct {
	// The data returned.
	Data *ListMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5748C531-80B1-5C31-8421-63A1830B9E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) GetData() *ListMessagesResponseBodyData {
	return s.Data
}

func (s *ListMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessagesResponseBody) SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessagesResponseBodyData struct {
	// Response parameters.
	Content []*ListMessagesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of results requested.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// eKDyCM0zFQ5op7jVMWmNNA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyData) GetContent() []*ListMessagesResponseBodyDataContent {
	return s.Content
}

func (s *ListMessagesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMessagesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMessagesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMessagesResponseBodyData) SetContent(v []*ListMessagesResponseBodyDataContent) *ListMessagesResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMessagesResponseBodyData) SetMaxResults(v int32) *ListMessagesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetNextToken(v string) *ListMessagesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTotalCount(v int64) *ListMessagesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMessagesResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessagesResponseBodyDataContent struct {
	// Message ID.
	//
	// example:
	//
	// zgrjap8j-us04-owef-fpmo-kdpr80pbss0k
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Message level.
	//
	// example:
	//
	// WARNING
	MessageLevel *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
	// Message name.
	//
	// example:
	//
	// MyMessage
	MessageName *string `json:"MessageName,omitempty" xml:"MessageName,omitempty"`
	// Message source ID.
	//
	// example:
	//
	// p-123***7890
	MessageSourceId *string `json:"MessageSourceId,omitempty" xml:"MessageSourceId,omitempty"`
	// Message source region ID.
	//
	// example:
	//
	// cn-hangzhou
	MessageSourceRegionId *string `json:"MessageSourceRegionId,omitempty" xml:"MessageSourceRegionId,omitempty"`
	// Message source type.
	//
	// example:
	//
	// PROTECTION_POLICY
	MessageSourceType *string `json:"MessageSourceType,omitempty" xml:"MessageSourceType,omitempty"`
	// Message time.
	//
	// example:
	//
	// 1740019609
	MessageTime *int64 `json:"MessageTime,omitempty" xml:"MessageTime,omitempty"`
	// Message type.
	//
	// example:
	//
	// SUB_PROTECTION_POLICY_MODIFIED
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s ListMessagesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataContent) GetMessageId() *string {
	return s.MessageId
}

func (s *ListMessagesResponseBodyDataContent) GetMessageLevel() *string {
	return s.MessageLevel
}

func (s *ListMessagesResponseBodyDataContent) GetMessageName() *string {
	return s.MessageName
}

func (s *ListMessagesResponseBodyDataContent) GetMessageSourceId() *string {
	return s.MessageSourceId
}

func (s *ListMessagesResponseBodyDataContent) GetMessageSourceRegionId() *string {
	return s.MessageSourceRegionId
}

func (s *ListMessagesResponseBodyDataContent) GetMessageSourceType() *string {
	return s.MessageSourceType
}

func (s *ListMessagesResponseBodyDataContent) GetMessageTime() *int64 {
	return s.MessageTime
}

func (s *ListMessagesResponseBodyDataContent) GetMessageType() *string {
	return s.MessageType
}

func (s *ListMessagesResponseBodyDataContent) SetMessageId(v string) *ListMessagesResponseBodyDataContent {
	s.MessageId = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageLevel(v string) *ListMessagesResponseBodyDataContent {
	s.MessageLevel = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageName(v string) *ListMessagesResponseBodyDataContent {
	s.MessageName = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageSourceId(v string) *ListMessagesResponseBodyDataContent {
	s.MessageSourceId = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageSourceRegionId(v string) *ListMessagesResponseBodyDataContent {
	s.MessageSourceRegionId = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageSourceType(v string) *ListMessagesResponseBodyDataContent {
	s.MessageSourceType = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageTime(v int64) *ListMessagesResponseBodyDataContent {
	s.MessageTime = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) SetMessageType(v string) *ListMessagesResponseBodyDataContent {
	s.MessageType = &v
	return s
}

func (s *ListMessagesResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
