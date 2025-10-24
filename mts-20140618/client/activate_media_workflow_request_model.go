// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowId(v string) *ActivateMediaWorkflowRequest
	GetMediaWorkflowId() *string
	SetOwnerAccount(v string) *ActivateMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ActivateMediaWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ActivateMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ActivateMediaWorkflowRequest
	GetResourceOwnerId() *int64
}

type ActivateMediaWorkflowRequest struct {
	// The ID of the media workflow. You can obtain the ID from the response of the [AddMediaWorkflow](https://help.aliyun.com/document_detail/44437.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e9181d4****
	MediaWorkflowId      *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ActivateMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *ActivateMediaWorkflowRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *ActivateMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ActivateMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ActivateMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ActivateMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ActivateMediaWorkflowRequest) SetMediaWorkflowId(v string) *ActivateMediaWorkflowRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *ActivateMediaWorkflowRequest) SetOwnerAccount(v string) *ActivateMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ActivateMediaWorkflowRequest) SetOwnerId(v int64) *ActivateMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *ActivateMediaWorkflowRequest) SetResourceOwnerAccount(v string) *ActivateMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ActivateMediaWorkflowRequest) SetResourceOwnerId(v int64) *ActivateMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActivateMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
