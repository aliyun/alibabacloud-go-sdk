// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIP(v string) *DescribeDcdnIpInfoRequest
	GetIP() *string
}

type DescribeDcdnIpInfoRequest struct {
	// The IP address. You can specify only one IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.10.10
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s DescribeDcdnIpInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoRequest) GetIP() *string {
	return s.IP
}

func (s *DescribeDcdnIpInfoRequest) SetIP(v string) *DescribeDcdnIpInfoRequest {
	s.IP = &v
	return s
}

func (s *DescribeDcdnIpInfoRequest) Validate() error {
	return dara.Validate(s)
}
