// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeVulWhitelistResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeVulWhitelistResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVulWhitelistResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVulWhitelistResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVulWhitelistResponseBody
	GetTotalCount() *int32
	SetVulWhitelists(v []*DescribeVulWhitelistResponseBodyVulWhitelists) *DescribeVulWhitelistResponseBody
	GetVulWhitelists() []*DescribeVulWhitelistResponseBodyVulWhitelists
}

type DescribeVulWhitelistResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 74F97EF7-B543-43FD-A4E9-18456731F9C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array that consists of information about the whitelist of vulnerabilities.
	VulWhitelists []*DescribeVulWhitelistResponseBodyVulWhitelists `json:"VulWhitelists,omitempty" xml:"VulWhitelists,omitempty" type:"Repeated"`
}

func (s DescribeVulWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVulWhitelistResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulWhitelistResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulWhitelistResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulWhitelistResponseBody) GetVulWhitelists() []*DescribeVulWhitelistResponseBodyVulWhitelists {
	return s.VulWhitelists
}

func (s *DescribeVulWhitelistResponseBody) SetCount(v int32) *DescribeVulWhitelistResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetCurrentPage(v int32) *DescribeVulWhitelistResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetPageSize(v int32) *DescribeVulWhitelistResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetRequestId(v string) *DescribeVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetTotalCount(v int32) *DescribeVulWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetVulWhitelists(v []*DescribeVulWhitelistResponseBodyVulWhitelists) *DescribeVulWhitelistResponseBody {
	s.VulWhitelists = v
	return s
}

func (s *DescribeVulWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVulWhitelistResponseBodyVulWhitelists struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2017:3263: curl security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1275
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20173263
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The reason why the vulnerability is added to the whitelist.
	//
	// example:
	//
	// ignore
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The application scope of the rule. The value is a JSON string that contains the following fields:
	//
	// 	- **type**: the type of the assets to which the rule is applied. Valid values:
	//
	//     	- **Uuid**: server
	//
	//     	- **GroupId**: server group
	//
	// 	- **groupIds**: the ID of the server group
	//
	// 	- **uuids**: the UUID of the server
	//
	// > If this field is empty, the rule is applied to all assets.
	//
	// example:
	//
	// {"type":"GroupId","groupIds":[916****],"uuids":[]}
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
	// The type of the vulnerability.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The information about the vulnerability. The value of this parameter is in the JSON format. Valid values:
	//
	// 	- **name**: the name of the vulnerability
	//
	// 	- **type**: the type of the vulnerability.
	//
	// 	- **aliasName**: the alias of the vulnerability
	//
	// example:
	//
	// [{
	//
	// "name":"oval:com.redhat.rhsa:def:20173263",
	//
	// "type":"cve",
	//
	// "aliasName":"RHSA-2017:3263: curl security update"
	//
	// }]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetId() *string {
	return s.Id
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetName() *string {
	return s.Name
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetReason() *string {
	return s.Reason
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetTargetInfo() *string {
	return s.TargetInfo
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetType() *string {
	return s.Type
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) GetWhitelist() *string {
	return s.Whitelist
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetAliasName(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.AliasName = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetId(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Id = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetName(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Name = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetReason(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Reason = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetTargetInfo(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.TargetInfo = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetType(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Type = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetWhitelist(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Whitelist = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) Validate() error {
	return dara.Validate(s)
}
