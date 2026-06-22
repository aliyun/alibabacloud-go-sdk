// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckStructureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetCheckStructureRequest
	GetCurrentPage() *int32
	SetLang(v string) *GetCheckStructureRequest
	GetLang() *string
	SetPageSize(v int32) *GetCheckStructureRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetCheckStructureRequest
	GetRegionId() *string
	SetTaskSources(v []*string) *GetCheckStructureRequest
	GetTaskSources() []*string
}

type GetCheckStructureRequest struct {
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for requests and responses. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the asset resides. Valid values:
	//
	// - cn-hangzhou: China
	//
	// - ap-southeast-1: outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
}

func (s GetCheckStructureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureRequest) GoString() string {
	return s.String()
}

func (s *GetCheckStructureRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetCheckStructureRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCheckStructureRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCheckStructureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCheckStructureRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *GetCheckStructureRequest) SetCurrentPage(v int32) *GetCheckStructureRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetCheckStructureRequest) SetLang(v string) *GetCheckStructureRequest {
	s.Lang = &v
	return s
}

func (s *GetCheckStructureRequest) SetPageSize(v int32) *GetCheckStructureRequest {
	s.PageSize = &v
	return s
}

func (s *GetCheckStructureRequest) SetRegionId(v string) *GetCheckStructureRequest {
	s.RegionId = &v
	return s
}

func (s *GetCheckStructureRequest) SetTaskSources(v []*string) *GetCheckStructureRequest {
	s.TaskSources = v
	return s
}

func (s *GetCheckStructureRequest) Validate() error {
	return dara.Validate(s)
}
