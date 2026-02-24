// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetAggregatorId() *string
	SetFolderId(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetFolderId() *string
	SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceAccountId() *int64
	SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceType() *string
}

type GetAggregateResourceCountsGroupByRegionRequest struct {
	// This parameter is required.
	AggregatorId      *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	FolderId          *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetAggregatorId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetFolderId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.FolderId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceType(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) Validate() error {
	return dara.Validate(s)
}
