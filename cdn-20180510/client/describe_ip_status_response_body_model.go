// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpStatus(v []*DescribeIpStatusResponseBodyIpStatus) *DescribeIpStatusResponseBody
	GetIpStatus() []*DescribeIpStatusResponseBodyIpStatus
	SetRequestId(v string) *DescribeIpStatusResponseBody
	GetRequestId() *string
}

type DescribeIpStatusResponseBody struct {
	// The status of the IP addresses of the POPs.
	IpStatus []*DescribeIpStatusResponseBodyIpStatus `json:"IpStatus,omitempty" xml:"IpStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpStatusResponseBody) GetIpStatus() []*DescribeIpStatusResponseBodyIpStatus {
	return s.IpStatus
}

func (s *DescribeIpStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpStatusResponseBody) SetIpStatus(v []*DescribeIpStatusResponseBodyIpStatus) *DescribeIpStatusResponseBody {
	s.IpStatus = v
	return s
}

func (s *DescribeIpStatusResponseBody) SetRequestId(v string) *DescribeIpStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpStatusResponseBody) Validate() error {
	if s.IpStatus != nil {
		for _, item := range s.IpStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpStatusResponseBodyIpStatus struct {
	// The IP address of the POP.
	//
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// The status.
	//
	// 	- **nonali**: not an Alibaba Cloud CDN POP
	//
	// 	- **normal**: an available Alibaba Cloud CDN POP
	//
	// 	- **abnormal**: an unavailable Alibaba Cloud CDN POP
	//
	// example:
	//
	// abnormal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeIpStatusResponseBodyIpStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpStatusResponseBodyIpStatus) GoString() string {
	return s.String()
}

func (s *DescribeIpStatusResponseBodyIpStatus) GetIp() *string {
	return s.Ip
}

func (s *DescribeIpStatusResponseBodyIpStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpStatusResponseBodyIpStatus) SetIp(v string) *DescribeIpStatusResponseBodyIpStatus {
	s.Ip = &v
	return s
}

func (s *DescribeIpStatusResponseBodyIpStatus) SetStatus(v string) *DescribeIpStatusResponseBodyIpStatus {
	s.Status = &v
	return s
}

func (s *DescribeIpStatusResponseBodyIpStatus) Validate() error {
	return dara.Validate(s)
}
