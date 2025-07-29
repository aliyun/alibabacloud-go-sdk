// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeMachines(v []*DescribeEdgeMachinesResponseBodyEdgeMachines) *DescribeEdgeMachinesResponseBody
	GetEdgeMachines() []*DescribeEdgeMachinesResponseBodyEdgeMachines
	SetPageInfo(v *DescribeEdgeMachinesResponseBodyPageInfo) *DescribeEdgeMachinesResponseBody
	GetPageInfo() *DescribeEdgeMachinesResponseBodyPageInfo
}

type DescribeEdgeMachinesResponseBody struct {
	// The list of cloud-native boxes.
	EdgeMachines []*DescribeEdgeMachinesResponseBodyEdgeMachines `json:"edge_machines,omitempty" xml:"edge_machines,omitempty" type:"Repeated"`
	// The paging information.
	PageInfo *DescribeEdgeMachinesResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeEdgeMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachinesResponseBody) GetEdgeMachines() []*DescribeEdgeMachinesResponseBodyEdgeMachines {
	return s.EdgeMachines
}

func (s *DescribeEdgeMachinesResponseBody) GetPageInfo() *DescribeEdgeMachinesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeEdgeMachinesResponseBody) SetEdgeMachines(v []*DescribeEdgeMachinesResponseBodyEdgeMachines) *DescribeEdgeMachinesResponseBody {
	s.EdgeMachines = v
	return s
}

func (s *DescribeEdgeMachinesResponseBody) SetPageInfo(v *DescribeEdgeMachinesResponseBodyPageInfo) *DescribeEdgeMachinesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeEdgeMachinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEdgeMachinesResponseBodyEdgeMachines struct {
	// The time when the cloud-native box was activated.
	//
	// example:
	//
	// 2021-07-19T16:07:48+08:00
	ActiveTime *string `json:"active_time,omitempty" xml:"active_time,omitempty"`
	// The time when the cloud-native box was created.
	//
	// example:
	//
	// 2021-07-07T20:44:00+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The device ID.
	//
	// example:
	//
	// c61083909b13f4a95b8554bda9577****
	EdgeMachineId *string `json:"edge_machine_id,omitempty" xml:"edge_machine_id,omitempty"`
	// The `hostname` of the cloud-native box.
	//
	// example:
	//
	// ack-v-b010-ssdfw****
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The lifecycle of the cloud-native box.
	//
	// example:
	//
	// activated
	LifeState *string `json:"life_state,omitempty" xml:"life_state,omitempty"`
	// The model of the cloud-native box.
	//
	// example:
	//
	// ACK-V-B010
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The machine name.
	//
	// example:
	//
	// ack-v-b010-ssdfw****
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the cloud-native box.
	//
	// example:
	//
	// online
	OnlineState *string `json:"online_state,omitempty" xml:"online_state,omitempty"`
	// The serial number.
	//
	// example:
	//
	// ACK9GBL31SXX****
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
	// The time when the cloud-native box was last updated.
	//
	// example:
	//
	// 2021-07-07T20:44:00+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeEdgeMachinesResponseBodyEdgeMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachinesResponseBodyEdgeMachines) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetActiveTime() *string {
	return s.ActiveTime
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetCreated() *string {
	return s.Created
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetEdgeMachineId() *string {
	return s.EdgeMachineId
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetLifeState() *string {
	return s.LifeState
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetModel() *string {
	return s.Model
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetName() *string {
	return s.Name
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetOnlineState() *string {
	return s.OnlineState
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetSn() *string {
	return s.Sn
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetActiveTime(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.ActiveTime = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetCreated(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Created = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetEdgeMachineId(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.EdgeMachineId = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetHostname(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Hostname = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetLifeState(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.LifeState = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetModel(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Model = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetName(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Name = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetOnlineState(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.OnlineState = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetSn(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Sn = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) SetUpdated(v string) *DescribeEdgeMachinesResponseBodyEdgeMachines {
	s.Updated = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyEdgeMachines) Validate() error {
	return dara.Validate(s)
}

type DescribeEdgeMachinesResponseBodyPageInfo struct {
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeEdgeMachinesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachinesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) SetPageNumber(v int32) *DescribeEdgeMachinesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) SetPageSize(v int32) *DescribeEdgeMachinesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeEdgeMachinesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeEdgeMachinesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
