// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzeResult(v string) *ListUnknownThreatDetectProcessRequest
	GetAnalyzeResult() *string
	SetCurrentPage(v int32) *ListUnknownThreatDetectProcessRequest
	GetCurrentPage() *int32
	SetFirstTimeEnd(v int64) *ListUnknownThreatDetectProcessRequest
	GetFirstTimeEnd() *int64
	SetFirstTimeStart(v int64) *ListUnknownThreatDetectProcessRequest
	GetFirstTimeStart() *int64
	SetMd5(v string) *ListUnknownThreatDetectProcessRequest
	GetMd5() *string
	SetPageSize(v int32) *ListUnknownThreatDetectProcessRequest
	GetPageSize() *int32
	SetPath(v string) *ListUnknownThreatDetectProcessRequest
	GetPath() *string
	SetProcessPath(v string) *ListUnknownThreatDetectProcessRequest
	GetProcessPath() *string
	SetRemark(v string) *ListUnknownThreatDetectProcessRequest
	GetRemark() *string
	SetSha256(v string) *ListUnknownThreatDetectProcessRequest
	GetSha256() *string
	SetUuid(v string) *ListUnknownThreatDetectProcessRequest
	GetUuid() *string
}

type ListUnknownThreatDetectProcessRequest struct {
	// example:
	//
	// white
	AnalyzeResult *string `json:"AnalyzeResult,omitempty" xml:"AnalyzeResult,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1768891966345
	FirstTimeEnd *int64 `json:"FirstTimeEnd,omitempty" xml:"FirstTimeEnd,omitempty"`
	// example:
	//
	// 1768891966344
	FirstTimeStart *int64 `json:"FirstTimeStart,omitempty" xml:"FirstTimeStart,omitempty"`
	// example:
	//
	// 0552c44e243abdea1729d4507bce****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// /etc/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// /bin/rm
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// example:
	//
	// 172.20.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// f204693a7d2ce99d6c4434e550d985ee1c7be7cb5dd9a76094369af0d2******
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnknownThreatDetectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectProcessRequest) GetAnalyzeResult() *string {
	return s.AnalyzeResult
}

func (s *ListUnknownThreatDetectProcessRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectProcessRequest) GetFirstTimeEnd() *int64 {
	return s.FirstTimeEnd
}

func (s *ListUnknownThreatDetectProcessRequest) GetFirstTimeStart() *int64 {
	return s.FirstTimeStart
}

func (s *ListUnknownThreatDetectProcessRequest) GetMd5() *string {
	return s.Md5
}

func (s *ListUnknownThreatDetectProcessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectProcessRequest) GetPath() *string {
	return s.Path
}

func (s *ListUnknownThreatDetectProcessRequest) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *ListUnknownThreatDetectProcessRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListUnknownThreatDetectProcessRequest) GetSha256() *string {
	return s.Sha256
}

func (s *ListUnknownThreatDetectProcessRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnknownThreatDetectProcessRequest) SetAnalyzeResult(v string) *ListUnknownThreatDetectProcessRequest {
	s.AnalyzeResult = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetCurrentPage(v int32) *ListUnknownThreatDetectProcessRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetFirstTimeEnd(v int64) *ListUnknownThreatDetectProcessRequest {
	s.FirstTimeEnd = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetFirstTimeStart(v int64) *ListUnknownThreatDetectProcessRequest {
	s.FirstTimeStart = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetMd5(v string) *ListUnknownThreatDetectProcessRequest {
	s.Md5 = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetPageSize(v int32) *ListUnknownThreatDetectProcessRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetPath(v string) *ListUnknownThreatDetectProcessRequest {
	s.Path = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetProcessPath(v string) *ListUnknownThreatDetectProcessRequest {
	s.ProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetRemark(v string) *ListUnknownThreatDetectProcessRequest {
	s.Remark = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetSha256(v string) *ListUnknownThreatDetectProcessRequest {
	s.Sha256 = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) SetUuid(v string) *ListUnknownThreatDetectProcessRequest {
	s.Uuid = &v
	return s
}

func (s *ListUnknownThreatDetectProcessRequest) Validate() error {
	return dara.Validate(s)
}
