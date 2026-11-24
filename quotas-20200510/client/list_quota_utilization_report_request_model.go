// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaUtilizationReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustable(v bool) *ListQuotaUtilizationReportRequest
	GetAdjustable() *bool
	SetKeyWords(v string) *ListQuotaUtilizationReportRequest
	GetKeyWords() *string
	SetMaxResults(v int32) *ListQuotaUtilizationReportRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaUtilizationReportRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaUtilizationReportRequest
	GetProductCode() *string
	SetQuotaCategory(v string) *ListQuotaUtilizationReportRequest
	GetQuotaCategory() *string
	SetQuotaRegionId(v string) *ListQuotaUtilizationReportRequest
	GetQuotaRegionId() *string
	SetUtilizationStatuses(v []*string) *ListQuotaUtilizationReportRequest
	GetUtilizationStatuses() []*string
}

type ListQuotaUtilizationReportRequest struct {
	// Specifies whether to filter quotas by adjustability. Valid values:
	//
	// - true: Queries only adjustable quotas.
	//
	// - false: Queries only non-adjustable quotas.
	//
	// If this parameter is not specified, quotas are not filtered by adjustability.
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The search keyword. Performs a fuzzy match based on the quota name or quota ID.
	//
	// example:
	//
	// abc
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// The maximum number of entries to return for this query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first query. For subsequent queries, pass the NextToken value returned in the previous response as is, and keep the quota type and filter conditions unchanged. Do not parse or construct the token.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwLCJvdHNOZXh0VG9rZW4iOiJDQUVTR3dvWkNoY0tFM1J2ZEdGc1gzVnpZV2RsWDNCbGNtTmxiblFRQVJnQkl1RUJDZ2tCQUFBQUFBQUFXVUFLMHdFRHpnQUFBREZUTXpFek9ETXdNemN6TnpNNE16VXpPVE0xTXpjek5qTXpNekV6TlRNNE16Z3VVek01TXpjek5UTTJOalkyTWpZeE16Y3laRFl5TmpJek16WXhNbVF6TkRNMk16RXpOREprTmpFek9UTTRNemN5WkRZMk16WTJORE16TXpVMk16TXhNemt6TURZMk16RXpNUzVUTmpFMk16YzBOamsyWmpabE56UTNNall4TmprMll5NVRNelkyTWpZMU16a3pPRE0yTmpRek1qSmtNekV6TlRNeU5qRXlaRE0wTXpFek1UWTJNbVEyTWpZME16TXpNVEprTXpNMk16WXlOalV6TkRNMU16VXpNalkwTXpVek1qTTMiLCJ0YXNrSWQiOiI5NzU2ZmJhNy1iYjNhLTQ2MTQtYTk4Ny1mNmQzNWMxOTBmMTEiLCJ1aWQiOiIxODA3Nzg1OTU3NjMxNTg4IiwidXRpbGl6YXRpb25TdGF0dXNlcyI6IiJ9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the cloud service.
	//
	// > For information about how to obtain the abbreviation of a cloud service, see ProductCode in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// CommonQuota, FlowControl
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The region ID of the quota, such as cn-hangzhou. If this parameter is not specified or is left empty, quotas are not filtered by region.
	//
	// example:
	//
	// cn-hangzhou
	QuotaRegionId *string `json:"QuotaRegionId,omitempty" xml:"QuotaRegionId,omitempty"`
	// The utilization statuses. You can specify multiple values. If this parameter is not specified or an empty list is passed, quotas are not filtered by utilization status.
	UtilizationStatuses []*string `json:"UtilizationStatuses,omitempty" xml:"UtilizationStatuses,omitempty" type:"Repeated"`
}

func (s ListQuotaUtilizationReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportRequest) GetAdjustable() *bool {
	return s.Adjustable
}

func (s *ListQuotaUtilizationReportRequest) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ListQuotaUtilizationReportRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaUtilizationReportRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaUtilizationReportRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaUtilizationReportRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaUtilizationReportRequest) GetQuotaRegionId() *string {
	return s.QuotaRegionId
}

func (s *ListQuotaUtilizationReportRequest) GetUtilizationStatuses() []*string {
	return s.UtilizationStatuses
}

func (s *ListQuotaUtilizationReportRequest) SetAdjustable(v bool) *ListQuotaUtilizationReportRequest {
	s.Adjustable = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetKeyWords(v string) *ListQuotaUtilizationReportRequest {
	s.KeyWords = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetMaxResults(v int32) *ListQuotaUtilizationReportRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetNextToken(v string) *ListQuotaUtilizationReportRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetProductCode(v string) *ListQuotaUtilizationReportRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetQuotaCategory(v string) *ListQuotaUtilizationReportRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetQuotaRegionId(v string) *ListQuotaUtilizationReportRequest {
	s.QuotaRegionId = &v
	return s
}

func (s *ListQuotaUtilizationReportRequest) SetUtilizationStatuses(v []*string) *ListQuotaUtilizationReportRequest {
	s.UtilizationStatuses = v
	return s
}

func (s *ListQuotaUtilizationReportRequest) Validate() error {
	return dara.Validate(s)
}
