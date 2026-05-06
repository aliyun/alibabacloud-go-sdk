// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryAvgMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgMetricList(v []*QueryHistoryAvgMetricListResponseBodyAvgMetricList) *QueryHistoryAvgMetricListResponseBody
	GetAvgMetricList() []*QueryHistoryAvgMetricListResponseBodyAvgMetricList
	SetRequestId(v string) *QueryHistoryAvgMetricListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryHistoryAvgMetricListResponseBody
	GetTotalCount() *int64
}

type QueryHistoryAvgMetricListResponseBody struct {
	AvgMetricList []*QueryHistoryAvgMetricListResponseBodyAvgMetricList `json:"AvgMetricList,omitempty" xml:"AvgMetricList,omitempty" type:"Repeated"`
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHistoryAvgMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListResponseBody) GetAvgMetricList() []*QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	return s.AvgMetricList
}

func (s *QueryHistoryAvgMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoryAvgMetricListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryHistoryAvgMetricListResponseBody) SetAvgMetricList(v []*QueryHistoryAvgMetricListResponseBodyAvgMetricList) *QueryHistoryAvgMetricListResponseBody {
	s.AvgMetricList = v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBody) SetRequestId(v string) *QueryHistoryAvgMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBody) SetTotalCount(v int64) *QueryHistoryAvgMetricListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBody) Validate() error {
	if s.AvgMetricList != nil {
		for _, item := range s.AvgMetricList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryAvgMetricListResponseBodyAvgMetricList struct {
	// example:
	//
	// 20
	AvgValue *float32 `json:"AvgValue,omitempty" xml:"AvgValue,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// dg-bogo95eob5avnis9k
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// example:
	//
	// ecd-bx9i0nsjd3zmibnzq
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// C-051
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// example:
	//
	// eds.enterprise_office.8c16g
	DesktopType *string   `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	EndUserIds  []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// example:
	//
	// 0
	ManagementFlag *string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	// example:
	//
	// 2048
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// true
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// example:
	//
	// Winserver2025
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sessions []*QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// monthPackage
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
}

func (s QueryHistoryAvgMetricListResponseBodyAvgMetricList) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListResponseBodyAvgMetricList) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetAvgValue() *float32 {
	return s.AvgValue
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetDesktopId() *string {
	return s.DesktopId
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetDesktopName() *string {
	return s.DesktopName
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetDesktopType() *string {
	return s.DesktopType
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetManagementFlag() *string {
	return s.ManagementFlag
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetMemory() *int64 {
	return s.Memory
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetMultiResource() *bool {
	return s.MultiResource
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetPlatform() *string {
	return s.Platform
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetSessions() []*QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions {
	return s.Sessions
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) GetSubPayType() *string {
	return s.SubPayType
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetAvgValue(v float32) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.AvgValue = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetChargeType(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.ChargeType = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetCpu(v int32) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.Cpu = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetDesktopGroupId(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.DesktopGroupId = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetDesktopId(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.DesktopId = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetDesktopName(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.DesktopName = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetDesktopStatus(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.DesktopStatus = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetDesktopType(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.DesktopType = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetEndUserIds(v []*string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.EndUserIds = v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetGpuSpec(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.GpuSpec = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetManagementFlag(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.ManagementFlag = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetMemory(v int64) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.Memory = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetMultiResource(v bool) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.MultiResource = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetPlatform(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.Platform = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetRegionId(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.RegionId = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetSessions(v []*QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.Sessions = v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) SetSubPayType(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricList {
	s.SubPayType = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricList) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions struct {
	// example:
	//
	// testUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// 2026-04-17T14:51:53Z
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
	// example:
	//
	// testUserName
	ExternalUserName *string `json:"ExternalUserName,omitempty" xml:"ExternalUserName,omitempty"`
	// example:
	//
	// sz-sygc-07-03
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) GetEstablishmentTime() *string {
	return s.EstablishmentTime
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) GetExternalUserName() *string {
	return s.ExternalUserName
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) GetNickName() *string {
	return s.NickName
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) SetEndUserId(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions {
	s.EndUserId = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) SetEstablishmentTime(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions {
	s.EstablishmentTime = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) SetExternalUserName(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions {
	s.ExternalUserName = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) SetNickName(v string) *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions {
	s.NickName = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponseBodyAvgMetricListSessions) Validate() error {
	return dara.Validate(s)
}
