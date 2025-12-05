// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListVpcConfigsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListVpcConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVpcConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVpcConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVpcConfigsResponseBody
	GetTotalCount() *int64
	SetVpcConfigList(v []*ListVpcConfigsResponseBodyVpcConfigList) *ListVpcConfigsResponseBody
	GetVpcConfigList() []*ListVpcConfigsResponseBodyVpcConfigList
}

type ListVpcConfigsResponseBody struct {
	// example:
	//
	// The specified parameter is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9E4FE41E-782D-5CF9-BFAF-2F369EFD803A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount    *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcConfigList []*ListVpcConfigsResponseBodyVpcConfigList `json:"VpcConfigList,omitempty" xml:"VpcConfigList,omitempty" type:"Repeated"`
}

func (s ListVpcConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVpcConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVpcConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVpcConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVpcConfigsResponseBody) GetVpcConfigList() []*ListVpcConfigsResponseBodyVpcConfigList {
	return s.VpcConfigList
}

func (s *ListVpcConfigsResponseBody) SetMessage(v string) *ListVpcConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpcConfigsResponseBody) SetPageNumber(v int32) *ListVpcConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVpcConfigsResponseBody) SetPageSize(v int32) *ListVpcConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcConfigsResponseBody) SetRequestId(v string) *ListVpcConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcConfigsResponseBody) SetTotalCount(v int64) *ListVpcConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcConfigsResponseBody) SetVpcConfigList(v []*ListVpcConfigsResponseBodyVpcConfigList) *ListVpcConfigsResponseBody {
	s.VpcConfigList = v
	return s
}

func (s *ListVpcConfigsResponseBody) Validate() error {
	if s.VpcConfigList != nil {
		for _, item := range s.VpcConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcConfigsResponseBodyVpcConfigList struct {
	// example:
	//
	// VPC配置信息描述
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	// example:
	//
	// 05AHW
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// my-test-vpc-config-info
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-2zeid0dd7bhkynxqpaly
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 1756954374266
	SortPriority *int64 `json:"SortPriority,omitempty" xml:"SortPriority,omitempty"`
	// example:
	//
	// vsw-2zes82msebuel5lbw0w1r
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	VpcCidr *string `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
	// example:
	//
	// vpc-bp11w3qellkjwnhqxzmt2
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcConfigsResponseBodyVpcConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListVpcConfigsResponseBodyVpcConfigList) GoString() string {
	return s.String()
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetConfigDescription() *string {
	return s.ConfigDescription
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetSortPriority() *int64 {
	return s.SortPriority
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetVpcCidr() *string {
	return s.VpcCidr
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetConfigDescription(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.ConfigDescription = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetConfigId(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.ConfigId = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetConfigName(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.ConfigName = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetRegionId(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.RegionId = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetSecurityGroupId(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetSortPriority(v int64) *ListVpcConfigsResponseBodyVpcConfigList {
	s.SortPriority = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetVSwitchId(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetVpcCidr(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.VpcCidr = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) SetVpcId(v string) *ListVpcConfigsResponseBodyVpcConfigList {
	s.VpcId = &v
	return s
}

func (s *ListVpcConfigsResponseBodyVpcConfigList) Validate() error {
	return dara.Validate(s)
}
