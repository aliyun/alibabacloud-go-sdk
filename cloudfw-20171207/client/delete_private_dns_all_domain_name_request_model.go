// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsAllDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DeletePrivateDnsAllDomainNameRequest
	GetAccessInstanceId() *string
	SetRegionNo(v string) *DeletePrivateDnsAllDomainNameRequest
	GetRegionNo() *string
}

type DeletePrivateDnsAllDomainNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DeletePrivateDnsAllDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsAllDomainNameRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsAllDomainNameRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DeletePrivateDnsAllDomainNameRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DeletePrivateDnsAllDomainNameRequest) SetAccessInstanceId(v string) *DeletePrivateDnsAllDomainNameRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DeletePrivateDnsAllDomainNameRequest) SetRegionNo(v string) *DeletePrivateDnsAllDomainNameRequest {
	s.RegionNo = &v
	return s
}

func (s *DeletePrivateDnsAllDomainNameRequest) Validate() error {
	return dara.Validate(s)
}
