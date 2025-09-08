// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventDisposeAndWhiteruleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest
	GetEventDispose() *string
	SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest
	GetIncidentUuid() *string
	SetReceiverInfo(v string) *PostEventDisposeAndWhiteruleListRequest
	GetReceiverInfo() *string
	SetRegionId(v string) *PostEventDisposeAndWhiteruleListRequest
	GetRegionId() *string
	SetRemark(v string) *PostEventDisposeAndWhiteruleListRequest
	GetRemark() *string
	SetRoleFor(v int64) *PostEventDisposeAndWhiteruleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostEventDisposeAndWhiteruleListRequest
	GetRoleType() *int32
	SetStatus(v int32) *PostEventDisposeAndWhiteruleListRequest
	GetStatus() *int32
	SetThreatLevel(v string) *PostEventDisposeAndWhiteruleListRequest
	GetThreatLevel() *string
}

type PostEventDisposeAndWhiteruleListRequest struct {
	// The configuration of event handling. The value is a JSON object.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "entityId": "104466118",
	//
	//             "scope": [
	//
	//                   "176618589410****"
	//
	//             ],
	//
	//             "startTime": 1604168946281,
	//
	//             "endTime": 1614168946281
	//
	//       },
	//
	//       {
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "entityId": "104466118",
	//
	//             "scope": [
	//
	//                   {
	//
	//                         "instanceId": "waf-cn-n6w1oy1****",
	//
	//                         "domains": [
	//
	//                               "lmfip.wafqax.***"
	//
	//                         ]
	//
	//                   }
	//
	//             ],
	//
	//             "startTime": 1604168946281,
	//
	//             "endTime": 1614168946281
	//
	//       }
	//
	// ]
	EventDispose *string `json:"EventDispose,omitempty" xml:"EventDispose,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The configuration of the alert recipient. The value is a JSON object.
	//
	// example:
	//
	// {
	//
	//       "messageTitle": "test",
	//
	//       "receiver": "xiaowang",
	//
	//       "channel": "message"
	//
	// }
	ReceiverInfo *string `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty"`
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
	// The remarks of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListRequest) String() string {
	return dara.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetEventDispose() *string {
	return s.EventDispose
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetReceiverInfo() *string {
	return s.ReceiverInfo
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetRemark() *string {
	return s.Remark
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.EventDispose = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetReceiverInfo(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ReceiverInfo = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRegionId(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRemark(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.Remark = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRoleFor(v int64) *PostEventDisposeAndWhiteruleListRequest {
	s.RoleFor = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRoleType(v int32) *PostEventDisposeAndWhiteruleListRequest {
	s.RoleType = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetStatus(v int32) *PostEventDisposeAndWhiteruleListRequest {
	s.Status = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetThreatLevel(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ThreatLevel = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) Validate() error {
	return dara.Validate(s)
}
