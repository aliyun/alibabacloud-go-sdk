// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SearchAlertRulesRequest
	GetAppType() *string
	SetCurrentPage(v int32) *SearchAlertRulesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *SearchAlertRulesRequest
	GetPageSize() *int32
	SetPid(v string) *SearchAlertRulesRequest
	GetPid() *string
	SetRegionId(v string) *SearchAlertRulesRequest
	GetRegionId() *string
	SetTitle(v string) *SearchAlertRulesRequest
	GetTitle() *string
	SetType(v string) *SearchAlertRulesRequest
	GetType() *string
}

type SearchAlertRulesRequest struct {
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchAlertRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchAlertRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertRulesRequest) GetPid() *string {
	return s.Pid
}

func (s *SearchAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertRulesRequest) GetTitle() *string {
	return s.Title
}

func (s *SearchAlertRulesRequest) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesRequest) SetAppType(v string) *SearchAlertRulesRequest {
	s.AppType = &v
	return s
}

func (s *SearchAlertRulesRequest) SetCurrentPage(v int32) *SearchAlertRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPageSize(v int32) *SearchAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPid(v string) *SearchAlertRulesRequest {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesRequest) SetRegionId(v string) *SearchAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetTitle(v string) *SearchAlertRulesRequest {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesRequest) SetType(v string) *SearchAlertRulesRequest {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
