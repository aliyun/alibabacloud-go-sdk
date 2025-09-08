// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventWhiteruleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *PostEventWhiteruleListRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *PostEventWhiteruleListRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostEventWhiteruleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostEventWhiteruleListRequest
	GetRoleType() *int32
	SetWhiteruleList(v string) *PostEventWhiteruleListRequest
	GetWhiteruleList() *string
}

type PostEventWhiteruleListRequest struct {
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The alert whitelist rule. The value is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "alertName": "webshell",
	//
	//             "alertNameId": "webshell",
	//
	//             "alertType": "command",
	//
	//             "alertTypeId": "command",
	//
	//             "expression": {
	//
	//                   "status": 1,
	//
	//                   "conditions": [
	//
	//                         {
	//
	//                               "isNot": false,
	//
	//                               "left": {
	//
	//                                     "value": "file_path"
	//
	//                               },
	//
	//                               "operator": "gt",
	//
	//                               "right": {
	//
	//                                     "value": "cp"
	//
	//                               }
	//
	//                         }
	//
	//                   ]
	//
	//             }
	//
	//       }
	//
	// ]
	WhiteruleList *string `json:"WhiteruleList,omitempty" xml:"WhiteruleList,omitempty"`
}

func (s PostEventWhiteruleListRequest) String() string {
	return dara.Prettify(s)
}

func (s PostEventWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *PostEventWhiteruleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostEventWhiteruleListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostEventWhiteruleListRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostEventWhiteruleListRequest) GetWhiteruleList() *string {
	return s.WhiteruleList
}

func (s *PostEventWhiteruleListRequest) SetIncidentUuid(v string) *PostEventWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRegionId(v string) *PostEventWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRoleFor(v int64) *PostEventWhiteruleListRequest {
	s.RoleFor = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRoleType(v int32) *PostEventWhiteruleListRequest {
	s.RoleType = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetWhiteruleList(v string) *PostEventWhiteruleListRequest {
	s.WhiteruleList = &v
	return s
}

func (s *PostEventWhiteruleListRequest) Validate() error {
	return dara.Validate(s)
}
