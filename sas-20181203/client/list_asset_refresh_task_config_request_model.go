// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetRefreshTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshConfigType(v int32) *ListAssetRefreshTaskConfigRequest
	GetRefreshConfigType() *int32
	SetRegionId(v string) *ListAssetRefreshTaskConfigRequest
	GetRegionId() *string
	SetTargetId(v int64) *ListAssetRefreshTaskConfigRequest
	GetTargetId() *int64
}

type ListAssetRefreshTaskConfigRequest struct {
	// The configuration type. Valid values:
	//
	// - **0**: host refresh task
	//
	// - **1**: cloud service refresh task
	//
	// - **2**: AccessKey scheduled verification task.
	//
	// example:
	//
	// 2
	RefreshConfigType *int32 `json:"RefreshConfigType,omitempty" xml:"RefreshConfigType,omitempty"`
	// The region of the Security Center instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the AccessKey record specified when querying an AccessKey scheduled verification task.
	//
	// example:
	//
	// 2295
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s ListAssetRefreshTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetRefreshTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *ListAssetRefreshTaskConfigRequest) GetRefreshConfigType() *int32 {
	return s.RefreshConfigType
}

func (s *ListAssetRefreshTaskConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAssetRefreshTaskConfigRequest) GetTargetId() *int64 {
	return s.TargetId
}

func (s *ListAssetRefreshTaskConfigRequest) SetRefreshConfigType(v int32) *ListAssetRefreshTaskConfigRequest {
	s.RefreshConfigType = &v
	return s
}

func (s *ListAssetRefreshTaskConfigRequest) SetRegionId(v string) *ListAssetRefreshTaskConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ListAssetRefreshTaskConfigRequest) SetTargetId(v int64) *ListAssetRefreshTaskConfigRequest {
	s.TargetId = &v
	return s
}

func (s *ListAssetRefreshTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
