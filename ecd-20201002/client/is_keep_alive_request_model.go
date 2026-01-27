// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsKeepAliveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *IsKeepAliveRequest
	GetClientId() *string
	SetOfficeSiteId(v string) *IsKeepAliveRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *IsKeepAliveRequest
	GetRegionId() *string
}

type IsKeepAliveRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s IsKeepAliveRequest) String() string {
	return dara.Prettify(s)
}

func (s IsKeepAliveRequest) GoString() string {
	return s.String()
}

func (s *IsKeepAliveRequest) GetClientId() *string {
	return s.ClientId
}

func (s *IsKeepAliveRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *IsKeepAliveRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *IsKeepAliveRequest) SetClientId(v string) *IsKeepAliveRequest {
	s.ClientId = &v
	return s
}

func (s *IsKeepAliveRequest) SetOfficeSiteId(v string) *IsKeepAliveRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveRequest) SetRegionId(v string) *IsKeepAliveRequest {
	s.RegionId = &v
	return s
}

func (s *IsKeepAliveRequest) Validate() error {
	return dara.Validate(s)
}
