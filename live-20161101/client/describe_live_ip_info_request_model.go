// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveIpInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIP(v string) *DescribeLiveIpInfoRequest
	GetIP() *string
	SetOwnerId(v int64) *DescribeLiveIpInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveIpInfoRequest
	GetRegionId() *string
}

type DescribeLiveIpInfoRequest struct {
	// The IP address that you want to check. You can specify only one IP address in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	IP       *string `json:"IP,omitempty" xml:"IP,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveIpInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveIpInfoRequest) GetIP() *string {
	return s.IP
}

func (s *DescribeLiveIpInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveIpInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveIpInfoRequest) SetIP(v string) *DescribeLiveIpInfoRequest {
	s.IP = &v
	return s
}

func (s *DescribeLiveIpInfoRequest) SetOwnerId(v int64) *DescribeLiveIpInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveIpInfoRequest) SetRegionId(v string) *DescribeLiveIpInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveIpInfoRequest) Validate() error {
	return dara.Validate(s)
}
