// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DeletePrivateDnsDomainNameRequest
	GetAccessInstanceId() *string
	SetDomainNameList(v []*string) *DeletePrivateDnsDomainNameRequest
	GetDomainNameList() []*string
	SetRegionNo(v string) *DeletePrivateDnsDomainNameRequest
	GetRegionNo() *string
}

type DeletePrivateDnsDomainNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// This parameter is required.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DeletePrivateDnsDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsDomainNameRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsDomainNameRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DeletePrivateDnsDomainNameRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DeletePrivateDnsDomainNameRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DeletePrivateDnsDomainNameRequest) SetAccessInstanceId(v string) *DeletePrivateDnsDomainNameRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DeletePrivateDnsDomainNameRequest) SetDomainNameList(v []*string) *DeletePrivateDnsDomainNameRequest {
	s.DomainNameList = v
	return s
}

func (s *DeletePrivateDnsDomainNameRequest) SetRegionNo(v string) *DeletePrivateDnsDomainNameRequest {
	s.RegionNo = &v
	return s
}

func (s *DeletePrivateDnsDomainNameRequest) Validate() error {
	return dara.Validate(s)
}
