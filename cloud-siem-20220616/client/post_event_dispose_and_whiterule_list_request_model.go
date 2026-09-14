// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventDisposeAndWhiteruleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PostEventDisposeAndWhiteruleListRequest
	GetClientToken() *string
	SetDisposeStrategyIds(v string) *PostEventDisposeAndWhiteruleListRequest
	GetDisposeStrategyIds() *string
	SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest
	GetEventDispose() *string
	SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest
	GetIncidentUuid() *string
	SetOwner(v string) *PostEventDisposeAndWhiteruleListRequest
	GetOwner() *string
	SetReceiverInfo(v string) *PostEventDisposeAndWhiteruleListRequest
	GetReceiverInfo() *string
	SetRegionId(v string) *PostEventDisposeAndWhiteruleListRequest
	GetRegionId() *string
	SetRemark(v string) *PostEventDisposeAndWhiteruleListRequest
	GetRemark() *string
	SetResponseSource(v string) *PostEventDisposeAndWhiteruleListRequest
	GetResponseSource() *string
	SetRoleFor(v int64) *PostEventDisposeAndWhiteruleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostEventDisposeAndWhiteruleListRequest
	GetRoleType() *int32
	SetStatus(v int32) *PostEventDisposeAndWhiteruleListRequest
	GetStatus() *int32
	SetSyncAlertStatus(v bool) *PostEventDisposeAndWhiteruleListRequest
	GetSyncAlertStatus() *bool
	SetThreatLevel(v string) *PostEventDisposeAndWhiteruleListRequest
	GetThreatLevel() *string
}

type PostEventDisposeAndWhiteruleListRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of handling policy IDs.
	//
	// example:
	//
	// 12,13,14
	DisposeStrategyIds *string `json:"DisposeStrategyIds,omitempty" xml:"DisposeStrategyIds,omitempty"`
	// The incident handling configuration as a JSON object.
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
	// The globally unique UUID of the incident.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The account UID of the incident owner.
	//
	// example:
	//
	// 1234567890xxxxxx
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert recipient configuration as a JSON object.
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
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets reside in regions in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets reside in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks for the incident.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source of the handling policy.
	//
	// example:
	//
	// system
	ResponseSource *string `json:"ResponseSource,omitempty" xml:"ResponseSource,omitempty"`
	// The ID of the user for whom the administrator switches to a member view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the China account view.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The incident status. Valid values:
	//
	// - 0: unhandled
	//
	// - 1: handling
	//
	// - 5: handling failed
	//
	// - 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to restore associated handled alerts to unhandled status when reopening the incident.
	SyncAlertStatus *bool `json:"SyncAlertStatus,omitempty" xml:"SyncAlertStatus,omitempty"`
	// The threat level. Valid values:
	//
	// - serious: high
	//
	// - suspicious: medium
	//
	// - remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListRequest) String() string {
	return dara.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetDisposeStrategyIds() *string {
	return s.DisposeStrategyIds
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetEventDispose() *string {
	return s.EventDispose
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetOwner() *string {
	return s.Owner
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

func (s *PostEventDisposeAndWhiteruleListRequest) GetResponseSource() *string {
	return s.ResponseSource
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

func (s *PostEventDisposeAndWhiteruleListRequest) GetSyncAlertStatus() *bool {
	return s.SyncAlertStatus
}

func (s *PostEventDisposeAndWhiteruleListRequest) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetClientToken(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ClientToken = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetDisposeStrategyIds(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.DisposeStrategyIds = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.EventDispose = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetOwner(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.Owner = &v
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

func (s *PostEventDisposeAndWhiteruleListRequest) SetResponseSource(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ResponseSource = &v
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

func (s *PostEventDisposeAndWhiteruleListRequest) SetSyncAlertStatus(v bool) *PostEventDisposeAndWhiteruleListRequest {
	s.SyncAlertStatus = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetThreatLevel(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ThreatLevel = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) Validate() error {
	return dara.Validate(s)
}
