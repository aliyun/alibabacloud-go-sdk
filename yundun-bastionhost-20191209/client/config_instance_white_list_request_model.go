// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceWhiteListRequest
	GetInstanceId() *string
	SetRegionId(v string) *ConfigInstanceWhiteListRequest
	GetRegionId() *string
	SetWhiteList(v []*string) *ConfigInstanceWhiteListRequest
	GetWhiteList() []*string
	SetWhiteListPolicies(v []*ConfigInstanceWhiteListRequestWhiteListPolicies) *ConfigInstanceWhiteListRequest
	GetWhiteListPolicies() []*ConfigInstanceWhiteListRequestWhiteListPolicies
}

type ConfigInstanceWhiteListRequest struct {
	// The ID of the Bastionhost instance to configure.
	//
	// > To obtain the instance ID, call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of IP addresses to add to the whitelist.
	//
	// example:
	//
	// 10.162.XX.XX
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
	// The policies for the public IP address whitelist.
	WhiteListPolicies []*ConfigInstanceWhiteListRequestWhiteListPolicies `json:"WhiteListPolicies,omitempty" xml:"WhiteListPolicies,omitempty" type:"Repeated"`
}

func (s ConfigInstanceWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceWhiteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigInstanceWhiteListRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *ConfigInstanceWhiteListRequest) GetWhiteListPolicies() []*ConfigInstanceWhiteListRequestWhiteListPolicies {
	return s.WhiteListPolicies
}

func (s *ConfigInstanceWhiteListRequest) SetInstanceId(v string) *ConfigInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetRegionId(v string) *ConfigInstanceWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetWhiteList(v []*string) *ConfigInstanceWhiteListRequest {
	s.WhiteList = v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetWhiteListPolicies(v []*ConfigInstanceWhiteListRequestWhiteListPolicies) *ConfigInstanceWhiteListRequest {
	s.WhiteListPolicies = v
	return s
}

func (s *ConfigInstanceWhiteListRequest) Validate() error {
	if s.WhiteListPolicies != nil {
		for _, item := range s.WhiteListPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConfigInstanceWhiteListRequestWhiteListPolicies struct {
	// The description of this whitelist rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP addresses to add to the whitelist. You can specify up to 50 IP addresses, separated by a comma.
	//
	// example:
	//
	// 10.162.XX.XX,192.168.XX.XX
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ConfigInstanceWhiteListRequestWhiteListPolicies) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhiteListRequestWhiteListPolicies) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListRequestWhiteListPolicies) GetDescription() *string {
	return s.Description
}

func (s *ConfigInstanceWhiteListRequestWhiteListPolicies) GetEntry() *string {
	return s.Entry
}

func (s *ConfigInstanceWhiteListRequestWhiteListPolicies) SetDescription(v string) *ConfigInstanceWhiteListRequestWhiteListPolicies {
	s.Description = &v
	return s
}

func (s *ConfigInstanceWhiteListRequestWhiteListPolicies) SetEntry(v string) *ConfigInstanceWhiteListRequestWhiteListPolicies {
	s.Entry = &v
	return s
}

func (s *ConfigInstanceWhiteListRequestWhiteListPolicies) Validate() error {
	return dara.Validate(s)
}
