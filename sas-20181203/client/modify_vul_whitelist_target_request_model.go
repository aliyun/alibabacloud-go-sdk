// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulWhitelistTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyVulWhitelistTargetRequest
	GetId() *int64
	SetReason(v string) *ModifyVulWhitelistTargetRequest
	GetReason() *string
	SetSourceIp(v string) *ModifyVulWhitelistTargetRequest
	GetSourceIp() *string
	SetTargetInfo(v string) *ModifyVulWhitelistTargetRequest
	GetTargetInfo() *string
}

type ModifyVulWhitelistTargetRequest struct {
	// The ID of the whitelist.
	//
	// >  You can call the [DescribeVulWhitelist](~~DescribeVulWhitelist~~) operation to query the IDs of whitelists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2533681
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The reason why you add the server to the whitelist.
	//
	// example:
	//
	// 1221
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 42.120.75.150
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The applicable scope of the whitelist. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: the type of the applicable scope. Valid values:
	//
	//     	- **GroupId**: the ID of a server group
	//
	//     	- **Uuid**: the UUID of a server
	//
	// 	- **uuids**: the UUIDs of servers
	//
	// 	- **groupIds**: the IDs of server groups
	//
	// >  If you leave this parameter empty, all servers are added to the whitelist. If you set the **type*	- field to **GroupId**, you must also specify the **groupIds*	- field. If you set the **type*	- field to **Uuid**, you must also specify the **uuids*	- field.
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
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
}

func (s ModifyVulWhitelistTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulWhitelistTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyVulWhitelistTargetRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyVulWhitelistTargetRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyVulWhitelistTargetRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyVulWhitelistTargetRequest) GetTargetInfo() *string {
	return s.TargetInfo
}

func (s *ModifyVulWhitelistTargetRequest) SetId(v int64) *ModifyVulWhitelistTargetRequest {
	s.Id = &v
	return s
}

func (s *ModifyVulWhitelistTargetRequest) SetReason(v string) *ModifyVulWhitelistTargetRequest {
	s.Reason = &v
	return s
}

func (s *ModifyVulWhitelistTargetRequest) SetSourceIp(v string) *ModifyVulWhitelistTargetRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVulWhitelistTargetRequest) SetTargetInfo(v string) *ModifyVulWhitelistTargetRequest {
	s.TargetInfo = &v
	return s
}

func (s *ModifyVulWhitelistTargetRequest) Validate() error {
	return dara.Validate(s)
}
