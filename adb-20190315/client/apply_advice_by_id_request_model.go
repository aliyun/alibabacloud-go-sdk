// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAdviceByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceDate(v int64) *ApplyAdviceByIdRequest
	GetAdviceDate() *int64
	SetAdviceId(v string) *ApplyAdviceByIdRequest
	GetAdviceId() *string
	SetApplyType(v string) *ApplyAdviceByIdRequest
	GetApplyType() *string
	SetBuildImmediately(v bool) *ApplyAdviceByIdRequest
	GetBuildImmediately() *bool
	SetDBClusterId(v string) *ApplyAdviceByIdRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ApplyAdviceByIdRequest
	GetRegionId() *string
}

type ApplyAdviceByIdRequest struct {
	// The date when the suggestion is generated. Specify the date in the yyyyMMdd format. The date must be in UTC.
	//
	// example:
	//
	// 20221101
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The suggestion ID.
	//
	// example:
	//
	// 0baf1f52-53df-487f-8292-99a03716****
	AdviceId         *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	ApplyType        *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	BuildImmediately *bool   `json:"BuildImmediately,omitempty" xml:"BuildImmediately,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyAdviceByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAdviceByIdRequest) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdRequest) GetAdviceDate() *int64 {
	return s.AdviceDate
}

func (s *ApplyAdviceByIdRequest) GetAdviceId() *string {
	return s.AdviceId
}

func (s *ApplyAdviceByIdRequest) GetApplyType() *string {
	return s.ApplyType
}

func (s *ApplyAdviceByIdRequest) GetBuildImmediately() *bool {
	return s.BuildImmediately
}

func (s *ApplyAdviceByIdRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ApplyAdviceByIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyAdviceByIdRequest) SetAdviceDate(v int64) *ApplyAdviceByIdRequest {
	s.AdviceDate = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetAdviceId(v string) *ApplyAdviceByIdRequest {
	s.AdviceId = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetApplyType(v string) *ApplyAdviceByIdRequest {
	s.ApplyType = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetBuildImmediately(v bool) *ApplyAdviceByIdRequest {
	s.BuildImmediately = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetDBClusterId(v string) *ApplyAdviceByIdRequest {
	s.DBClusterId = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetRegionId(v string) *ApplyAdviceByIdRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyAdviceByIdRequest) Validate() error {
	return dara.Validate(s)
}
