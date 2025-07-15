// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImageTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FetchImageTaskRequest
	GetAgentKey() *string
	SetArticleTaskId(v string) *FetchImageTaskRequest
	GetArticleTaskId() *string
	SetTaskIdList(v []*string) *FetchImageTaskRequest
	GetTaskIdList() []*string
}

type FetchImageTaskRequest struct {
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
	TaskIdList []*string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FetchImageTaskRequest) GetArticleTaskId() *string {
	return s.ArticleTaskId
}

func (s *FetchImageTaskRequest) GetTaskIdList() []*string {
	return s.TaskIdList
}

func (s *FetchImageTaskRequest) SetAgentKey(v string) *FetchImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskRequest) SetArticleTaskId(v string) *FetchImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskRequest) SetTaskIdList(v []*string) *FetchImageTaskRequest {
	s.TaskIdList = v
	return s
}

func (s *FetchImageTaskRequest) Validate() error {
	return dara.Validate(s)
}
