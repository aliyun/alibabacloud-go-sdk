// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetHotspotAnalysisRequest
	GetAppType() *string
	SetBegEnd(v int64) *GetHotspotAnalysisRequest
	GetBegEnd() *int64
	SetBegStart(v int64) *GetHotspotAnalysisRequest
	GetBegStart() *int64
	SetInstance(v string) *GetHotspotAnalysisRequest
	GetInstance() *string
	SetPid(v int64) *GetHotspotAnalysisRequest
	GetPid() *int64
	SetTable(v string) *GetHotspotAnalysisRequest
	GetTable() *string
}

type GetHotspotAnalysisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GR
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
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
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1657494
	Pid *int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetHotspotAnalysisRequest) GetBegEnd() *int64 {
	return s.BegEnd
}

func (s *GetHotspotAnalysisRequest) GetBegStart() *int64 {
	return s.BegStart
}

func (s *GetHotspotAnalysisRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHotspotAnalysisRequest) GetPid() *int64 {
	return s.Pid
}

func (s *GetHotspotAnalysisRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotspotAnalysisRequest) SetAppType(v string) *GetHotspotAnalysisRequest {
	s.AppType = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetBegEnd(v int64) *GetHotspotAnalysisRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetBegStart(v int64) *GetHotspotAnalysisRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetInstance(v string) *GetHotspotAnalysisRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetPid(v int64) *GetHotspotAnalysisRequest {
	s.Pid = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetTable(v string) *GetHotspotAnalysisRequest {
	s.Table = &v
	return s
}

func (s *GetHotspotAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
