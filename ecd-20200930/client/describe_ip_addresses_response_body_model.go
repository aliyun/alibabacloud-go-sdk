// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpAddresses(v []*DescribeIpAddressesResponseBodyIpAddresses) *DescribeIpAddressesResponseBody
	GetIpAddresses() []*DescribeIpAddressesResponseBodyIpAddresses
	SetMaxResults(v int32) *DescribeIpAddressesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeIpAddressesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeIpAddressesResponseBody
	GetRequestId() *string
}

type DescribeIpAddressesResponseBody struct {
	IpAddresses []*DescribeIpAddressesResponseBodyIpAddresses `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	MaxResults  *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpAddressesResponseBody) GetIpAddresses() []*DescribeIpAddressesResponseBodyIpAddresses {
	return s.IpAddresses
}

func (s *DescribeIpAddressesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeIpAddressesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpAddressesResponseBody) SetIpAddresses(v []*DescribeIpAddressesResponseBodyIpAddresses) *DescribeIpAddressesResponseBody {
	s.IpAddresses = v
	return s
}

func (s *DescribeIpAddressesResponseBody) SetMaxResults(v int32) *DescribeIpAddressesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeIpAddressesResponseBody) SetNextToken(v string) *DescribeIpAddressesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeIpAddressesResponseBody) SetRequestId(v string) *DescribeIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpAddressesResponseBody) Validate() error {
	if s.IpAddresses != nil {
		for _, item := range s.IpAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpAddressesResponseBodyIpAddresses struct {
	CreateByWuying *bool   `json:"CreateByWuying,omitempty" xml:"CreateByWuying,omitempty"`
	EipAddress     *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	EipId          *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	EipStatus      *string `json:"EipStatus,omitempty" xml:"EipStatus,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeIpAddressesResponseBodyIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAddressesResponseBodyIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetCreateByWuying() *bool {
	return s.CreateByWuying
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetEipAddress() *string {
	return s.EipAddress
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetEipId() *string {
	return s.EipId
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetEipStatus() *string {
	return s.EipStatus
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetCreateByWuying(v bool) *DescribeIpAddressesResponseBodyIpAddresses {
	s.CreateByWuying = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetEipAddress(v string) *DescribeIpAddressesResponseBodyIpAddresses {
	s.EipAddress = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetEipId(v string) *DescribeIpAddressesResponseBodyIpAddresses {
	s.EipId = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetEipStatus(v string) *DescribeIpAddressesResponseBodyIpAddresses {
	s.EipStatus = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetInstanceId(v string) *DescribeIpAddressesResponseBodyIpAddresses {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) SetInstanceType(v string) *DescribeIpAddressesResponseBodyIpAddresses {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpAddressesResponseBodyIpAddresses) Validate() error {
	return dara.Validate(s)
}
