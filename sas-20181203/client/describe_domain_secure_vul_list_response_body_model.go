// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureVulListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainSecureVulListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDomainSecureVulListResponseBody
	GetTotalCount() *int32
	SetVulList(v []*DescribeDomainSecureVulListResponseBodyVulList) *DescribeDomainSecureVulListResponseBody
	GetVulList() []*DescribeDomainSecureVulListResponseBodyVulList
}

type DescribeDomainSecureVulListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The domain name-related vulnerabilities.
	VulList []*DescribeDomainSecureVulListResponseBodyVulList `json:"VulList,omitempty" xml:"VulList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSecureVulListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureVulListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecureVulListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDomainSecureVulListResponseBody) GetVulList() []*DescribeDomainSecureVulListResponseBodyVulList {
	return s.VulList
}

func (s *DescribeDomainSecureVulListResponseBody) SetRequestId(v string) *DescribeDomainSecureVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBody) SetTotalCount(v int32) *DescribeDomainSecureVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBody) SetVulList(v []*DescribeDomainSecureVulListResponseBodyVulList) *DescribeDomainSecureVulListResponseBody {
	s.VulList = v
	return s
}

func (s *DescribeDomainSecureVulListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSecureVulListResponseBodyVulList struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// CESA-2023:3555: python Security Update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The number of the vulnerabilities that have the **high*	- priority.
	//
	// example:
	//
	// 50
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// The timestamp when the vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1639371446000
	GmtLast *int64 `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	// The number of handled vulnerabilities.
	//
	// example:
	//
	// 33
	HandledCount *int32 `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	// The number of the vulnerabilities that have the **medium*	- priority.
	//
	// example:
	//
	// 30
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20170574
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the vulnerabilities that have the **low*	- priority.
	//
	// example:
	//
	// 20
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// The tag that is added to the vulnerability. Valid values:
	//
	// 	- Restart required
	//
	// 	- Remote utilization
	//
	// 	- EXP exists
	//
	// 	- Available
	//
	// 	- Elevation of Privilege
	//
	// 	- Code Execution
	//
	// example:
	//
	// Code Execution
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the vulnerability. Default value: cve. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// 	- **app**: application vulnerability that is detected by network scanning.
	//
	// 	- **sca**: application vulnerability that is detected by using software component analysis.
	//
	// example:
	//
	// app
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainSecureVulListResponseBodyVulList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureVulListResponseBodyVulList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetAsapCount() *int32 {
	return s.AsapCount
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetGmtLast() *int64 {
	return s.GmtLast
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetHandledCount() *int32 {
	return s.HandledCount
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetLaterCount() *int32 {
	return s.LaterCount
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetName() *string {
	return s.Name
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetNntfCount() *int32 {
	return s.NntfCount
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetTags() *string {
	return s.Tags
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) GetType() *string {
	return s.Type
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetAliasName(v string) *DescribeDomainSecureVulListResponseBodyVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetAsapCount(v int32) *DescribeDomainSecureVulListResponseBodyVulList {
	s.AsapCount = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetGmtLast(v int64) *DescribeDomainSecureVulListResponseBodyVulList {
	s.GmtLast = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetHandledCount(v int32) *DescribeDomainSecureVulListResponseBodyVulList {
	s.HandledCount = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetLaterCount(v int32) *DescribeDomainSecureVulListResponseBodyVulList {
	s.LaterCount = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetName(v string) *DescribeDomainSecureVulListResponseBodyVulList {
	s.Name = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetNntfCount(v int32) *DescribeDomainSecureVulListResponseBodyVulList {
	s.NntfCount = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetTags(v string) *DescribeDomainSecureVulListResponseBodyVulList {
	s.Tags = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) SetType(v string) *DescribeDomainSecureVulListResponseBodyVulList {
	s.Type = &v
	return s
}

func (s *DescribeDomainSecureVulListResponseBodyVulList) Validate() error {
	return dara.Validate(s)
}
