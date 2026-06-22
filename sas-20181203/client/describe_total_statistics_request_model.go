// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeTotalStatisticsRequest
	GetFrom() *string
	SetGroupId(v int64) *DescribeTotalStatisticsRequest
	GetGroupId() *int64
	SetRemark(v string) *DescribeTotalStatisticsRequest
	GetRemark() *string
}

type DescribeTotalStatisticsRequest struct {
	// The source of the data request. Default value: **aqs**. Valid values:
	//
	// - **sas**: The data request comes from Security Center.
	//
	// - **aqs**: The data request comes from Server Guard.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the asset group to query.
	//
	// > You can call the [DescribeAllGroups](https://help.aliyun.com/document_detail/130972.html) operation to obtain this parameter.
	//
	// example:
	//
	// 8076980
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The asset name or public IP address.
	//
	// example:
	//
	// 222.185.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeTotalStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeTotalStatisticsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeTotalStatisticsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeTotalStatisticsRequest) SetFrom(v string) *DescribeTotalStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeTotalStatisticsRequest) SetGroupId(v int64) *DescribeTotalStatisticsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTotalStatisticsRequest) SetRemark(v string) *DescribeTotalStatisticsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeTotalStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
