// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionIds(v []*string) *ListDataIngestionsRequest
	GetDataIngestionIds() []*string
	SetDataIngestionStatus(v string) *ListDataIngestionsRequest
	GetDataIngestionStatus() *string
	SetDataIngestionTemplateIds(v []*string) *ListDataIngestionsRequest
	GetDataIngestionTemplateIds() []*string
	SetLang(v string) *ListDataIngestionsRequest
	GetLang() *string
	SetProductId(v string) *ListDataIngestionsRequest
	GetProductId() *string
	SetRegionId(v string) *ListDataIngestionsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataIngestionsRequest
	GetRoleFor() *int64
}

type ListDataIngestionsRequest struct {
	DataIngestionIds []*string `json:"DataIngestionIds,omitempty" xml:"DataIngestionIds,omitempty" type:"Repeated"`
	// example:
	//
	// enabled。
	DataIngestionStatus      *string   `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	DataIngestionTemplateIds []*string `json:"DataIngestionTemplateIds,omitempty" xml:"DataIngestionTemplateIds,omitempty" type:"Repeated"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataIngestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataIngestionsRequest) GetDataIngestionIds() []*string {
	return s.DataIngestionIds
}

func (s *ListDataIngestionsRequest) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *ListDataIngestionsRequest) GetDataIngestionTemplateIds() []*string {
	return s.DataIngestionTemplateIds
}

func (s *ListDataIngestionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataIngestionsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListDataIngestionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataIngestionsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataIngestionsRequest) SetDataIngestionIds(v []*string) *ListDataIngestionsRequest {
	s.DataIngestionIds = v
	return s
}

func (s *ListDataIngestionsRequest) SetDataIngestionStatus(v string) *ListDataIngestionsRequest {
	s.DataIngestionStatus = &v
	return s
}

func (s *ListDataIngestionsRequest) SetDataIngestionTemplateIds(v []*string) *ListDataIngestionsRequest {
	s.DataIngestionTemplateIds = v
	return s
}

func (s *ListDataIngestionsRequest) SetLang(v string) *ListDataIngestionsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataIngestionsRequest) SetProductId(v string) *ListDataIngestionsRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataIngestionsRequest) SetRegionId(v string) *ListDataIngestionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataIngestionsRequest) SetRoleFor(v int64) *ListDataIngestionsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataIngestionsRequest) Validate() error {
	return dara.Validate(s)
}
