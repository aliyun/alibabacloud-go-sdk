// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowId(v string) *DeleteMediaWorkflowRequest
	GetMediaWorkflowId() *string
	SetOwnerAccount(v string) *DeleteMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteMediaWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMediaWorkflowRequest
	GetResourceOwnerId() *int64
}

type DeleteMediaWorkflowRequest struct {
	// The ID of the media workflow that you want to delete. To obtain the ID of the media workflow, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings*	- in the left-side navigation pane.
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

func (s DeleteMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaWorkflowRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *DeleteMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMediaWorkflowRequest) SetMediaWorkflowId(v string) *DeleteMediaWorkflowRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *DeleteMediaWorkflowRequest) SetOwnerAccount(v string) *DeleteMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMediaWorkflowRequest) SetOwnerId(v int64) *DeleteMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMediaWorkflowRequest) SetResourceOwnerAccount(v string) *DeleteMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMediaWorkflowRequest) SetResourceOwnerId(v int64) *DeleteMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
