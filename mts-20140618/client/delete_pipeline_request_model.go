// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeletePipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePipelineRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *DeletePipelineRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *DeletePipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePipelineRequest
	GetResourceOwnerId() *int64
}

type DeletePipelineRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue that you want to delete. To obtain the ID of the MPS queue, you can log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeletePipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePipelineRequest) SetOwnerAccount(v string) *DeletePipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePipelineRequest) SetOwnerId(v int64) *DeletePipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePipelineRequest) SetPipelineId(v string) *DeletePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *DeletePipelineRequest) SetResourceOwnerAccount(v string) *DeletePipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePipelineRequest) SetResourceOwnerId(v int64) *DeletePipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePipelineRequest) Validate() error {
	return dara.Validate(s)
}
