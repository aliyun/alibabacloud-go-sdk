// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigSets(v []*ConfigSetListResponseBodyConfigSets) *ConfigSetListResponseBody
	GetConfigSets() []*ConfigSetListResponseBodyConfigSets
	SetCurrentPage(v int32) *ConfigSetListResponseBody
	GetCurrentPage() *int32
	SetHasMore(v bool) *ConfigSetListResponseBody
	GetHasMore() *bool
	SetPageSize(v int32) *ConfigSetListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ConfigSetListResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ConfigSetListResponseBody
	GetTotalCounts() *int32
}

type ConfigSetListResponseBody struct {
	ConfigSets []*ConfigSetListResponseBodyConfigSets `json:"ConfigSets,omitempty" xml:"ConfigSets,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ConfigSetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetListResponseBody) GetConfigSets() []*ConfigSetListResponseBodyConfigSets {
	return s.ConfigSets
}

func (s *ConfigSetListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ConfigSetListResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ConfigSetListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ConfigSetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetListResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ConfigSetListResponseBody) SetConfigSets(v []*ConfigSetListResponseBodyConfigSets) *ConfigSetListResponseBody {
	s.ConfigSets = v
	return s
}

func (s *ConfigSetListResponseBody) SetCurrentPage(v int32) *ConfigSetListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ConfigSetListResponseBody) SetHasMore(v bool) *ConfigSetListResponseBody {
	s.HasMore = &v
	return s
}

func (s *ConfigSetListResponseBody) SetPageSize(v int32) *ConfigSetListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ConfigSetListResponseBody) SetRequestId(v string) *ConfigSetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetListResponseBody) SetTotalCounts(v int32) *ConfigSetListResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ConfigSetListResponseBody) Validate() error {
	if s.ConfigSets != nil {
		for _, item := range s.ConfigSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConfigSetListResponseBodyConfigSets struct {
	// example:
	//
	// xxx
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	FromAddresses []*string `json:"FromAddresses,omitempty" xml:"FromAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	Id     *string                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	IpPool *ConfigSetListResponseBodyConfigSetsIpPool `json:"IpPool,omitempty" xml:"IpPool,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ConfigSetListResponseBodyConfigSets) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetListResponseBodyConfigSets) GoString() string {
	return s.String()
}

func (s *ConfigSetListResponseBodyConfigSets) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetListResponseBodyConfigSets) GetFromAddresses() []*string {
	return s.FromAddresses
}

func (s *ConfigSetListResponseBodyConfigSets) GetId() *string {
	return s.Id
}

func (s *ConfigSetListResponseBodyConfigSets) GetIpPool() *ConfigSetListResponseBodyConfigSetsIpPool {
	return s.IpPool
}

func (s *ConfigSetListResponseBodyConfigSets) GetName() *string {
	return s.Name
}

func (s *ConfigSetListResponseBodyConfigSets) SetDescription(v string) *ConfigSetListResponseBodyConfigSets {
	s.Description = &v
	return s
}

func (s *ConfigSetListResponseBodyConfigSets) SetFromAddresses(v []*string) *ConfigSetListResponseBodyConfigSets {
	s.FromAddresses = v
	return s
}

func (s *ConfigSetListResponseBodyConfigSets) SetId(v string) *ConfigSetListResponseBodyConfigSets {
	s.Id = &v
	return s
}

func (s *ConfigSetListResponseBodyConfigSets) SetIpPool(v *ConfigSetListResponseBodyConfigSetsIpPool) *ConfigSetListResponseBodyConfigSets {
	s.IpPool = v
	return s
}

func (s *ConfigSetListResponseBodyConfigSets) SetName(v string) *ConfigSetListResponseBodyConfigSets {
	s.Name = &v
	return s
}

func (s *ConfigSetListResponseBodyConfigSets) Validate() error {
	if s.IpPool != nil {
		if err := s.IpPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetListResponseBodyConfigSetsIpPool struct {
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// example:
	//
	// xxx
	IpPoolName *string `json:"IpPoolName,omitempty" xml:"IpPoolName,omitempty"`
}

func (s ConfigSetListResponseBodyConfigSetsIpPool) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetListResponseBodyConfigSetsIpPool) GoString() string {
	return s.String()
}

func (s *ConfigSetListResponseBodyConfigSetsIpPool) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetListResponseBodyConfigSetsIpPool) GetIpPoolName() *string {
	return s.IpPoolName
}

func (s *ConfigSetListResponseBodyConfigSetsIpPool) SetIpPoolId(v string) *ConfigSetListResponseBodyConfigSetsIpPool {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetListResponseBodyConfigSetsIpPool) SetIpPoolName(v string) *ConfigSetListResponseBodyConfigSetsIpPool {
	s.IpPoolName = &v
	return s
}

func (s *ConfigSetListResponseBodyConfigSetsIpPool) Validate() error {
	return dara.Validate(s)
}
