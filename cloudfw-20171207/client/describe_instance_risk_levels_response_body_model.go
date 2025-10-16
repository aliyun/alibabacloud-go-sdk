// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRiskLevelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRisks(v []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks) *DescribeInstanceRiskLevelsResponseBody
	GetInstanceRisks() []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks
	SetRequestId(v string) *DescribeInstanceRiskLevelsResponseBody
	GetRequestId() *string
}

type DescribeInstanceRiskLevelsResponseBody struct {
	// The information about the instances.
	InstanceRisks []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks `json:"InstanceRisks,omitempty" xml:"InstanceRisks,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9AABB1B7-C81F-5158-9EF9-B2DD5D3DA014
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBody) GetInstanceRisks() []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	return s.InstanceRisks
}

func (s *DescribeInstanceRiskLevelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRiskLevelsResponseBody) SetInstanceRisks(v []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks) *DescribeInstanceRiskLevelsResponseBody {
	s.InstanceRisks = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBody) SetRequestId(v string) *DescribeInstanceRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBody) Validate() error {
	if s.InstanceRisks != nil {
		for _, item := range s.InstanceRisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRiskLevelsResponseBodyInstanceRisks struct {
	// The risk levels of the Elastic Compute Service (ECS) instance.
	Details []*DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The instance ID of your Cloud Firewall.
	//
	// example:
	//
	// vipcloudfw-cn-7mz2fj8nm0u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The risk levels. Valid values:
	//
	// 	- **medium**
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisks) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) GetDetails() []*DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	return s.Details
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) GetLevel() *string {
	return s.Level
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetDetails(v []*DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.Details = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetInstanceId(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetLevel(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.Level = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails struct {
	// The IP addresses of servers.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The risk levels. Valid values:
	//
	// 	- **medium**
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The type.
	//
	// example:
	//
	// ResourceNotProtected
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) GetLevel() *string {
	return s.Level
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetIp(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetLevel(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Level = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetType(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Type = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) Validate() error {
	return dara.Validate(s)
}
