// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnStagingIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIPV4s(v *DescribeDcdnStagingIpResponseBodyIPV4s) *DescribeDcdnStagingIpResponseBody
	GetIPV4s() *DescribeDcdnStagingIpResponseBodyIPV4s
	SetRequestId(v string) *DescribeDcdnStagingIpResponseBody
	GetRequestId() *string
}

type DescribeDcdnStagingIpResponseBody struct {
	IPV4s *DescribeDcdnStagingIpResponseBodyIPV4s `json:"IPV4s,omitempty" xml:"IPV4s,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1B9E0E83-24AC-49F4-9EE0-BF5EB03E8381
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnStagingIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnStagingIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponseBody) GetIPV4s() *DescribeDcdnStagingIpResponseBodyIPV4s {
	return s.IPV4s
}

func (s *DescribeDcdnStagingIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnStagingIpResponseBody) SetIPV4s(v *DescribeDcdnStagingIpResponseBodyIPV4s) *DescribeDcdnStagingIpResponseBody {
	s.IPV4s = v
	return s
}

func (s *DescribeDcdnStagingIpResponseBody) SetRequestId(v string) *DescribeDcdnStagingIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnStagingIpResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnStagingIpResponseBodyIPV4s struct {
	IPV4 []*string `json:"IPV4,omitempty" xml:"IPV4,omitempty" type:"Repeated"`
}

func (s DescribeDcdnStagingIpResponseBodyIPV4s) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnStagingIpResponseBodyIPV4s) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponseBodyIPV4s) GetIPV4() []*string {
	return s.IPV4
}

func (s *DescribeDcdnStagingIpResponseBodyIPV4s) SetIPV4(v []*string) *DescribeDcdnStagingIpResponseBodyIPV4s {
	s.IPV4 = v
	return s
}

func (s *DescribeDcdnStagingIpResponseBodyIPV4s) Validate() error {
	return dara.Validate(s)
}
