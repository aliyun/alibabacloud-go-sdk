// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendNextActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrompt(v string) *RecommendNextActionsRequest
	GetCustomPrompt() *string
	SetOutputType(v string) *RecommendNextActionsRequest
	GetOutputType() *string
	SetRecentMessageCount(v int64) *RecommendNextActionsRequest
	GetRecentMessageCount() *int64
	SetSessionId(v string) *RecommendNextActionsRequest
	GetSessionId() *string
	SetTenantId(v string) *RecommendNextActionsRequest
	GetTenantId() *string
}

type RecommendNextActionsRequest struct {
	// The extraction instruction.
	//
	// example:
	//
	// Only recommend next steps related to data analysis
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// The output type: `conversation/skill/task`.
	//
	// example:
	//
	// followUpOnly
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
	// The number of recent messages used to assemble contextual information.
	//
	// example:
	//
	// 10
	RecentMessageCount *int64 `json:"recentMessageCount,omitempty" xml:"recentMessageCount,omitempty"`
	// The session ID to filter by. If specified, returns all Active/Expired status information associated with this session.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through winnexo-cli using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RecommendNextActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s RecommendNextActionsRequest) GoString() string {
	return s.String()
}

func (s *RecommendNextActionsRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RecommendNextActionsRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *RecommendNextActionsRequest) GetRecentMessageCount() *int64 {
	return s.RecentMessageCount
}

func (s *RecommendNextActionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RecommendNextActionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RecommendNextActionsRequest) SetCustomPrompt(v string) *RecommendNextActionsRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RecommendNextActionsRequest) SetOutputType(v string) *RecommendNextActionsRequest {
	s.OutputType = &v
	return s
}

func (s *RecommendNextActionsRequest) SetRecentMessageCount(v int64) *RecommendNextActionsRequest {
	s.RecentMessageCount = &v
	return s
}

func (s *RecommendNextActionsRequest) SetSessionId(v string) *RecommendNextActionsRequest {
	s.SessionId = &v
	return s
}

func (s *RecommendNextActionsRequest) SetTenantId(v string) *RecommendNextActionsRequest {
	s.TenantId = &v
	return s
}

func (s *RecommendNextActionsRequest) Validate() error {
	return dara.Validate(s)
}
