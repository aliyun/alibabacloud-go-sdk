// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMajorProtectionBlackIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpList(v []*DescribeMajorProtectionBlackIpsResponseBodyIpList) *DescribeMajorProtectionBlackIpsResponseBody
	GetIpList() []*DescribeMajorProtectionBlackIpsResponseBodyIpList
	SetRequestId(v string) *DescribeMajorProtectionBlackIpsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeMajorProtectionBlackIpsResponseBody
	GetTotalCount() *int64
}

type DescribeMajorProtectionBlackIpsResponseBody struct {
	// An array of IP addresses in the IP address blacklist.
	IpList []*DescribeMajorProtectionBlackIpsResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 41631674-EEB0-5B02-BEB4-40A758E9B841
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of IP addresses in the blacklist.
	//
	// example:
	//
	// 63
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) GetIpList() []*DescribeMajorProtectionBlackIpsResponseBodyIpList {
	return s.IpList
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetIpList(v []*DescribeMajorProtectionBlackIpsResponseBodyIpList) *DescribeMajorProtectionBlackIpsResponseBody {
	s.IpList = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetRequestId(v string) *DescribeMajorProtectionBlackIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetTotalCount(v int64) *DescribeMajorProtectionBlackIpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMajorProtectionBlackIpsResponseBodyIpList struct {
	// The description of the IP address in the blacklist.
	//
	// example:
	//
	// test0003asdffas
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If the value of this parameter is **0**, the blacklist is permanently valid.
	//
	// example:
	//
	// 1662603328
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The most recent time when the IP address blacklist was modified.
	//
	// example:
	//
	// 1665456202000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP address in the IP address blacklist.
	//
	// example:
	//
	// 192.0.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	//
	// example:
	//
	// 8508970
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the rule template for major event protection.
	//
	// example:
	//
	// 9684
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetDescription() *string {
	return s.Description
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetIp() *string {
	return s.Ip
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetDescription(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Description = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetExpiredTime(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetGmtModified(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.GmtModified = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetIp(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.TemplateId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) Validate() error {
	return dara.Validate(s)
}
