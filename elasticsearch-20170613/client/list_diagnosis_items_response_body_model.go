// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDiagnosisItemsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDiagnosisItemsResponseBodyResult) *ListDiagnosisItemsResponseBody
	GetResult() []*ListDiagnosisItemsResponseBodyResult
}

type ListDiagnosisItemsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result []*ListDiagnosisItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnosisItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosisItemsResponseBody) GetResult() []*ListDiagnosisItemsResponseBodyResult {
	return s.Result
}

func (s *ListDiagnosisItemsResponseBody) SetRequestId(v string) *ListDiagnosisItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisItemsResponseBody) SetResult(v []*ListDiagnosisItemsResponseBodyResult) *ListDiagnosisItemsResponseBody {
	s.Result = v
	return s
}

func (s *ListDiagnosisItemsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDiagnosisItemsResponseBodyResult struct {
	// Indicates whether billable tokens are consumed. The value is true when level is ADVANCED.
	//
	// example:
	//
	// true
	Billable *bool `json:"billable,omitempty" xml:"billable,omitempty"`
	// The category code. You can use this value to group diagnostic items by category.
	//
	// example:
	//
	// CLUSTER_HEALTH
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The diagnostic item description.
	//
	// example:
	//
	// Diagnoses whether data write operations are backlogged in the cluster. When data write operations are backlogged, BulkReject exceptions occur, which may cause data loss and severe system resource consumption
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the cluster API is accessed.
	//
	// example:
	//
	// true
	EsApiRequired *bool `json:"esApiRequired,omitempty" xml:"esApiRequired,omitempty"`
	// The diagnostic item identifier.
	//
	// example:
	//
	// ClusterBulkRejectDiagnostic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The diagnostic item level. Valid values:
	//
	// - BASIC: basic inspection item (free).
	//
	// - ADVANCED: advanced inspection item (consumes billable tokens).
	//
	// example:
	//
	// BASIC
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The diagnostic item name.
	//
	// example:
	//
	// Index Write BulkReject Diagnostics
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The sort order number for display.
	//
	// example:
	//
	// 1
	SortOrder *int32 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The supported execution modes. Basic items support RULE and AGENT. Advanced items support only AGENT.
	SupportedModes []*string `json:"supportedModes,omitempty" xml:"supportedModes,omitempty" type:"Repeated"`
}

func (s ListDiagnosisItemsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsResponseBodyResult) GetBillable() *bool {
	return s.Billable
}

func (s *ListDiagnosisItemsResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *ListDiagnosisItemsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListDiagnosisItemsResponseBodyResult) GetEsApiRequired() *bool {
	return s.EsApiRequired
}

func (s *ListDiagnosisItemsResponseBodyResult) GetKey() *string {
	return s.Key
}

func (s *ListDiagnosisItemsResponseBodyResult) GetLevel() *string {
	return s.Level
}

func (s *ListDiagnosisItemsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDiagnosisItemsResponseBodyResult) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *ListDiagnosisItemsResponseBodyResult) GetSupportedModes() []*string {
	return s.SupportedModes
}

func (s *ListDiagnosisItemsResponseBodyResult) SetBillable(v bool) *ListDiagnosisItemsResponseBodyResult {
	s.Billable = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetCategory(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetDescription(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetEsApiRequired(v bool) *ListDiagnosisItemsResponseBodyResult {
	s.EsApiRequired = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetKey(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Key = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetLevel(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Level = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetName(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetSortOrder(v int32) *ListDiagnosisItemsResponseBodyResult {
	s.SortOrder = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetSupportedModes(v []*string) *ListDiagnosisItemsResponseBodyResult {
	s.SupportedModes = v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
