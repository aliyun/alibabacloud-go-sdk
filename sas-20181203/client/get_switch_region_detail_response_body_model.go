// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwitchRegionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSwitchRegionDetailResponseBodyData) *GetSwitchRegionDetailResponseBody
	GetData() *GetSwitchRegionDetailResponseBodyData
	SetRequestId(v string) *GetSwitchRegionDetailResponseBody
	GetRequestId() *string
}

type GetSwitchRegionDetailResponseBody struct {
	// The response parameters.
	Data *GetSwitchRegionDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSwitchRegionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSwitchRegionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwitchRegionDetailResponseBody) GetData() *GetSwitchRegionDetailResponseBodyData {
	return s.Data
}

func (s *GetSwitchRegionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSwitchRegionDetailResponseBody) SetData(v *GetSwitchRegionDetailResponseBodyData) *GetSwitchRegionDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSwitchRegionDetailResponseBody) SetRequestId(v string) *GetSwitchRegionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSwitchRegionDetailResponseBodyData struct {
	// The time when the permissions were modified.
	//
	// example:
	//
	// 1692858597000
	GmtIsAgreeModified *int64 `json:"GmtIsAgreeModified,omitempty" xml:"GmtIsAgreeModified,omitempty"`
	// The notification time.
	//
	// example:
	//
	// 1692858597000
	GmtNoticed *int64 `json:"GmtNoticed,omitempty" xml:"GmtNoticed,omitempty"`
	// Indicates whether the migration is approved.
	//
	// example:
	//
	// true
	IsAgree *string `json:"IsAgree,omitempty" xml:"IsAgree,omitempty"`
	// Indicates whether the notification is sent.
	//
	// example:
	//
	// YES
	IsNoticed  *string `json:"IsNoticed,omitempty" xml:"IsNoticed,omitempty"`
	NeedSwitch *bool   `json:"NeedSwitch,omitempty" xml:"NeedSwitch,omitempty"`
	// The status of the switching to the region.
	RegionStatus []*GetSwitchRegionDetailResponseBodyDataRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s GetSwitchRegionDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSwitchRegionDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSwitchRegionDetailResponseBodyData) GetGmtIsAgreeModified() *int64 {
	return s.GmtIsAgreeModified
}

func (s *GetSwitchRegionDetailResponseBodyData) GetGmtNoticed() *int64 {
	return s.GmtNoticed
}

func (s *GetSwitchRegionDetailResponseBodyData) GetIsAgree() *string {
	return s.IsAgree
}

func (s *GetSwitchRegionDetailResponseBodyData) GetIsNoticed() *string {
	return s.IsNoticed
}

func (s *GetSwitchRegionDetailResponseBodyData) GetNeedSwitch() *bool {
	return s.NeedSwitch
}

func (s *GetSwitchRegionDetailResponseBodyData) GetRegionStatus() []*GetSwitchRegionDetailResponseBodyDataRegionStatus {
	return s.RegionStatus
}

func (s *GetSwitchRegionDetailResponseBodyData) SetGmtIsAgreeModified(v int64) *GetSwitchRegionDetailResponseBodyData {
	s.GmtIsAgreeModified = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) SetGmtNoticed(v int64) *GetSwitchRegionDetailResponseBodyData {
	s.GmtNoticed = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) SetIsAgree(v string) *GetSwitchRegionDetailResponseBodyData {
	s.IsAgree = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) SetIsNoticed(v string) *GetSwitchRegionDetailResponseBodyData {
	s.IsNoticed = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) SetNeedSwitch(v bool) *GetSwitchRegionDetailResponseBodyData {
	s.NeedSwitch = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) SetRegionStatus(v []*GetSwitchRegionDetailResponseBodyDataRegionStatus) *GetSwitchRegionDetailResponseBodyData {
	s.RegionStatus = v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSwitchRegionDetailResponseBodyDataRegionStatus struct {
	// The number of ECS instances.
	//
	// example:
	//
	// 27
	EcsCount *int32 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// The time when the migration is scheduled.
	//
	// example:
	//
	// 1692858597000
	GmtPlanSwitchTime *int64 `json:"GmtPlanSwitchTime,omitempty" xml:"GmtPlanSwitchTime,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// us-east-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The migration status. Valid values:
	//
	// 	- **0**: pending
	//
	// 	- **1**: successful
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSwitchRegionDetailResponseBodyDataRegionStatus) String() string {
	return dara.Prettify(s)
}

func (s GetSwitchRegionDetailResponseBodyDataRegionStatus) GoString() string {
	return s.String()
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) GetGmtPlanSwitchTime() *int64 {
	return s.GmtPlanSwitchTime
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) GetStatus() *int32 {
	return s.Status
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) SetEcsCount(v int32) *GetSwitchRegionDetailResponseBodyDataRegionStatus {
	s.EcsCount = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) SetGmtPlanSwitchTime(v int64) *GetSwitchRegionDetailResponseBodyDataRegionStatus {
	s.GmtPlanSwitchTime = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) SetRegionId(v string) *GetSwitchRegionDetailResponseBodyDataRegionStatus {
	s.RegionId = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) SetStatus(v int32) *GetSwitchRegionDetailResponseBodyDataRegionStatus {
	s.Status = &v
	return s
}

func (s *GetSwitchRegionDetailResponseBodyDataRegionStatus) Validate() error {
	return dara.Validate(s)
}
