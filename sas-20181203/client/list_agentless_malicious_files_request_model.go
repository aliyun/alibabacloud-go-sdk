// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessMaliciousFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessMaliciousFilesRequest
	GetCurrentPage() *int32
	SetDealed(v string) *ListAgentlessMaliciousFilesRequest
	GetDealed() *string
	SetEventId(v int64) *ListAgentlessMaliciousFilesRequest
	GetEventId() *int64
	SetFuzzyMaliciousName(v string) *ListAgentlessMaliciousFilesRequest
	GetFuzzyMaliciousName() *string
	SetLang(v string) *ListAgentlessMaliciousFilesRequest
	GetLang() *string
	SetLevels(v string) *ListAgentlessMaliciousFilesRequest
	GetLevels() *string
	SetMaliciousMd5(v string) *ListAgentlessMaliciousFilesRequest
	GetMaliciousMd5() *string
	SetMaliciousType(v string) *ListAgentlessMaliciousFilesRequest
	GetMaliciousType() *string
	SetPageSize(v string) *ListAgentlessMaliciousFilesRequest
	GetPageSize() *string
	SetRemark(v string) *ListAgentlessMaliciousFilesRequest
	GetRemark() *string
	SetScanRange(v []*string) *ListAgentlessMaliciousFilesRequest
	GetScanRange() []*string
	SetUuid(v string) *ListAgentlessMaliciousFilesRequest
	GetUuid() *string
}

type ListAgentlessMaliciousFilesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the vulnerability is handled. Valid values:
	//
	// 	- Y: The vulnerability is handled.
	//
	// 	- N: The vulnerability is not handled.
	//
	// example:
	//
	// Y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 81****
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the malicious file that you want to query.
	//
	// > Fuzzy match is supported.
	//
	// example:
	//
	// WebShell
	FuzzyMaliciousName *string `json:"FuzzyMaliciousName,omitempty" xml:"FuzzyMaliciousName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severities of the malicious files. Separate multiple values with commas (,). Valid values:
	//
	// 	- serious
	//
	// 	- suspicious
	//
	// 	- remind
	//
	// example:
	//
	// remind,suspicious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The MD5 hash value of the malicious file.
	//
	// example:
	//
	// d836968041f7683b5459****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The alert type.
	//
	// Valid values when Lang is set to zh:
	//
	// 	- WebShell
	//
	// 	- Malicious Software
	//
	// 	- Malicious Script
	//
	// Valid values when Lang is set to en:
	//
	// 	- WebShell
	//
	// 	- Malicious Software
	//
	// 	- Malicious Script
	//
	// example:
	//
	// WebShell
	MaliciousType *string `json:"MaliciousType,omitempty" xml:"MaliciousType,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The remarks for the asset affected by the vulnerability. The value can be the private IP address, public IP address, or name of the asset. Fuzzy match is supported.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source of the malicious file.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The UUID of the asset.
	//
	// example:
	//
	// d2d94e8b-bb25-4744-8004-1e08a53c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessMaliciousFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessMaliciousFilesRequest) GetDealed() *string {
	return s.Dealed
}

func (s *ListAgentlessMaliciousFilesRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *ListAgentlessMaliciousFilesRequest) GetFuzzyMaliciousName() *string {
	return s.FuzzyMaliciousName
}

func (s *ListAgentlessMaliciousFilesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAgentlessMaliciousFilesRequest) GetLevels() *string {
	return s.Levels
}

func (s *ListAgentlessMaliciousFilesRequest) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *ListAgentlessMaliciousFilesRequest) GetMaliciousType() *string {
	return s.MaliciousType
}

func (s *ListAgentlessMaliciousFilesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAgentlessMaliciousFilesRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListAgentlessMaliciousFilesRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *ListAgentlessMaliciousFilesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessMaliciousFilesRequest) SetCurrentPage(v int32) *ListAgentlessMaliciousFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetDealed(v string) *ListAgentlessMaliciousFilesRequest {
	s.Dealed = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetEventId(v int64) *ListAgentlessMaliciousFilesRequest {
	s.EventId = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetFuzzyMaliciousName(v string) *ListAgentlessMaliciousFilesRequest {
	s.FuzzyMaliciousName = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetLang(v string) *ListAgentlessMaliciousFilesRequest {
	s.Lang = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetLevels(v string) *ListAgentlessMaliciousFilesRequest {
	s.Levels = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetMaliciousMd5(v string) *ListAgentlessMaliciousFilesRequest {
	s.MaliciousMd5 = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetMaliciousType(v string) *ListAgentlessMaliciousFilesRequest {
	s.MaliciousType = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetPageSize(v string) *ListAgentlessMaliciousFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetRemark(v string) *ListAgentlessMaliciousFilesRequest {
	s.Remark = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetScanRange(v []*string) *ListAgentlessMaliciousFilesRequest {
	s.ScanRange = v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) SetUuid(v string) *ListAgentlessMaliciousFilesRequest {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessMaliciousFilesRequest) Validate() error {
	return dara.Validate(s)
}
