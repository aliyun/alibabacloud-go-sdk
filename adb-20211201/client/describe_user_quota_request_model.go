// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeUserQuotaRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeUserQuotaRequest
	GetRegionId() *string
}

type DescribeUserQuotaRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1qjt3o18d86987
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeUserQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserQuotaRequest) SetDBClusterId(v string) *DescribeUserQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeUserQuotaRequest) SetRegionId(v string) *DescribeUserQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
