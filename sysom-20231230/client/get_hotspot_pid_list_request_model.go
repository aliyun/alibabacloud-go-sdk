// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotPidListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetHotspotPidListRequest
	GetXDebugId() *string
	SetBegEnd(v int64) *GetHotspotPidListRequest
	GetBegEnd() *int64
	SetBegStart(v int64) *GetHotspotPidListRequest
	GetBegStart() *int64
	SetInstance(v string) *GetHotspotPidListRequest
	GetInstance() *string
	SetTable(v string) *GetHotspotPidListRequest
	GetTable() *string
	SetXSysomInvokeSource(v string) *GetHotspotPidListRequest
	GetXSysomInvokeSource() *string
}

type GetHotspotPidListRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725413947000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// The start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725413647000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
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

func (s GetHotspotPidListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotPidListRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetHotspotPidListRequest) GetBegEnd() *int64 {
	return s.BegEnd
}

func (s *GetHotspotPidListRequest) GetBegStart() *int64 {
	return s.BegStart
}

func (s *GetHotspotPidListRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHotspotPidListRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotspotPidListRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetHotspotPidListRequest) SetXDebugId(v string) *GetHotspotPidListRequest {
	s.XDebugId = &v
	return s
}

func (s *GetHotspotPidListRequest) SetBegEnd(v int64) *GetHotspotPidListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotPidListRequest) SetBegStart(v int64) *GetHotspotPidListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotPidListRequest) SetInstance(v string) *GetHotspotPidListRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotPidListRequest) SetTable(v string) *GetHotspotPidListRequest {
	s.Table = &v
	return s
}

func (s *GetHotspotPidListRequest) SetXSysomInvokeSource(v string) *GetHotspotPidListRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetHotspotPidListRequest) Validate() error {
	return dara.Validate(s)
}
