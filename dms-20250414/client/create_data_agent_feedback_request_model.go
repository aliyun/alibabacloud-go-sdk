// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateDataAgentFeedbackRequest
	GetDMSUnit() *string
	SetFeedbackContent(v string) *CreateDataAgentFeedbackRequest
	GetFeedbackContent() *string
	SetFeedbackType(v string) *CreateDataAgentFeedbackRequest
	GetFeedbackType() *string
	SetLikeValue(v int32) *CreateDataAgentFeedbackRequest
	GetLikeValue() *int32
	SetSessionId(v string) *CreateDataAgentFeedbackRequest
	GetSessionId() *string
	SetTargetId(v string) *CreateDataAgentFeedbackRequest
	GetTargetId() *string
	SetTargetType(v string) *CreateDataAgentFeedbackRequest
	GetTargetType() *string
	SetWorkspaceId(v string) *CreateDataAgentFeedbackRequest
	GetWorkspaceId() *string
}

type CreateDataAgentFeedbackRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// {"feedback_type":"PRODUCT_SUGGESTION","user_feedback": "test","email":"yourname@example.com","is_authorized":"Y"}
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// example:
	//
	// ISSUE_REPORT
	FeedbackType *string `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	// example:
	//
	// 1
	LikeValue *int32 `json:"LikeValue,omitempty" xml:"LikeValue,omitempty"`
	// example:
	//
	// h8r********4fch
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// h8r********4fch_sdesfews
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// example:
	//
	// SESSION
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentFeedbackRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentFeedbackRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateDataAgentFeedbackRequest) GetFeedbackContent() *string {
	return s.FeedbackContent
}

func (s *CreateDataAgentFeedbackRequest) GetFeedbackType() *string {
	return s.FeedbackType
}

func (s *CreateDataAgentFeedbackRequest) GetLikeValue() *int32 {
	return s.LikeValue
}

func (s *CreateDataAgentFeedbackRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDataAgentFeedbackRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateDataAgentFeedbackRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDataAgentFeedbackRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentFeedbackRequest) SetDMSUnit(v string) *CreateDataAgentFeedbackRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetFeedbackContent(v string) *CreateDataAgentFeedbackRequest {
	s.FeedbackContent = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetFeedbackType(v string) *CreateDataAgentFeedbackRequest {
	s.FeedbackType = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetLikeValue(v int32) *CreateDataAgentFeedbackRequest {
	s.LikeValue = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetSessionId(v string) *CreateDataAgentFeedbackRequest {
	s.SessionId = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetTargetId(v string) *CreateDataAgentFeedbackRequest {
	s.TargetId = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetTargetType(v string) *CreateDataAgentFeedbackRequest {
	s.TargetType = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) SetWorkspaceId(v string) *CreateDataAgentFeedbackRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
