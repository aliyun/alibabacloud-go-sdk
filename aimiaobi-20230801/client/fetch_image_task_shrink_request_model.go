// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImageTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FetchImageTaskShrinkRequest
	GetAgentKey() *string
	SetArticleTaskId(v string) *FetchImageTaskShrinkRequest
	GetArticleTaskId() *string
	SetTaskIdListShrink(v string) *FetchImageTaskShrinkRequest
	GetTaskIdListShrink() *string
}

type FetchImageTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["9d8c9185-3f75-4a20-aca1-c5bb53dd97b3"]
	TaskIdListShrink *string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty"`
}

func (s FetchImageTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FetchImageTaskShrinkRequest) GetArticleTaskId() *string {
	return s.ArticleTaskId
}

func (s *FetchImageTaskShrinkRequest) GetTaskIdListShrink() *string {
	return s.TaskIdListShrink
}

func (s *FetchImageTaskShrinkRequest) SetAgentKey(v string) *FetchImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetArticleTaskId(v string) *FetchImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetTaskIdListShrink(v string) *FetchImageTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
