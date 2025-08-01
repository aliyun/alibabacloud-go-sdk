// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTrackingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBegEnd(v int64) *GetHotspotTrackingRequest
	GetBegEnd() *int64
	SetBegStart(v int64) *GetHotspotTrackingRequest
	GetBegStart() *int64
	SetHotType(v string) *GetHotspotTrackingRequest
	GetHotType() *string
	SetInstance(v string) *GetHotspotTrackingRequest
	GetInstance() *string
	SetPid(v int64) *GetHotspotTrackingRequest
	GetPid() *int64
	SetTable(v string) *GetHotspotTrackingRequest
	GetTable() *string
}

type GetHotspotTrackingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1725413948000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725410348000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	HotType *string `json:"hot_type,omitempty" xml:"hot_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1657494
	Pid *int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotTrackingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingRequest) GetBegEnd() *int64 {
	return s.BegEnd
}

func (s *GetHotspotTrackingRequest) GetBegStart() *int64 {
	return s.BegStart
}

func (s *GetHotspotTrackingRequest) GetHotType() *string {
	return s.HotType
}

func (s *GetHotspotTrackingRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHotspotTrackingRequest) GetPid() *int64 {
	return s.Pid
}

func (s *GetHotspotTrackingRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotspotTrackingRequest) SetBegEnd(v int64) *GetHotspotTrackingRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetBegStart(v int64) *GetHotspotTrackingRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetHotType(v string) *GetHotspotTrackingRequest {
	s.HotType = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetInstance(v string) *GetHotspotTrackingRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetPid(v int64) *GetHotspotTrackingRequest {
	s.Pid = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetTable(v string) *GetHotspotTrackingRequest {
	s.Table = &v
	return s
}

func (s *GetHotspotTrackingRequest) Validate() error {
	return dara.Validate(s)
}
