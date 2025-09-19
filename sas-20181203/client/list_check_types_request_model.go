// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCheckTypesRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListCheckTypesRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckTypesRequest
	GetPageSize() *int32
	SetRiskId(v int64) *ListCheckTypesRequest
	GetRiskId() *int64
	SetShowChecks(v bool) *ListCheckTypesRequest
	GetShowChecks() *bool
	SetSource(v string) *ListCheckTypesRequest
	GetSource() *string
	SetUuid(v string) *ListCheckTypesRequest
	GetUuid() *string
}

type ListCheckTypesRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the baseline.
	//
	// >  You can call the [DescribeCheckWarningSummary](https://help.aliyun.com/document_detail/116179.html) operation to query the IDs of baselines.
	//
	// example:
	//
	// 34
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// Whether to query the check item list. The default value is false. Valid values:
	//
	// - **false**: Not Query
	//
	// - **true**: Query
	//
	// example:
	//
	// false
	ShowChecks *bool `json:"ShowChecks,omitempty" xml:"ShowChecks,omitempty"`
	// The data source. Default value: **default**. Valid values:
	//
	// 	- **agentless**: The check items of baselines for agentless detection.
	//
	// 	- **default**: The check items of baselines for hosts.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 293b07cb-db2d-4f39-941f-b2e4abb8****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListCheckTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckTypesRequest) GoString() string {
	return s.String()
}

func (s *ListCheckTypesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckTypesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckTypesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckTypesRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *ListCheckTypesRequest) GetShowChecks() *bool {
	return s.ShowChecks
}

func (s *ListCheckTypesRequest) GetSource() *string {
	return s.Source
}

func (s *ListCheckTypesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListCheckTypesRequest) SetCurrentPage(v int32) *ListCheckTypesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckTypesRequest) SetLang(v string) *ListCheckTypesRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckTypesRequest) SetPageSize(v int32) *ListCheckTypesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckTypesRequest) SetRiskId(v int64) *ListCheckTypesRequest {
	s.RiskId = &v
	return s
}

func (s *ListCheckTypesRequest) SetShowChecks(v bool) *ListCheckTypesRequest {
	s.ShowChecks = &v
	return s
}

func (s *ListCheckTypesRequest) SetSource(v string) *ListCheckTypesRequest {
	s.Source = &v
	return s
}

func (s *ListCheckTypesRequest) SetUuid(v string) *ListCheckTypesRequest {
	s.Uuid = &v
	return s
}

func (s *ListCheckTypesRequest) Validate() error {
	return dara.Validate(s)
}
