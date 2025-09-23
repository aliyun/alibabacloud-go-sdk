// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncWritingBiddingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyKeyword(v string) *AsyncWritingBiddingDocRequest
	GetCompanyKeyword() *string
	SetPrompt(v string) *AsyncWritingBiddingDocRequest
	GetPrompt() *string
	SetTaskId(v string) *AsyncWritingBiddingDocRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *AsyncWritingBiddingDocRequest
	GetWorkspaceId() *string
}

type AsyncWritingBiddingDocRequest struct {
	// example:
	//
	// comany name
	CompanyKeyword *string `json:"CompanyKeyword,omitempty" xml:"CompanyKeyword,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncWritingBiddingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncWritingBiddingDocRequest) GoString() string {
	return s.String()
}

func (s *AsyncWritingBiddingDocRequest) GetCompanyKeyword() *string {
	return s.CompanyKeyword
}

func (s *AsyncWritingBiddingDocRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *AsyncWritingBiddingDocRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncWritingBiddingDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncWritingBiddingDocRequest) SetCompanyKeyword(v string) *AsyncWritingBiddingDocRequest {
	s.CompanyKeyword = &v
	return s
}

func (s *AsyncWritingBiddingDocRequest) SetPrompt(v string) *AsyncWritingBiddingDocRequest {
	s.Prompt = &v
	return s
}

func (s *AsyncWritingBiddingDocRequest) SetTaskId(v string) *AsyncWritingBiddingDocRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncWritingBiddingDocRequest) SetWorkspaceId(v string) *AsyncWritingBiddingDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncWritingBiddingDocRequest) Validate() error {
	return dara.Validate(s)
}
