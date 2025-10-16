// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRiskLevelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstanceRiskLevelsRequestInstances) *DescribeInstanceRiskLevelsRequest
	GetInstances() []*DescribeInstanceRiskLevelsRequestInstances
	SetLang(v string) *DescribeInstanceRiskLevelsRequest
	GetLang() *string
}

type DescribeInstanceRiskLevelsRequest struct {
	// The information about the instances.
	Instances []*DescribeInstanceRiskLevelsRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeInstanceRiskLevelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsRequest) GetInstances() []*DescribeInstanceRiskLevelsRequestInstances {
	return s.Instances
}

func (s *DescribeInstanceRiskLevelsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstanceRiskLevelsRequest) SetInstances(v []*DescribeInstanceRiskLevelsRequestInstances) *DescribeInstanceRiskLevelsRequest {
	s.Instances = v
	return s
}

func (s *DescribeInstanceRiskLevelsRequest) SetLang(v string) *DescribeInstanceRiskLevelsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRiskLevelsRequestInstances struct {
	// The instance ID of your Cloud Firewall.
	//
	// example:
	//
	// vipcloudfw-cn-7mz2fj8nm0u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP addresses of instances.
	InternetIp []*string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty" type:"Repeated"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 172.17.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// 181ad081-e4f2-4e3e-b925-03b67f648397
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceRiskLevelsRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsRequestInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRiskLevelsRequestInstances) GetInternetIp() []*string {
	return s.InternetIp
}

func (s *DescribeInstanceRiskLevelsRequestInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeInstanceRiskLevelsRequestInstances) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetInstanceId(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetInternetIp(v []*string) *DescribeInstanceRiskLevelsRequestInstances {
	s.InternetIp = v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetIntranetIp(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetUuid(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) Validate() error {
	return dara.Validate(s)
}
