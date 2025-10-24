// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowId(v string) *DeactivateMediaWorkflowRequest
	GetMediaWorkflowId() *string
	SetOwnerAccount(v string) *DeactivateMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeactivateMediaWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeactivateMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeactivateMediaWorkflowRequest
	GetResourceOwnerId() *int64
}

type DeactivateMediaWorkflowRequest struct {
	// The ID of the media workflow that is deactivated.
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

func (s DeactivateMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeactivateMediaWorkflowRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *DeactivateMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeactivateMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeactivateMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeactivateMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeactivateMediaWorkflowRequest) SetMediaWorkflowId(v string) *DeactivateMediaWorkflowRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *DeactivateMediaWorkflowRequest) SetOwnerAccount(v string) *DeactivateMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactivateMediaWorkflowRequest) SetOwnerId(v int64) *DeactivateMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactivateMediaWorkflowRequest) SetResourceOwnerAccount(v string) *DeactivateMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactivateMediaWorkflowRequest) SetResourceOwnerId(v int64) *DeactivateMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeactivateMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
