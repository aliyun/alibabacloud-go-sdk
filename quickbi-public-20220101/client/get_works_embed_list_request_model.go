// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorksEmbedListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *GetWorksEmbedListRequest
	GetKeyword() *string
	SetPageNo(v int32) *GetWorksEmbedListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetWorksEmbedListRequest
	GetPageSize() *int32
	SetWorksType(v string) *GetWorksEmbedListRequest
	GetWorksType() *string
	SetWsId(v string) *GetWorksEmbedListRequest
	GetWsId() *string
}

type GetWorksEmbedListRequest struct {
	// The report name. Fuzzy match is supported.
	//
	// example:
	//
	// Test dataset
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The report type. Valid values:
	//
	// - page: Dashboard.
	//
	// - screen: Data screen.
	//
	// - report: Workbook.
	//
	// - ANALYSIS: Ad hoc analysis.
	//
	// - dashboardOfflineQuery: Self-service data retrieval.
	//
	// - dataForm: Data entry.
	//
	// example:
	//
	// page
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 919818-***-*****-wdasd
	WsId *string `json:"WsId,omitempty" xml:"WsId,omitempty"`
}

func (s GetWorksEmbedListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorksEmbedListRequest) GoString() string {
	return s.String()
}

func (s *GetWorksEmbedListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetWorksEmbedListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetWorksEmbedListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetWorksEmbedListRequest) GetWorksType() *string {
	return s.WorksType
}

func (s *GetWorksEmbedListRequest) GetWsId() *string {
	return s.WsId
}

func (s *GetWorksEmbedListRequest) SetKeyword(v string) *GetWorksEmbedListRequest {
	s.Keyword = &v
	return s
}

func (s *GetWorksEmbedListRequest) SetPageNo(v int32) *GetWorksEmbedListRequest {
	s.PageNo = &v
	return s
}

func (s *GetWorksEmbedListRequest) SetPageSize(v int32) *GetWorksEmbedListRequest {
	s.PageSize = &v
	return s
}

func (s *GetWorksEmbedListRequest) SetWorksType(v string) *GetWorksEmbedListRequest {
	s.WorksType = &v
	return s
}

func (s *GetWorksEmbedListRequest) SetWsId(v string) *GetWorksEmbedListRequest {
	s.WsId = &v
	return s
}

func (s *GetWorksEmbedListRequest) Validate() error {
	return dara.Validate(s)
}
