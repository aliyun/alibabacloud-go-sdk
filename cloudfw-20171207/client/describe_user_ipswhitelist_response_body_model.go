// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserIPSWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Whitelists(v []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) *DescribeUserIPSWhitelistResponseBody
	GetIpv6Whitelists() []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists
	SetRequestId(v string) *DescribeUserIPSWhitelistResponseBody
	GetRequestId() *string
	SetWhitelists(v []*DescribeUserIPSWhitelistResponseBodyWhitelists) *DescribeUserIPSWhitelistResponseBody
	GetWhitelists() []*DescribeUserIPSWhitelistResponseBodyWhitelists
}

type DescribeUserIPSWhitelistResponseBody struct {
	Ipv6Whitelists []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists `json:"Ipv6Whitelists,omitempty" xml:"Ipv6Whitelists,omitempty" type:"Repeated"`
	// example:
	//
	// 04F788A5-6A47-5EA9-AC30-CA4DB98AD520
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Whitelists []*DescribeUserIPSWhitelistResponseBodyWhitelists `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s DescribeUserIPSWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBody) GetIpv6Whitelists() []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	return s.Ipv6Whitelists
}

func (s *DescribeUserIPSWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserIPSWhitelistResponseBody) GetWhitelists() []*DescribeUserIPSWhitelistResponseBodyWhitelists {
	return s.Whitelists
}

func (s *DescribeUserIPSWhitelistResponseBody) SetIpv6Whitelists(v []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) *DescribeUserIPSWhitelistResponseBody {
	s.Ipv6Whitelists = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBody) SetRequestId(v string) *DescribeUserIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBody) SetWhitelists(v []*DescribeUserIPSWhitelistResponseBodyWhitelists) *DescribeUserIPSWhitelistResponseBody {
	s.Whitelists = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBody) Validate() error {
	if s.Ipv6Whitelists != nil {
		for _, item := range s.Ipv6Whitelists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Whitelists != nil {
		for _, item := range s.Whitelists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserIPSWhitelistResponseBodyIpv6Whitelists struct {
	// example:
	//
	// 1
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 0
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// example:
	//
	// 2408:400a:81a:7900:a77d:ea36:fcbf:de40/128
	ListValue      *string   `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GetDirection() *int64 {
	return s.Direction
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GetListType() *int64 {
	return s.ListType
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GetListValue() *string {
	return s.ListValue
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GetWhiteListValue() []*string {
	return s.WhiteListValue
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GetWhiteType() *int64 {
	return s.WhiteType
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetDirection(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.Direction = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetListType(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.ListType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetListValue(v string) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetWhiteListValue(v []*string) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetWhiteType(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.WhiteType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) Validate() error {
	return dara.Validate(s)
}

type DescribeUserIPSWhitelistResponseBodyWhitelists struct {
	// example:
	//
	// 1
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 1
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// example:
	//
	// 10.10.200.4/32,10.10.200.25/32
	ListValue      *string   `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeUserIPSWhitelistResponseBodyWhitelists) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBodyWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) GetDirection() *int64 {
	return s.Direction
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) GetListType() *int64 {
	return s.ListType
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) GetListValue() *string {
	return s.ListValue
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) GetWhiteListValue() []*string {
	return s.WhiteListValue
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) GetWhiteType() *int64 {
	return s.WhiteType
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetDirection(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.Direction = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetListType(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.ListType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetListValue(v string) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetWhiteListValue(v []*string) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetWhiteType(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.WhiteType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) Validate() error {
	return dara.Validate(s)
}
