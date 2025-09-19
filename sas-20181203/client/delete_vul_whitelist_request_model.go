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
	SetWhitelist(v string) *DeleteVulWhitelistRequest
	GetWhitelist() *string
}

type DeleteVulWhitelistRequest struct {
	// The ID of the whitelist.
	//
	// >  To delete a vulnerability whitelist, you must provide the ID of the whitelist. You can call the [DescribeVulWhitelist](~~DescribeVulWhitelist~~) operation to query the IDs of whitelists.
	//
	// example:
	//
	// 131231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the whitelist. The value is a JSON string that contains the following fields:
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
	//     	- **emg**: urgent vulnerability
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

func (s DeleteVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteVulWhitelistRequest) GetId() *string {
	return s.Id
}

func (s *DeleteVulWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *DeleteVulWhitelistRequest) SetId(v string) *DeleteVulWhitelistRequest {
	s.Id = &v
	return s
}

func (s *DeleteVulWhitelistRequest) SetWhitelist(v string) *DeleteVulWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *DeleteVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
