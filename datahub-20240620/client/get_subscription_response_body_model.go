// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v string) *GetSubscriptionResponseBody
	GetApplication() *string
	SetComment(v string) *GetSubscriptionResponseBody
	GetComment() *string
	SetCreateTime(v int64) *GetSubscriptionResponseBody
	GetCreateTime() *int64
	SetCreator(v string) *GetSubscriptionResponseBody
	GetCreator() *string
	SetProjectName(v string) *GetSubscriptionResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetSubscriptionResponseBody
	GetRequestId() *string
	SetState(v int32) *GetSubscriptionResponseBody
	GetState() *int32
	SetSubscriptionId(v string) *GetSubscriptionResponseBody
	GetSubscriptionId() *string
	SetSuccess(v bool) *GetSubscriptionResponseBody
	GetSuccess() *bool
	SetTopicName(v string) *GetSubscriptionResponseBody
	GetTopicName() *string
	SetType(v string) *GetSubscriptionResponseBody
	GetType() *string
	SetUpdateTime(v int64) *GetSubscriptionResponseBody
	GetUpdateTime() *int64
}

type GetSubscriptionResponseBody struct {
	// example:
	//
	// test_application_name
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1724041098000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1559031978056215
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1741072334529RFEF7
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// USER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1724041098000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBody) GetApplication() *string {
	return s.Application
}

func (s *GetSubscriptionResponseBody) GetComment() *string {
	return s.Comment
}

func (s *GetSubscriptionResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetSubscriptionResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetSubscriptionResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionResponseBody) GetState() *int32 {
	return s.State
}

func (s *GetSubscriptionResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *GetSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSubscriptionResponseBody) GetType() *string {
	return s.Type
}

func (s *GetSubscriptionResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetSubscriptionResponseBody) SetApplication(v string) *GetSubscriptionResponseBody {
	s.Application = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetComment(v string) *GetSubscriptionResponseBody {
	s.Comment = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetCreateTime(v int64) *GetSubscriptionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetCreator(v string) *GetSubscriptionResponseBody {
	s.Creator = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetProjectName(v string) *GetSubscriptionResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetRequestId(v string) *GetSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetState(v int32) *GetSubscriptionResponseBody {
	s.State = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetSubscriptionId(v string) *GetSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetSuccess(v bool) *GetSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetTopicName(v string) *GetSubscriptionResponseBody {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetType(v string) *GetSubscriptionResponseBody {
	s.Type = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetUpdateTime(v int64) *GetSubscriptionResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
