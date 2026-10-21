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
	SetNormalizationSchemaIds(v []*string) *ListDataIngestionsRequest
	GetNormalizationSchemaIds() []*string
	SetProductId(v string) *ListDataIngestionsRequest
	GetProductId() *string
	SetRegionId(v string) *ListDataIngestionsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataIngestionsRequest
	GetRoleFor() *int64
}

type ListDataIngestionsRequest struct {
	// The list of data ingestion policy IDs.
	DataIngestionIds []*string `json:"DataIngestionIds,omitempty" xml:"DataIngestionIds,omitempty" type:"Repeated"`
	// The accessed state of data ingestion. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	DataIngestionStatus *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// The list of data ingestion template IDs.
	DataIngestionTemplateIds []*string `json:"DataIngestionTemplateIds,omitempty" xml:"DataIngestionTemplateIds,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of normalization schema IDs.
	NormalizationSchemaIds []*string `json:"NormalizationSchemaIds,omitempty" xml:"NormalizationSchemaIds,omitempty" type:"Repeated"`
	// The product ID.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 173326*******
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

func (s *ListDataIngestionsRequest) GetNormalizationSchemaIds() []*string {
	return s.NormalizationSchemaIds
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

func (s *ListDataIngestionsRequest) SetNormalizationSchemaIds(v []*string) *ListDataIngestionsRequest {
	s.NormalizationSchemaIds = v
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
