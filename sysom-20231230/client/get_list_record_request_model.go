// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetListRecordRequest
	GetXDebugId() *string
	SetAnalysisId(v string) *GetListRecordRequest
	GetAnalysisId() *string
	SetCurrent(v int64) *GetListRecordRequest
	GetCurrent() *int64
	SetCustomId(v int64) *GetListRecordRequest
	GetCustomId() *int64
	SetPageSize(v int64) *GetListRecordRequest
	GetPageSize() *int64
	SetRegion(v string) *GetListRecordRequest
	GetRegion() *string
	SetXSysomInvokeSource(v string) *GetListRecordRequest
	GetXSysomInvokeSource() *string
}

type GetListRecordRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The ID of the analysis record to query the status of.
	//
	// example:
	//
	// xxxxxxxxx-f37b-478d-87f6-2259ed8766bf
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 5
	Current  *int64 `json:"current,omitempty" xml:"current,omitempty"`
	CustomId *int64 `json:"customId,omitempty" xml:"customId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetListRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListRecordRequest) GoString() string {
	return s.String()
}

func (s *GetListRecordRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetListRecordRequest) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *GetListRecordRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *GetListRecordRequest) GetCustomId() *int64 {
	return s.CustomId
}

func (s *GetListRecordRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetListRecordRequest) GetRegion() *string {
	return s.Region
}

func (s *GetListRecordRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetListRecordRequest) SetXDebugId(v string) *GetListRecordRequest {
	s.XDebugId = &v
	return s
}

func (s *GetListRecordRequest) SetAnalysisId(v string) *GetListRecordRequest {
	s.AnalysisId = &v
	return s
}

func (s *GetListRecordRequest) SetCurrent(v int64) *GetListRecordRequest {
	s.Current = &v
	return s
}

func (s *GetListRecordRequest) SetCustomId(v int64) *GetListRecordRequest {
	s.CustomId = &v
	return s
}

func (s *GetListRecordRequest) SetPageSize(v int64) *GetListRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetListRecordRequest) SetRegion(v string) *GetListRecordRequest {
	s.Region = &v
	return s
}

func (s *GetListRecordRequest) SetXSysomInvokeSource(v string) *GetListRecordRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetListRecordRequest) Validate() error {
	return dara.Validate(s)
}
