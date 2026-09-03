// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetChargeType(v string) *DescribeNetworkPackagesRequest
	GetInternetChargeType() *string
	SetMaxResults(v int32) *DescribeNetworkPackagesRequest
	GetMaxResults() *int32
	SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest
	GetNetworkPackageId() []*string
	SetNextToken(v string) *DescribeNetworkPackagesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeNetworkPackagesRequest
	GetRegionId() *string
	SetTag(v []*DescribeNetworkPackagesRequestTag) *DescribeNetworkPackagesRequest
	GetTag() []*DescribeNetworkPackagesRequestTag
}

type DescribeNetworkPackagesRequest struct {
	// The billing method of the pay-as-you-go premium Internet bandwidth plan.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The number of entries per page in a paged query.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The IDs of premium Internet bandwidth plans. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	// The pagination token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags. A maximum of 20 tags are supported.
	Tag []*DescribeNetworkPackagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeNetworkPackagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNetworkPackagesRequest) GetNetworkPackageId() []*string {
	return s.NetworkPackageId
}

func (s *DescribeNetworkPackagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNetworkPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkPackagesRequest) GetTag() []*DescribeNetworkPackagesRequestTag {
	return s.Tag
}

func (s *DescribeNetworkPackagesRequest) SetInternetChargeType(v string) *DescribeNetworkPackagesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetMaxResults(v int32) *DescribeNetworkPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNextToken(v string) *DescribeNetworkPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetRegionId(v string) *DescribeNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetTag(v []*DescribeNetworkPackagesRequestTag) *DescribeNetworkPackagesRequest {
	s.Tag = v
	return s
}

func (s *DescribeNetworkPackagesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkPackagesRequestTag struct {
	// The tag key. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkPackagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkPackagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkPackagesRequestTag) SetKey(v string) *DescribeNetworkPackagesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworkPackagesRequestTag) SetValue(v string) *DescribeNetworkPackagesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworkPackagesRequestTag) Validate() error {
	return dara.Validate(s)
}
