// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ReadFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ReadFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *ReadFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *ReadFlowVersionShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *ReadFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReadFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReadFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ReadFlowVersionShrinkRequest
	GetStatus() *string
}

type ReadFlowVersionShrinkRequest struct {
	// The business tenant code. Default value: ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The business extension information. Default value: an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The flow code. You can view the flow code on the [flow editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. You can click a flow name on the [flow editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page to go to the canvas orchestration page and view the flow version, or call [ListFlowVersion](https://help.aliyun.com/document_detail/2937202.html) to retrieve the flow version.
	//
	// example:
	//
	// 1
	FlowVersion          *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The flow version status. Valid values:
	//
	// - DRAFT: draft.
	//
	// - DELETED: deleted.
	//
	// - ONLINE: online.
	//
	// - OFFLINE: offline.
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReadFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ReadFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ReadFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ReadFlowVersionShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *ReadFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReadFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReadFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReadFlowVersionShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ReadFlowVersionShrinkRequest) SetBizCode(v string) *ReadFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetBizExtendShrink(v string) *ReadFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetFlowCode(v string) *ReadFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetFlowVersion(v string) *ReadFlowVersionShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetOwnerId(v int64) *ReadFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *ReadFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *ReadFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) SetStatus(v string) *ReadFlowVersionShrinkRequest {
	s.Status = &v
	return s
}

func (s *ReadFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
