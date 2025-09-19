// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVulWhitelistResponseBody
	GetRequestId() *string
	SetVulWhitelist(v *GetVulWhitelistResponseBodyVulWhitelist) *GetVulWhitelistResponseBody
	GetVulWhitelist() *GetVulWhitelistResponseBodyVulWhitelist
}

type GetVulWhitelistResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the whitelist.
	VulWhitelist *GetVulWhitelistResponseBodyVulWhitelist `json:"VulWhitelist,omitempty" xml:"VulWhitelist,omitempty" type:"Struct"`
}

func (s GetVulWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulWhitelistResponseBody) GetVulWhitelist() *GetVulWhitelistResponseBodyVulWhitelist {
	return s.VulWhitelist
}

func (s *GetVulWhitelistResponseBody) SetRequestId(v string) *GetVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulWhitelistResponseBody) SetVulWhitelist(v *GetVulWhitelistResponseBodyVulWhitelist) *GetVulWhitelistResponseBody {
	s.VulWhitelist = v
	return s
}

func (s *GetVulWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVulWhitelistResponseBodyVulWhitelist struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2017:3263: curl security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the whitelist.
	//
	// example:
	//
	// 1275
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// Ignore
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
	// > If this parameter is empty, the rule is applied to all types of assets.
	//
	// example:
	//
	// {
	//
	//       "type": "GroupId",
	//
	//       "uuids": [],
	//
	//       "groupIds": [
	//
	//             10782678
	//
	//       ]
	//
	// }
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the vulnerability.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The information about the vulnerability that is added to the whitelist. The value is a JSON string that contains the following fields:
	//
	// 	- **Name**: the name of the vulnerability.
	//
	// 	- **Type**: the type of the vulnerability. Valid values:
	//
	//     	- **cve**: Linux software vulnerability
	//
	//     	- **sys**: Windows system vulnerability
	//
	//     	- **cms**: Web-CMS vulnerability
	//
	//     	- **app**: application vulnerability
	//
	//     	- **emg**: urgent vulnerabilities
	//
	// 	- **AliasName**: the alias of the vulnerability.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "Name": "oval:com.redhat.rhsa:def:20173263",
	//
	//             "Type": "cve",
	//
	//             "AliasName": "RHSA-2017:3263: curl security update"
	//
	//       }
	//
	// ]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s GetVulWhitelistResponseBodyVulWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetVulWhitelistResponseBodyVulWhitelist) GoString() string {
	return s.String()
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetAliasName() *string {
	return s.AliasName
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetId() *int64 {
	return s.Id
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetName() *string {
	return s.Name
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetReason() *string {
	return s.Reason
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetTarget() *string {
	return s.Target
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetType() *string {
	return s.Type
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) GetWhitelist() *string {
	return s.Whitelist
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetAliasName(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.AliasName = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetId(v int64) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Id = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetName(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Name = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetReason(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Reason = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetTarget(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Target = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetType(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Type = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) SetWhitelist(v string) *GetVulWhitelistResponseBodyVulWhitelist {
	s.Whitelist = &v
	return s
}

func (s *GetVulWhitelistResponseBodyVulWhitelist) Validate() error {
	return dara.Validate(s)
}
