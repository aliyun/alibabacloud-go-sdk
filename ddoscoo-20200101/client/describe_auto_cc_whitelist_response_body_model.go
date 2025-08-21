// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCcWhitelist(v []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) *DescribeAutoCcWhitelistResponseBody
	GetAutoCcWhitelist() []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist
	SetRequestId(v string) *DescribeAutoCcWhitelistResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAutoCcWhitelistResponseBody
	GetTotalCount() *int64
}

type DescribeAutoCcWhitelistResponseBody struct {
	// An array that consists of details of the IP address in the whitelist of the instance.
	AutoCcWhitelist []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist `json:"AutoCcWhitelist,omitempty" xml:"AutoCcWhitelist,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F09D085E-5E0F-4FF2-B32E-F4A644049162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned IP addresses in the whitelist.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoCcWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponseBody) GetAutoCcWhitelist() []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	return s.AutoCcWhitelist
}

func (s *DescribeAutoCcWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoCcWhitelistResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAutoCcWhitelistResponseBody) SetAutoCcWhitelist(v []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) *DescribeAutoCcWhitelistResponseBody {
	s.AutoCcWhitelist = v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBody) SetRequestId(v string) *DescribeAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBody) SetTotalCount(v int64) *DescribeAutoCcWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist struct {
	// The IP address of the instance.
	//
	// example:
	//
	// 203.***.***.117
	DestIp *string `json:"DestIp,omitempty" xml:"DestIp,omitempty"`
	// The validity period of the IP address in the whitelist. Unit: seconds. **0*	- indicates that the IP address in the whitelist never expires.
	//
	// example:
	//
	// 0
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP addresses that is contained in the IP address whitelist.
	//
	// example:
	//
	// 2.2.2.2
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The mode of how an IP address is added to the whitelist. Valid values:
	//
	// 	- **manual**: manually added
	//
	// 	- **auto**: automatically added
	//
	// example:
	//
	// manual
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GetDestIp() *string {
	return s.DestIp
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GetType() *string {
	return s.Type
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetDestIp(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.DestIp = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetEndTime(v int64) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetSourceIp(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetType(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.Type = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) Validate() error {
	return dara.Validate(s)
}
