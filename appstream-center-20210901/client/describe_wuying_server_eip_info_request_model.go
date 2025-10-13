// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWuyingServerEipInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsp(v string) *DescribeWuyingServerEipInfoRequest
	GetIsp() *string
	SetWuyingServerId(v string) *DescribeWuyingServerEipInfoRequest
	GetWuyingServerId() *string
}

type DescribeWuyingServerEipInfoRequest struct {
	// example:
	//
	// ChinaTelecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s DescribeWuyingServerEipInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerEipInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerEipInfoRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeWuyingServerEipInfoRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *DescribeWuyingServerEipInfoRequest) SetIsp(v string) *DescribeWuyingServerEipInfoRequest {
	s.Isp = &v
	return s
}

func (s *DescribeWuyingServerEipInfoRequest) SetWuyingServerId(v string) *DescribeWuyingServerEipInfoRequest {
	s.WuyingServerId = &v
	return s
}

func (s *DescribeWuyingServerEipInfoRequest) Validate() error {
	return dara.Validate(s)
}
