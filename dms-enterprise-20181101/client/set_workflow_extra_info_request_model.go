// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowExtraInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderAddApprovalNode(v bool) *SetWorkflowExtraInfoRequest
	GetRenderAddApprovalNode() *bool
	SetRenderAgree(v bool) *SetWorkflowExtraInfoRequest
	GetRenderAgree() *bool
	SetRenderCancel(v bool) *SetWorkflowExtraInfoRequest
	GetRenderCancel() *bool
	SetRenderReject(v bool) *SetWorkflowExtraInfoRequest
	GetRenderReject() *bool
	SetRenderTransfer(v bool) *SetWorkflowExtraInfoRequest
	GetRenderTransfer() *bool
	SetThirdpartyWorkflowComment(v string) *SetWorkflowExtraInfoRequest
	GetThirdpartyWorkflowComment() *string
	SetThirdpartyWorkflowUrl(v string) *SetWorkflowExtraInfoRequest
	GetThirdpartyWorkflowUrl() *string
	SetTid(v int64) *SetWorkflowExtraInfoRequest
	GetTid() *int64
	SetWorkflowInstanceId(v int64) *SetWorkflowExtraInfoRequest
	GetWorkflowInstanceId() *int64
}

type SetWorkflowExtraInfoRequest struct {
	// example:
	//
	// false
	RenderAddApprovalNode *bool `json:"RenderAddApprovalNode,omitempty" xml:"RenderAddApprovalNode,omitempty"`
	// example:
	//
	// true
	RenderAgree *bool `json:"RenderAgree,omitempty" xml:"RenderAgree,omitempty"`
	// example:
	//
	// true
	RenderCancel *bool `json:"RenderCancel,omitempty" xml:"RenderCancel,omitempty"`
	// example:
	//
	// false
	RenderReject *bool `json:"RenderReject,omitempty" xml:"RenderReject,omitempty"`
	// example:
	//
	// true
	RenderTransfer *bool `json:"RenderTransfer,omitempty" xml:"RenderTransfer,omitempty"`
	// example:
	//
	// test
	ThirdpartyWorkflowComment *string `json:"ThirdpartyWorkflowComment,omitempty" xml:"ThirdpartyWorkflowComment,omitempty"`
	// example:
	//
	// https://xxx
	ThirdpartyWorkflowUrl *string `json:"ThirdpartyWorkflowUrl,omitempty" xml:"ThirdpartyWorkflowUrl,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 184****
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s SetWorkflowExtraInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowExtraInfoRequest) GoString() string {
	return s.String()
}

func (s *SetWorkflowExtraInfoRequest) GetRenderAddApprovalNode() *bool {
	return s.RenderAddApprovalNode
}

func (s *SetWorkflowExtraInfoRequest) GetRenderAgree() *bool {
	return s.RenderAgree
}

func (s *SetWorkflowExtraInfoRequest) GetRenderCancel() *bool {
	return s.RenderCancel
}

func (s *SetWorkflowExtraInfoRequest) GetRenderReject() *bool {
	return s.RenderReject
}

func (s *SetWorkflowExtraInfoRequest) GetRenderTransfer() *bool {
	return s.RenderTransfer
}

func (s *SetWorkflowExtraInfoRequest) GetThirdpartyWorkflowComment() *string {
	return s.ThirdpartyWorkflowComment
}

func (s *SetWorkflowExtraInfoRequest) GetThirdpartyWorkflowUrl() *string {
	return s.ThirdpartyWorkflowUrl
}

func (s *SetWorkflowExtraInfoRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SetWorkflowExtraInfoRequest) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *SetWorkflowExtraInfoRequest) SetRenderAddApprovalNode(v bool) *SetWorkflowExtraInfoRequest {
	s.RenderAddApprovalNode = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetRenderAgree(v bool) *SetWorkflowExtraInfoRequest {
	s.RenderAgree = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetRenderCancel(v bool) *SetWorkflowExtraInfoRequest {
	s.RenderCancel = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetRenderReject(v bool) *SetWorkflowExtraInfoRequest {
	s.RenderReject = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetRenderTransfer(v bool) *SetWorkflowExtraInfoRequest {
	s.RenderTransfer = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetThirdpartyWorkflowComment(v string) *SetWorkflowExtraInfoRequest {
	s.ThirdpartyWorkflowComment = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetThirdpartyWorkflowUrl(v string) *SetWorkflowExtraInfoRequest {
	s.ThirdpartyWorkflowUrl = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetTid(v int64) *SetWorkflowExtraInfoRequest {
	s.Tid = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) SetWorkflowInstanceId(v int64) *SetWorkflowExtraInfoRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *SetWorkflowExtraInfoRequest) Validate() error {
	return dara.Validate(s)
}
