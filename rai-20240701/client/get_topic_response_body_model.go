// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTopicResponseBody
	GetCode() *string
	SetGmtModified(v int64) *GetTopicResponseBody
	GetGmtModified() *int64
	SetHttpStatusCode(v int32) *GetTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTopicResponseBody
	GetMessage() *string
	SetPolicyInfoList(v []*GetTopicResponseBodyPolicyInfoList) *GetTopicResponseBody
	GetPolicyInfoList() []*GetTopicResponseBodyPolicyInfoList
	SetRequestId(v string) *GetTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicResponseBody
	GetSuccess() *bool
	SetTopicDefinition(v string) *GetTopicResponseBody
	GetTopicDefinition() *string
	SetTopicExampleInfoList(v []*GetTopicResponseBodyTopicExampleInfoList) *GetTopicResponseBody
	GetTopicExampleInfoList() []*GetTopicResponseBodyTopicExampleInfoList
	SetTopicName(v string) *GetTopicResponseBody
	GetTopicName() *string
}

type GetTopicResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1634122349000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyInfoList []*GetTopicResponseBodyPolicyInfoList `json:"PolicyInfoList,omitempty" xml:"PolicyInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success              *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicDefinition      *string                                     `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	TopicExampleInfoList []*GetTopicResponseBodyTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTopicResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicResponseBody) GetPolicyInfoList() []*GetTopicResponseBodyPolicyInfoList {
	return s.PolicyInfoList
}

func (s *GetTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicResponseBody) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *GetTopicResponseBody) GetTopicExampleInfoList() []*GetTopicResponseBodyTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *GetTopicResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicResponseBody) SetCode(v string) *GetTopicResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicResponseBody) SetGmtModified(v int64) *GetTopicResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetTopicResponseBody) SetHttpStatusCode(v int32) *GetTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicResponseBody) SetMessage(v string) *GetTopicResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicResponseBody) SetPolicyInfoList(v []*GetTopicResponseBodyPolicyInfoList) *GetTopicResponseBody {
	s.PolicyInfoList = v
	return s
}

func (s *GetTopicResponseBody) SetRequestId(v string) *GetTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicResponseBody) SetSuccess(v bool) *GetTopicResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicResponseBody) SetTopicDefinition(v string) *GetTopicResponseBody {
	s.TopicDefinition = &v
	return s
}

func (s *GetTopicResponseBody) SetTopicExampleInfoList(v []*GetTopicResponseBodyTopicExampleInfoList) *GetTopicResponseBody {
	s.TopicExampleInfoList = v
	return s
}

func (s *GetTopicResponseBody) SetTopicName(v string) *GetTopicResponseBody {
	s.TopicName = &v
	return s
}

func (s *GetTopicResponseBody) Validate() error {
	if s.PolicyInfoList != nil {
		for _, item := range s.PolicyInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopicExampleInfoList != nil {
		for _, item := range s.TopicExampleInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicResponseBodyPolicyInfoList struct {
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s GetTopicResponseBodyPolicyInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBodyPolicyInfoList) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyPolicyInfoList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *GetTopicResponseBodyPolicyInfoList) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *GetTopicResponseBodyPolicyInfoList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetTopicResponseBodyPolicyInfoList) SetPolicyId(v int64) *GetTopicResponseBodyPolicyInfoList {
	s.PolicyId = &v
	return s
}

func (s *GetTopicResponseBodyPolicyInfoList) SetPolicyIdentifier(v string) *GetTopicResponseBodyPolicyInfoList {
	s.PolicyIdentifier = &v
	return s
}

func (s *GetTopicResponseBodyPolicyInfoList) SetPolicyName(v string) *GetTopicResponseBodyPolicyInfoList {
	s.PolicyName = &v
	return s
}

func (s *GetTopicResponseBodyPolicyInfoList) Validate() error {
	return dara.Validate(s)
}

type GetTopicResponseBodyTopicExampleInfoList struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2
	ExampleId *int64 `json:"ExampleId,omitempty" xml:"ExampleId,omitempty"`
	// example:
	//
	// 0
	ExampleType *int32 `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s GetTopicResponseBodyTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBodyTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *GetTopicResponseBodyTopicExampleInfoList) GetExampleId() *int64 {
	return s.ExampleId
}

func (s *GetTopicResponseBodyTopicExampleInfoList) GetExampleType() *int32 {
	return s.ExampleType
}

func (s *GetTopicResponseBodyTopicExampleInfoList) SetContent(v string) *GetTopicResponseBodyTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *GetTopicResponseBodyTopicExampleInfoList) SetExampleId(v int64) *GetTopicResponseBodyTopicExampleInfoList {
	s.ExampleId = &v
	return s
}

func (s *GetTopicResponseBodyTopicExampleInfoList) SetExampleType(v int32) *GetTopicResponseBodyTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *GetTopicResponseBodyTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}
