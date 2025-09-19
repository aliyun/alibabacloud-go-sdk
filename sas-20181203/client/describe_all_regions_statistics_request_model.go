// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllRegionsStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeAllRegionsStatisticsRequest
	GetFrom() *string
	SetGroupId(v int64) *DescribeAllRegionsStatisticsRequest
	GetGroupId() *int64
	SetRemark(v string) *DescribeAllRegionsStatisticsRequest
	GetRemark() *string
	SetSourceIp(v string) *DescribeAllRegionsStatisticsRequest
	GetSourceIp() *string
}

type DescribeAllRegionsStatisticsRequest struct {
	// The source of the request. Default value: **aqs**. Valid values:
	//
	// 	- **sas**: Security Center.
	//
	// 	- **aqs**: Server Guard.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the asset group that you want to query.
	//
	// >  You can call the [DescribeAllGroups](https://help.aliyun.com/document_detail/130972.html) operation to query the ID.
	//
	// example:
	//
	// 1161****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name or public IP address of the asset.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 33.80.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAllRegionsStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRegionsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeAllRegionsStatisticsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeAllRegionsStatisticsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAllRegionsStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAllRegionsStatisticsRequest) SetFrom(v string) *DescribeAllRegionsStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetGroupId(v int64) *DescribeAllRegionsStatisticsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetRemark(v string) *DescribeAllRegionsStatisticsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetSourceIp(v string) *DescribeAllRegionsStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
