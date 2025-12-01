// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulItemPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *GetVulItemPageRequest
	GetAliasName() *string
	SetCurrentPage(v int32) *GetVulItemPageRequest
	GetCurrentPage() *int32
	SetDealed(v string) *GetVulItemPageRequest
	GetDealed() *string
	SetLevel(v string) *GetVulItemPageRequest
	GetLevel() *string
	SetName(v string) *GetVulItemPageRequest
	GetName() *string
	SetPageSize(v int32) *GetVulItemPageRequest
	GetPageSize() *int32
	SetScanType(v string) *GetVulItemPageRequest
	GetScanType() *string
}

type GetVulItemPageRequest struct {
	// Vulnerability alias.
	//
	// example:
	//
	// RHSA-2018:3665-Important: NetworkManager security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Processing status. y: processed; n: unprocessed; h: processing.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Risk level.
	//
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Vulnerability name.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20183665
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of items to display per page in the returned data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Vulnerability type.
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
}

func (s GetVulItemPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulItemPageRequest) GoString() string {
	return s.String()
}

func (s *GetVulItemPageRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *GetVulItemPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVulItemPageRequest) GetDealed() *string {
	return s.Dealed
}

func (s *GetVulItemPageRequest) GetLevel() *string {
	return s.Level
}

func (s *GetVulItemPageRequest) GetName() *string {
	return s.Name
}

func (s *GetVulItemPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVulItemPageRequest) GetScanType() *string {
	return s.ScanType
}

func (s *GetVulItemPageRequest) SetAliasName(v string) *GetVulItemPageRequest {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageRequest) SetCurrentPage(v int32) *GetVulItemPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageRequest) SetDealed(v string) *GetVulItemPageRequest {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageRequest) SetLevel(v string) *GetVulItemPageRequest {
	s.Level = &v
	return s
}

func (s *GetVulItemPageRequest) SetName(v string) *GetVulItemPageRequest {
	s.Name = &v
	return s
}

func (s *GetVulItemPageRequest) SetPageSize(v int32) *GetVulItemPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageRequest) SetScanType(v string) *GetVulItemPageRequest {
	s.ScanType = &v
	return s
}

func (s *GetVulItemPageRequest) Validate() error {
	return dara.Validate(s)
}
