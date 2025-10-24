// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *QueryPipelineListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryPipelineListRequest
	GetOwnerId() *int64
	SetPipelineIds(v string) *QueryPipelineListRequest
	GetPipelineIds() *string
	SetResourceOwnerAccount(v string) *QueryPipelineListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPipelineListRequest
	GetResourceOwnerId() *int64
}

type QueryPipelineListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of the MPS queues that you want to query. To view the IDs, you can log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane. You can query up to 10 MPS queues at a time. Separate multiple IDs of MPS queues with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****,72dfa5e679ab4be9a3ed9974c736****
	PipelineIds          *string `json:"PipelineIds,omitempty" xml:"PipelineIds,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPipelineListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListRequest) GoString() string {
	return s.String()
}

func (s *QueryPipelineListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryPipelineListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPipelineListRequest) GetPipelineIds() *string {
	return s.PipelineIds
}

func (s *QueryPipelineListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPipelineListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPipelineListRequest) SetOwnerAccount(v string) *QueryPipelineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryPipelineListRequest) SetOwnerId(v int64) *QueryPipelineListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPipelineListRequest) SetPipelineIds(v string) *QueryPipelineListRequest {
	s.PipelineIds = &v
	return s
}

func (s *QueryPipelineListRequest) SetResourceOwnerAccount(v string) *QueryPipelineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPipelineListRequest) SetResourceOwnerId(v int64) *QueryPipelineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPipelineListRequest) Validate() error {
	return dara.Validate(s)
}
