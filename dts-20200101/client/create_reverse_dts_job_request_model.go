// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReverseDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *CreateReverseDtsJobRequest
	GetDtsJobId() *string
	SetResourceGroupId(v string) *CreateReverseDtsJobRequest
	GetResourceGroupId() *string
	SetShardPassword(v string) *CreateReverseDtsJobRequest
	GetShardPassword() *string
	SetShardUsername(v string) *CreateReverseDtsJobRequest
	GetShardUsername() *string
}

type CreateReverseDtsJobRequest struct {
	// The ID of the synchronization or migration task, which can be queried by calling [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// n99m9jx822k****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Resource GroupId
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Shard Password
	//
	// example:
	//
	// DTStest****
	ShardPassword *string `json:"ShardPassword,omitempty" xml:"ShardPassword,omitempty"`
	// Shard User name
	//
	// example:
	//
	// dtstest
	ShardUsername *string `json:"ShardUsername,omitempty" xml:"ShardUsername,omitempty"`
}

func (s CreateReverseDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReverseDtsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateReverseDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateReverseDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateReverseDtsJobRequest) GetShardPassword() *string {
	return s.ShardPassword
}

func (s *CreateReverseDtsJobRequest) GetShardUsername() *string {
	return s.ShardUsername
}

func (s *CreateReverseDtsJobRequest) SetDtsJobId(v string) *CreateReverseDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateReverseDtsJobRequest) SetResourceGroupId(v string) *CreateReverseDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateReverseDtsJobRequest) SetShardPassword(v string) *CreateReverseDtsJobRequest {
	s.ShardPassword = &v
	return s
}

func (s *CreateReverseDtsJobRequest) SetShardUsername(v string) *CreateReverseDtsJobRequest {
	s.ShardUsername = &v
	return s
}

func (s *CreateReverseDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
