// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeVpcsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVpcsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVpcsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcsResponseBody
	GetTotalCount() *int32
	SetVpcs(v []*DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody
	GetVpcs() []*DescribeVpcsResponseBodyVpcs
}

type DescribeVpcsResponseBody struct {
	// The maximum number of entries returned on a single page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, it indicates that you have retrieved all the data.
	//
	// This parameter is required.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of VPCs that are available in the specified region.
	Vpcs []*DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVpcsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcsResponseBody) GetVpcs() []*DescribeVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsResponseBody) SetMaxResults(v int32) *DescribeVpcsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetNextToken(v string) *DescribeVpcsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v []*DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		for _, item := range s.Vpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcs struct {
	// Indicates whether the VPC is the default VPC in the specified region. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The status of the VPC. Valid values:
	//
	// 	- `Pending`: The VPC is being configured.
	//
	// 	- `Available`: The VPC is available for use.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of a VPC.
	//
	// example:
	//
	// vpc-bp1qkf2o3xmqc2519****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsResponseBodyVpcs) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetStatus(v string) *DescribeVpcsResponseBodyVpcs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcId(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcName(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) Validate() error {
	return dara.Validate(s)
}
