// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRiskListRequest
	GetInstanceId() *string
	SetPageSize(v int64) *GetRiskListRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetRiskListRequest
	GetRegionId() *string
	SetStartIndex(v int64) *GetRiskListRequest
	GetStartIndex() *int64
}

type GetRiskListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-pe335pgxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	StartIndex *int64 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s GetRiskListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRiskListRequest) GoString() string {
	return s.String()
}

func (s *GetRiskListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRiskListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRiskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRiskListRequest) GetStartIndex() *int64 {
	return s.StartIndex
}

func (s *GetRiskListRequest) SetInstanceId(v string) *GetRiskListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRiskListRequest) SetPageSize(v int64) *GetRiskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetRiskListRequest) SetRegionId(v string) *GetRiskListRequest {
	s.RegionId = &v
	return s
}

func (s *GetRiskListRequest) SetStartIndex(v int64) *GetRiskListRequest {
	s.StartIndex = &v
	return s
}

func (s *GetRiskListRequest) Validate() error {
	return dara.Validate(s)
}
