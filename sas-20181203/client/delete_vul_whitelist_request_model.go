// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteVulWhitelistRequest
	GetId() *string
	SetResourceDirectoryAccountId(v int64) *DeleteVulWhitelistRequest
	GetResourceDirectoryAccountId() *int64
	SetWhitelist(v string) *DeleteVulWhitelistRequest
	GetWhitelist() *string
}

type DeleteVulWhitelistRequest struct {
	// The ID of the vulnerability whitelist.
	//
	// > To delete a vulnerability whitelist, provide the vulnerability whitelist ID. You can obtain this ID by calling the [DescribeVulWhitelist](~~DescribeVulWhitelist~~) operation.
	//
	// example:
	//
	// 131231
	Id                         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The vulnerability whitelist information to delete. The value is a JSON string that contains the following fields:
	//
	// - **Name**: The name of the vulnerability.
	//
	// - **Type**: The type of the vulnerability. Valid values:
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
	// - **AliasName**: The alias of the vulnerability.
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

func (s DeleteVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteVulWhitelistRequest) GetId() *string {
	return s.Id
}

func (s *DeleteVulWhitelistRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DeleteVulWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *DeleteVulWhitelistRequest) SetId(v string) *DeleteVulWhitelistRequest {
	s.Id = &v
	return s
}

func (s *DeleteVulWhitelistRequest) SetResourceDirectoryAccountId(v int64) *DeleteVulWhitelistRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DeleteVulWhitelistRequest) SetWhitelist(v string) *DeleteVulWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *DeleteVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
