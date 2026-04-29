// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUnknownThreatDetectEventRequest
	GetCurrentPage() *int32
	SetHashKey(v string) *ListUnknownThreatDetectEventRequest
	GetHashKey() *string
	SetPageSize(v int32) *ListUnknownThreatDetectEventRequest
	GetPageSize() *int32
	SetParentProcessPath(v string) *ListUnknownThreatDetectEventRequest
	GetParentProcessPath() *string
	SetProcessPath(v string) *ListUnknownThreatDetectEventRequest
	GetProcessPath() *string
	SetRemark(v string) *ListUnknownThreatDetectEventRequest
	GetRemark() *string
	SetStatus(v int32) *ListUnknownThreatDetectEventRequest
	GetStatus() *int32
	SetUuid(v string) *ListUnknownThreatDetectEventRequest
	GetUuid() *string
}

type ListUnknownThreatDetectEventRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// /bin/bash
	ParentProcessPath *string `json:"ParentProcessPath,omitempty" xml:"ParentProcessPath,omitempty"`
	// example:
	//
	// /test
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// example:
	//
	// 10.167.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 18b7336e-d469-473b-af83-8e5420f9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnknownThreatDetectEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectEventRequest) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectEventRequest) GetHashKey() *string {
	return s.HashKey
}

func (s *ListUnknownThreatDetectEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectEventRequest) GetParentProcessPath() *string {
	return s.ParentProcessPath
}

func (s *ListUnknownThreatDetectEventRequest) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *ListUnknownThreatDetectEventRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListUnknownThreatDetectEventRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListUnknownThreatDetectEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnknownThreatDetectEventRequest) SetCurrentPage(v int32) *ListUnknownThreatDetectEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetHashKey(v string) *ListUnknownThreatDetectEventRequest {
	s.HashKey = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetPageSize(v int32) *ListUnknownThreatDetectEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetParentProcessPath(v string) *ListUnknownThreatDetectEventRequest {
	s.ParentProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetProcessPath(v string) *ListUnknownThreatDetectEventRequest {
	s.ProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetRemark(v string) *ListUnknownThreatDetectEventRequest {
	s.Remark = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetStatus(v int32) *ListUnknownThreatDetectEventRequest {
	s.Status = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) SetUuid(v string) *ListUnknownThreatDetectEventRequest {
	s.Uuid = &v
	return s
}

func (s *ListUnknownThreatDetectEventRequest) Validate() error {
	return dara.Validate(s)
}
