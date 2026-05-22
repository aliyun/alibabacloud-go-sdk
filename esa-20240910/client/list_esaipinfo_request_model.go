// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListESAIPInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVipInfo(v string) *ListESAIPInfoRequest
	GetVipInfo() *string
}

type ListESAIPInfoRequest struct {
	// You can enter IPv4 or IPv6 addresses. Separate multiple IP addresses with commas (,). You can enter up to 20 IP addresses at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.0.24,2408:8740:41FF:2:23::7FE,0.0.0.0,abcd
	VipInfo *string `json:"VipInfo,omitempty" xml:"VipInfo,omitempty"`
}

func (s ListESAIPInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListESAIPInfoRequest) GoString() string {
	return s.String()
}

func (s *ListESAIPInfoRequest) GetVipInfo() *string {
	return s.VipInfo
}

func (s *ListESAIPInfoRequest) SetVipInfo(v string) *ListESAIPInfoRequest {
	s.VipInfo = &v
	return s
}

func (s *ListESAIPInfoRequest) Validate() error {
	return dara.Validate(s)
}
