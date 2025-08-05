// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateDnsDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *AddPrivateDnsDomainNameRequest
	GetAccessInstanceId() *string
	SetDomainNameList(v []*string) *AddPrivateDnsDomainNameRequest
	GetDomainNameList() []*string
	SetRegionNo(v string) *AddPrivateDnsDomainNameRequest
	GetRegionNo() *string
}

type AddPrivateDnsDomainNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cfw-xxx
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// This parameter is required.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s AddPrivateDnsDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateDnsDomainNameRequest) GoString() string {
	return s.String()
}

func (s *AddPrivateDnsDomainNameRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *AddPrivateDnsDomainNameRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *AddPrivateDnsDomainNameRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *AddPrivateDnsDomainNameRequest) SetAccessInstanceId(v string) *AddPrivateDnsDomainNameRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *AddPrivateDnsDomainNameRequest) SetDomainNameList(v []*string) *AddPrivateDnsDomainNameRequest {
	s.DomainNameList = v
	return s
}

func (s *AddPrivateDnsDomainNameRequest) SetRegionNo(v string) *AddPrivateDnsDomainNameRequest {
	s.RegionNo = &v
	return s
}

func (s *AddPrivateDnsDomainNameRequest) Validate() error {
	return dara.Validate(s)
}
