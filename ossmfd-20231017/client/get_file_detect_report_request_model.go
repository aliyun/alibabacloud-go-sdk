// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *GetFileDetectReportRequest
	GetEventId() *int64
	SetField(v string) *GetFileDetectReportRequest
	GetField() *string
	SetFileHash(v string) *GetFileDetectReportRequest
	GetFileHash() *string
	SetLang(v string) *GetFileDetectReportRequest
	GetLang() *string
	SetSourceType(v string) *GetFileDetectReportRequest
	GetSourceType() *string
}

type GetFileDetectReportRequest struct {
	// example:
	//
	// 100******
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// Basic,ThreatTypes,Intelligences,Sandbox
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// 7ab4ee98d535623991bc178718******
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// object_scan
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetFileDetectReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectReportRequest) GoString() string {
	return s.String()
}

func (s *GetFileDetectReportRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *GetFileDetectReportRequest) GetField() *string {
	return s.Field
}

func (s *GetFileDetectReportRequest) GetFileHash() *string {
	return s.FileHash
}

func (s *GetFileDetectReportRequest) GetLang() *string {
	return s.Lang
}

func (s *GetFileDetectReportRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetFileDetectReportRequest) SetEventId(v int64) *GetFileDetectReportRequest {
	s.EventId = &v
	return s
}

func (s *GetFileDetectReportRequest) SetField(v string) *GetFileDetectReportRequest {
	s.Field = &v
	return s
}

func (s *GetFileDetectReportRequest) SetFileHash(v string) *GetFileDetectReportRequest {
	s.FileHash = &v
	return s
}

func (s *GetFileDetectReportRequest) SetLang(v string) *GetFileDetectReportRequest {
	s.Lang = &v
	return s
}

func (s *GetFileDetectReportRequest) SetSourceType(v string) *GetFileDetectReportRequest {
	s.SourceType = &v
	return s
}

func (s *GetFileDetectReportRequest) Validate() error {
	return dara.Validate(s)
}
