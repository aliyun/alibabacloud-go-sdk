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
	// Specifies whether the Sign button is displayed in the ticket approval section of the DMS console for a third-party approval workflow. Valid values:
	//
	// 	- **true*	- (default): The Sign button is displayed.
	//
	// 	- **false**: The Sign button is not displayed.
	//
	// example:
	//
	// false
	RenderAddApprovalNode *bool `json:"RenderAddApprovalNode,omitempty" xml:"RenderAddApprovalNode,omitempty"`
	// Specifies whether the Agree button is displayed in the ticket approval section of the DMS console for a third-party approval workflow. Valid values:
	//
	// 	- **true*	- (default): The Agree button is displayed.
	//
	// 	- **false**: The Agree button is not displayed.
	//
	// example:
	//
	// true
	RenderAgree *bool `json:"RenderAgree,omitempty" xml:"RenderAgree,omitempty"`
	// Specifies whether the Revoke button is displayed in the ticket approval section of the DMS console for a third-party approval workflow. Valid values:
	//
	// 	- **true*	- (default): The Revoke button is displayed.
	//
	// 	- **false**: The Revoke button is not displayed.
	//
	// example:
	//
	// true
	RenderCancel *bool `json:"RenderCancel,omitempty" xml:"RenderCancel,omitempty"`
	// Specifies whether the Reject button is displayed in the ticket approval section of the DMS console for a third-party approval workflow. Valid values:
	//
	// 	- **true*	- (default): The Reject button is displayed.
	//
	// 	- **false**: The Reject button is not displayed.
	//
	// example:
	//
	// false
	RenderReject *bool `json:"RenderReject,omitempty" xml:"RenderReject,omitempty"`
	// Specifies whether the Forward button is displayed in the ticket approval section of the DMS console for a third-party approval workflow. Valid values:
	//
	// 	- **true*	- (default): The Forward button is displayed.
	//
	// 	- **false**: The Forward button is not displayed.
	//
	// example:
	//
	// true
	RenderTransfer *bool `json:"RenderTransfer,omitempty" xml:"RenderTransfer,omitempty"`
	// The remarks of approval workflow for third parties.
	//
	// example:
	//
	// test
	ThirdpartyWorkflowComment *string `json:"ThirdpartyWorkflowComment,omitempty" xml:"ThirdpartyWorkflowComment,omitempty"`
	// The link of approval workflow for third parties.
	//
	// example:
	//
	// https://xxx
	ThirdpartyWorkflowUrl *string `json:"ThirdpartyWorkflowUrl,omitempty" xml:"ThirdpartyWorkflowUrl,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the approval workflow. You can call the [GetOrderBaseInfo](https://help.aliyun.com/document_detail/144642.html) operation to query the ID of the approval workflow.
	//
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
