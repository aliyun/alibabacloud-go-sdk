// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotCompareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetHotspotCompareRequest
	GetXDebugId() *string
	SetBeg1End(v int64) *GetHotspotCompareRequest
	GetBeg1End() *int64
	SetBeg1Start(v int64) *GetHotspotCompareRequest
	GetBeg1Start() *int64
	SetBeg2End(v int64) *GetHotspotCompareRequest
	GetBeg2End() *int64
	SetBeg2Start(v int64) *GetHotspotCompareRequest
	GetBeg2Start() *int64
	SetHotType(v string) *GetHotspotCompareRequest
	GetHotType() *string
	SetInstance1(v string) *GetHotspotCompareRequest
	GetInstance1() *string
	SetInstance2(v string) *GetHotspotCompareRequest
	GetInstance2() *string
	SetPid1(v int64) *GetHotspotCompareRequest
	GetPid1() *int64
	SetPid2(v int64) *GetHotspotCompareRequest
	GetPid2() *int64
	SetTable(v string) *GetHotspotCompareRequest
	GetTable() *string
	SetXSysomInvokeSource(v string) *GetHotspotCompareRequest
	GetXSysomInvokeSource() *string
}

type GetHotspotCompareRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The end time 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725415774000
	Beg1End *int64 `json:"beg1_end,omitempty" xml:"beg1_end,omitempty"`
	// The start time 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725415474000
	Beg1Start *int64 `json:"beg1_start,omitempty" xml:"beg1_start,omitempty"`
	// The end time 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725415774000
	Beg2End *int64 `json:"beg2_end,omitempty" xml:"beg2_end,omitempty"`
	// The start time 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725415474000
	Beg2Start *int64 `json:"beg2_start,omitempty" xml:"beg2_start,omitempty"`
	// The hot spot type.
	//
	// example:
	//
	// alloc
	HotType *string `json:"hot_type,omitempty" xml:"hot_type,omitempty"`
	// The instance 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zei55fwj8nnu31h3z46"
	Instance1 *string `json:"instance1,omitempty" xml:"instance1,omitempty"`
	// The instance 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Instance2 *string `json:"instance2,omitempty" xml:"instance2,omitempty"`
	// The process of instance 1.
	//
	// example:
	//
	// 0
	Pid1 *int64 `json:"pid1,omitempty" xml:"pid1,omitempty"`
	// The process of instance 2.
	//
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Pid2 *int64 `json:"pid2,omitempty" xml:"pid2,omitempty"`
	// The name of the table to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table              *string `json:"table,omitempty" xml:"table,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetHotspotCompareRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetHotspotCompareRequest) GetBeg1End() *int64 {
	return s.Beg1End
}

func (s *GetHotspotCompareRequest) GetBeg1Start() *int64 {
	return s.Beg1Start
}

func (s *GetHotspotCompareRequest) GetBeg2End() *int64 {
	return s.Beg2End
}

func (s *GetHotspotCompareRequest) GetBeg2Start() *int64 {
	return s.Beg2Start
}

func (s *GetHotspotCompareRequest) GetHotType() *string {
	return s.HotType
}

func (s *GetHotspotCompareRequest) GetInstance1() *string {
	return s.Instance1
}

func (s *GetHotspotCompareRequest) GetInstance2() *string {
	return s.Instance2
}

func (s *GetHotspotCompareRequest) GetPid1() *int64 {
	return s.Pid1
}

func (s *GetHotspotCompareRequest) GetPid2() *int64 {
	return s.Pid2
}

func (s *GetHotspotCompareRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotspotCompareRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetHotspotCompareRequest) SetXDebugId(v string) *GetHotspotCompareRequest {
	s.XDebugId = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg1End(v int64) *GetHotspotCompareRequest {
	s.Beg1End = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg1Start(v int64) *GetHotspotCompareRequest {
	s.Beg1Start = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg2End(v int64) *GetHotspotCompareRequest {
	s.Beg2End = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg2Start(v int64) *GetHotspotCompareRequest {
	s.Beg2Start = &v
	return s
}

func (s *GetHotspotCompareRequest) SetHotType(v string) *GetHotspotCompareRequest {
	s.HotType = &v
	return s
}

func (s *GetHotspotCompareRequest) SetInstance1(v string) *GetHotspotCompareRequest {
	s.Instance1 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetInstance2(v string) *GetHotspotCompareRequest {
	s.Instance2 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetPid1(v int64) *GetHotspotCompareRequest {
	s.Pid1 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetPid2(v int64) *GetHotspotCompareRequest {
	s.Pid2 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetTable(v string) *GetHotspotCompareRequest {
	s.Table = &v
	return s
}

func (s *GetHotspotCompareRequest) SetXSysomInvokeSource(v string) *GetHotspotCompareRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetHotspotCompareRequest) Validate() error {
	return dara.Validate(s)
}
