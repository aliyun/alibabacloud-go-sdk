// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingDocInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetBiddingDocInfoRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetBiddingDocInfoRequest
	GetWorkspaceId() *string
}

type GetBiddingDocInfoRequest struct {
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetBiddingDocInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingDocInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBiddingDocInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBiddingDocInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetBiddingDocInfoRequest) SetTaskId(v string) *GetBiddingDocInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetBiddingDocInfoRequest) SetWorkspaceId(v string) *GetBiddingDocInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetBiddingDocInfoRequest) Validate() error {
	return dara.Validate(s)
}
