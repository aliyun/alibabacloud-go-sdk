// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCreateVulWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyCreateVulWhitelistRequest
	GetClientToken() *string
	SetReason(v string) *ModifyCreateVulWhitelistRequest
	GetReason() *string
	SetResourceDirectoryAccountId(v int64) *ModifyCreateVulWhitelistRequest
	GetResourceDirectoryAccountId() *int64
	SetTargetInfo(v string) *ModifyCreateVulWhitelistRequest
	GetTargetInfo() *string
	SetWhitelist(v string) *ModifyCreateVulWhitelistRequest
	GetWhitelist() *string
}

type ModifyCreateVulWhitelistRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The reason for adding the vulnerability to the whitelist.
	//
	// example:
	//
	// This vulnerability is not harmful
	Reason                     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The scope in which the whitelist takes effect. The value is a JSON string that contains the following fields:
	//
	// - **type**: The scope type. Valid values:
	//
	//     - **GroupId**: server group
	//
	//     - **Uuid**: host asset
	//
	// - **uuids**: The collection of host asset UUIDs. The field type is String.
	//
	// - **groupIds**: The collection of server group IDs. The field type is Long.
	//
	// > If this parameter is left empty, the whitelist takes effect on all hosts. If **type*	- is set to **GroupId**, **groupIds*	- cannot be empty. If **type*	- is set to **Uuid**, **uuids*	- cannot be empty.
	//
	// example:
	//
	// {"type":"Uuid","uuids":["b31a708f-5fea-426e-bebe-a7b0893****","1f749687-3b5d-4e11-8140-d964673****"],"groupIds":[]}
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
	// The information about the vulnerability to add to the whitelist. The value is a JSON string that contains the following fields:
	//
	// - **Status**: The vulnerability status.
	//
	// - **GmtLast**: The timestamp when the vulnerability was last detected. Unit: milliseconds.
	//
	// - **LaterCount**: The number of medium-priority vulnerabilities.
	//
	// - **AsapCount**: The number of high-priority vulnerabilities.
	//
	// - **Name**: The vulnerability name.
	//
	// - **Type**: The vulnerability type. Valid values:
	//
	//     - **cve**: Linux software vulnerability
	//
	//     - **sys**: Windows system vulnerability
	//
	//     - **cms**: Web-CMS vulnerability
	//
	//     - **app**: application vulnerability
	//
	//     - **emg**: emergency vulnerability
	//
	// - **Related**: The CVE ID of the vulnerability.
	//
	// - **HandledCount**: The number of handled vulnerabilities.
	//
	// - **AliasName**: The alias of the vulnerability.
	//
	// - **RuleModifyTime**: The time when the vulnerability was last published.
	//
	// - **NntfCount**: The number of low-priority vulnerabilities.
	//
	// - **TotalFixCount**: The total number of fixed vulnerabilities.
	//
	// - **Tags**: The vulnerability tags.
	//
	// > You can call the [DescribeGroupedVul](~~DescribeGroupedVul~~) operation to obtain the vulnerability information to add to the whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Status":0,"GmtLast":1662281929000,"LaterCount":0,"AsapCount":1,"Name":"oval:com.redhat.rhsa:def:20172836","Type":"cve","Related":"CVE-2017-14491,CVE-2017-14492,CVE-2017-14493,CVE-2017-14494,CVE-2017-14495,CVE-2017-14496","HandledCount":1,"AliasName":"RHSA-2017:2836-Critical: dnsmasq security update","RuleModifyTime":1535542395000,"NntfCount":0,"TotalFixCount":196668,"Tags":"Exploit Exists,Code Execution"},{"Status":0,"GmtLast":1662281933000,"LaterCount":0,"AsapCount":1,"Name":"oval:com.redhat.rhsa:def:20173075","Type":"cve","Related":"CVE-2017-13089,CVE-2017-13090","HandledCount":1,"AliasName":"RHSA-2017:3075-Important: wget security update","RuleModifyTime":1551432867000,"NntfCount":0,"TotalFixCount":369136,"Tags":"Code Execution"}]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ModifyCreateVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCreateVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyCreateVulWhitelistRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyCreateVulWhitelistRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyCreateVulWhitelistRequest) GetTargetInfo() *string {
	return s.TargetInfo
}

func (s *ModifyCreateVulWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *ModifyCreateVulWhitelistRequest) SetClientToken(v string) *ModifyCreateVulWhitelistRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetReason(v string) *ModifyCreateVulWhitelistRequest {
	s.Reason = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetResourceDirectoryAccountId(v int64) *ModifyCreateVulWhitelistRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetTargetInfo(v string) *ModifyCreateVulWhitelistRequest {
	s.TargetInfo = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetWhitelist(v string) *ModifyCreateVulWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
