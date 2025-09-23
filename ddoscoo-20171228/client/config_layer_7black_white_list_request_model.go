// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7BlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackList(v []*string) *ConfigLayer7BlackWhiteListRequest
	GetBlackList() []*string
	SetDomain(v string) *ConfigLayer7BlackWhiteListRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ConfigLayer7BlackWhiteListRequest
	GetResourceGroupId() *string
	SetWhiteList(v []*string) *ConfigLayer7BlackWhiteListRequest
	GetWhiteList() []*string
}

type ConfigLayer7BlackWhiteListRequest struct {
	// example:
	//
	// 1.1.1.1
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigLayer7BlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListRequest) GetBlackList() []*string {
	return s.BlackList
}

func (s *ConfigLayer7BlackWhiteListRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigLayer7BlackWhiteListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigLayer7BlackWhiteListRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *ConfigLayer7BlackWhiteListRequest) SetBlackList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.BlackList = v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetDomain(v string) *ConfigLayer7BlackWhiteListRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetResourceGroupId(v string) *ConfigLayer7BlackWhiteListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetWhiteList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.WhiteList = v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
