// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest
	GetAggregatorId() *string
	SetFolderId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest
	GetFolderId() *string
	SetRegion(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest
	GetRegion() *string
	SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByResourceTypeRequest
	GetResourceAccountId() *int64
	SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByResourceTypeRequest
	GetResourceOwnerId() *int64
}

type GetAggregateResourceCountsGroupByResourceTypeRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a260626622af0005****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the folder in the resource directory.
	//
	// For more information about how to obtain the ID of a folder, see [View the basic information of a folder](https://help.aliyun.com/document_detail/111223.html).
	//
	// example:
	//
	// r-BU****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources in the account group belong.
	//
	// > You can specify either the ResourceAccountId or ResourceOwnerId parameter. We recommend that you specify the ResourceAccountId parameter.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetAggregatorId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetFolderId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.FolderId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetRegion(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
