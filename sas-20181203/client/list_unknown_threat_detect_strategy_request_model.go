// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *ListUnknownThreatDetectStrategyRequest
	GetCurrentPage() *string
	SetId(v string) *ListUnknownThreatDetectStrategyRequest
	GetId() *string
	SetName(v string) *ListUnknownThreatDetectStrategyRequest
	GetName() *string
	SetPageSize(v string) *ListUnknownThreatDetectStrategyRequest
	GetPageSize() *string
	SetStudyMode(v string) *ListUnknownThreatDetectStrategyRequest
	GetStudyMode() *string
}

type ListUnknownThreatDetectStrategyRequest struct {
	// The page number of the current page in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 210****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The policy name.
	//
	// example:
	//
	// strategy****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to display on each page in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The whitelist mode. Valid values:
	//
	// - **hash**: process hash
	//
	// - **path**: process path
	//
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
}

func (s ListUnknownThreatDetectStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectStrategyRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectStrategyRequest) GetId() *string {
	return s.Id
}

func (s *ListUnknownThreatDetectStrategyRequest) GetName() *string {
	return s.Name
}

func (s *ListUnknownThreatDetectStrategyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListUnknownThreatDetectStrategyRequest) GetStudyMode() *string {
	return s.StudyMode
}

func (s *ListUnknownThreatDetectStrategyRequest) SetCurrentPage(v string) *ListUnknownThreatDetectStrategyRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyRequest) SetId(v string) *ListUnknownThreatDetectStrategyRequest {
	s.Id = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyRequest) SetName(v string) *ListUnknownThreatDetectStrategyRequest {
	s.Name = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyRequest) SetPageSize(v string) *ListUnknownThreatDetectStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyRequest) SetStudyMode(v string) *ListUnknownThreatDetectStrategyRequest {
	s.StudyMode = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyRequest) Validate() error {
	return dara.Validate(s)
}
