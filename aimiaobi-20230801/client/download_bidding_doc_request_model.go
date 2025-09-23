// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadBiddingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DownloadBiddingDocRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *DownloadBiddingDocRequest
	GetWorkspaceId() *string
}

type DownloadBiddingDocRequest struct {
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DownloadBiddingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadBiddingDocRequest) GoString() string {
	return s.String()
}

func (s *DownloadBiddingDocRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DownloadBiddingDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DownloadBiddingDocRequest) SetTaskId(v string) *DownloadBiddingDocRequest {
	s.TaskId = &v
	return s
}

func (s *DownloadBiddingDocRequest) SetWorkspaceId(v string) *DownloadBiddingDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DownloadBiddingDocRequest) Validate() error {
	return dara.Validate(s)
}
