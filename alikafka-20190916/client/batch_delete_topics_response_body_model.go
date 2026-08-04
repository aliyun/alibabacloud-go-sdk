// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchDeleteTopicsResponseBody
	GetCode() *int32
	SetData(v *BatchDeleteTopicsResponseBodyData) *BatchDeleteTopicsResponseBody
	GetData() *BatchDeleteTopicsResponseBodyData
	SetMessage(v string) *BatchDeleteTopicsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteTopicsResponseBody
	GetSuccess() *bool
}

type BatchDeleteTopicsResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchDeleteTopicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchDeleteTopicsResponseBody) GetData() *BatchDeleteTopicsResponseBodyData {
	return s.Data
}

func (s *BatchDeleteTopicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteTopicsResponseBody) SetCode(v int32) *BatchDeleteTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteTopicsResponseBody) SetData(v *BatchDeleteTopicsResponseBodyData) *BatchDeleteTopicsResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteTopicsResponseBody) SetMessage(v string) *BatchDeleteTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteTopicsResponseBody) SetRequestId(v string) *BatchDeleteTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteTopicsResponseBody) SetSuccess(v bool) *BatchDeleteTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteTopicsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeleteTopicsResponseBodyData struct {
	FailedCount  *int32                                    `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	Results      *BatchDeleteTopicsResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	SuccessCount *int32                                    `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	Total        *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s BatchDeleteTopicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *BatchDeleteTopicsResponseBodyData) GetResults() *BatchDeleteTopicsResponseBodyDataResults {
	return s.Results
}

func (s *BatchDeleteTopicsResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchDeleteTopicsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *BatchDeleteTopicsResponseBodyData) SetFailedCount(v int32) *BatchDeleteTopicsResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyData) SetResults(v *BatchDeleteTopicsResponseBodyDataResults) *BatchDeleteTopicsResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchDeleteTopicsResponseBodyData) SetSuccessCount(v int32) *BatchDeleteTopicsResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyData) SetTotal(v int32) *BatchDeleteTopicsResponseBodyData {
	s.Total = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyData) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeleteTopicsResponseBodyDataResults struct {
	TopicDeleteResultItemVO []*BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO `json:"TopicDeleteResultItemVO,omitempty" xml:"TopicDeleteResultItemVO,omitempty" type:"Repeated"`
}

func (s BatchDeleteTopicsResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsResponseBodyDataResults) GetTopicDeleteResultItemVO() []*BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	return s.TopicDeleteResultItemVO
}

func (s *BatchDeleteTopicsResponseBodyDataResults) SetTopicDeleteResultItemVO(v []*BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) *BatchDeleteTopicsResponseBodyDataResults {
	s.TopicDeleteResultItemVO = v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResults) Validate() error {
	if s.TopicDeleteResultItemVO != nil {
		for _, item := range s.TopicDeleteResultItemVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetCode() *int32 {
	return s.Code
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetStatus() *string {
	return s.Status
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) GetTopic() *string {
	return s.Topic
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetCode(v int32) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.Code = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetDynamicCode(v string) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.DynamicCode = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetDynamicMessage(v string) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.DynamicMessage = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetMessage(v string) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.Message = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetStatus(v string) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.Status = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetSuccess(v bool) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.Success = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) SetTopic(v string) *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO {
	s.Topic = &v
	return s
}

func (s *BatchDeleteTopicsResponseBodyDataResultsTopicDeleteResultItemVO) Validate() error {
	return dara.Validate(s)
}
