// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCcBlacklist(v []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) *DescribeAutoCcBlacklistResponseBody
	GetAutoCcBlacklist() []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist
	SetRequestId(v string) *DescribeAutoCcBlacklistResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAutoCcBlacklistResponseBody
	GetTotalCount() *int64
}

type DescribeAutoCcBlacklistResponseBody struct {
	// An array that consists of the details of the IP addresses in the blacklist of the instance.
	AutoCcBlacklist []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist `json:"AutoCcBlacklist,omitempty" xml:"AutoCcBlacklist,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E78C8472-0B15-42D5-AF22-A32A78818AB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned IP addresses in the blacklist.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoCcBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponseBody) GetAutoCcBlacklist() []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	return s.AutoCcBlacklist
}

func (s *DescribeAutoCcBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoCcBlacklistResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAutoCcBlacklistResponseBody) SetAutoCcBlacklist(v []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) *DescribeAutoCcBlacklistResponseBody {
	s.AutoCcBlacklist = v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBody) SetRequestId(v string) *DescribeAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBody) SetTotalCount(v int64) *DescribeAutoCcBlacklistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBody) Validate() error {
	if s.AutoCcBlacklist != nil {
		for _, item := range s.AutoCcBlacklist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist struct {
	// The IP address of the instance.
	//
	// example:
	//
	// 192.0.XX.XX
	DestIp *string `json:"DestIp,omitempty" xml:"DestIp,omitempty"`
	// The validity period of the IP address in the blacklist. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1584093569
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address in the blacklist.
	//
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The mode of how the IP address is added to the blacklist. Valid values:
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

func (s DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GetDestIp() *string {
	return s.DestIp
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GetType() *string {
	return s.Type
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetDestIp(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.DestIp = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetEndTime(v int64) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetSourceIp(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetType(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.Type = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) Validate() error {
	return dara.Validate(s)
}
