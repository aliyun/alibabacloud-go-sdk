// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowPreviewUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetFlowPreviewUrlRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *GetFlowPreviewUrlRequest
	GetFlowId() *string
	SetOwnerId(v int64) *GetFlowPreviewUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetFlowPreviewUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFlowPreviewUrlRequest
	GetResourceOwnerId() *int64
}

type GetFlowPreviewUrlRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetFlowPreviewUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowPreviewUrlRequest) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetFlowPreviewUrlRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowPreviewUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFlowPreviewUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFlowPreviewUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFlowPreviewUrlRequest) SetCustSpaceId(v string) *GetFlowPreviewUrlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) SetFlowId(v string) *GetFlowPreviewUrlRequest {
	s.FlowId = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) SetOwnerId(v int64) *GetFlowPreviewUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) SetResourceOwnerAccount(v string) *GetFlowPreviewUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) SetResourceOwnerId(v int64) *GetFlowPreviewUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) Validate() error {
	return dara.Validate(s)
}
