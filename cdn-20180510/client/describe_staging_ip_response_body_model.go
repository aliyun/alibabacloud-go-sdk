// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStagingIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIPV4s(v *DescribeStagingIpResponseBodyIPV4s) *DescribeStagingIpResponseBody
	GetIPV4s() *DescribeStagingIpResponseBodyIPV4s
	SetRequestId(v string) *DescribeStagingIpResponseBody
	GetRequestId() *string
}

type DescribeStagingIpResponseBody struct {
	// IPv4 addresses.
	IPV4s *DescribeStagingIpResponseBodyIPV4s `json:"IPV4s,omitempty" xml:"IPV4s,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1B9E0E83-24AC-49F4-9EE0-BF5EB03E8381
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStagingIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStagingIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponseBody) GetIPV4s() *DescribeStagingIpResponseBodyIPV4s {
	return s.IPV4s
}

func (s *DescribeStagingIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStagingIpResponseBody) SetIPV4s(v *DescribeStagingIpResponseBodyIPV4s) *DescribeStagingIpResponseBody {
	s.IPV4s = v
	return s
}

func (s *DescribeStagingIpResponseBody) SetRequestId(v string) *DescribeStagingIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStagingIpResponseBody) Validate() error {
	if s.IPV4s != nil {
		if err := s.IPV4s.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStagingIpResponseBodyIPV4s struct {
	IPV4 []*string `json:"IPV4,omitempty" xml:"IPV4,omitempty" type:"Repeated"`
}

func (s DescribeStagingIpResponseBodyIPV4s) String() string {
	return dara.Prettify(s)
}

func (s DescribeStagingIpResponseBodyIPV4s) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponseBodyIPV4s) GetIPV4() []*string {
	return s.IPV4
}

func (s *DescribeStagingIpResponseBodyIPV4s) SetIPV4(v []*string) *DescribeStagingIpResponseBodyIPV4s {
	s.IPV4 = v
	return s
}

func (s *DescribeStagingIpResponseBodyIPV4s) Validate() error {
	return dara.Validate(s)
}
