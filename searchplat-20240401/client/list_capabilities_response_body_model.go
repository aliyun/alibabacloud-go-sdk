// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCapabilitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpCode(v int64) *ListCapabilitiesResponseBody
	GetHttpCode() *int64
	SetMaxResults(v int32) *ListCapabilitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCapabilitiesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCapabilitiesResponseBody
	GetRequestId() *string
	SetResult(v []*ListCapabilitiesResponseBodyResult) *ListCapabilitiesResponseBody
	GetResult() []*ListCapabilitiesResponseBodyResult
	SetStatus(v string) *ListCapabilitiesResponseBody
	GetStatus() *string
	SetTotalCount(v int64) *ListCapabilitiesResponseBody
	GetTotalCount() *int64
}

type ListCapabilitiesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// G5FG/nXfNOQ=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*ListCapabilitiesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCapabilitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCapabilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCapabilitiesResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListCapabilitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCapabilitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCapabilitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCapabilitiesResponseBody) GetResult() []*ListCapabilitiesResponseBodyResult {
	return s.Result
}

func (s *ListCapabilitiesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListCapabilitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCapabilitiesResponseBody) SetHttpCode(v int64) *ListCapabilitiesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListCapabilitiesResponseBody) SetMaxResults(v int32) *ListCapabilitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCapabilitiesResponseBody) SetNextToken(v string) *ListCapabilitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCapabilitiesResponseBody) SetRequestId(v string) *ListCapabilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCapabilitiesResponseBody) SetResult(v []*ListCapabilitiesResponseBodyResult) *ListCapabilitiesResponseBody {
	s.Result = v
	return s
}

func (s *ListCapabilitiesResponseBody) SetStatus(v string) *ListCapabilitiesResponseBody {
	s.Status = &v
	return s
}

func (s *ListCapabilitiesResponseBody) SetTotalCount(v int64) *ListCapabilitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCapabilitiesResponseBody) Validate() error {
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

type ListCapabilitiesResponseBodyResult struct {
	// The creation time.
	//
	// example:
	//
	// 1729665694
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// Indicates whether the configuration is the default configuration.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// The configuration category.
	//
	// example:
	//
	// ai_search_agent
	ItemCategory *string `json:"itemCategory,omitempty" xml:"itemCategory,omitempty"`
	// The configuration description.
	//
	// example:
	//
	// 描述
	ItemDesc *string `json:"itemDesc,omitempty" xml:"itemDesc,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// es_knowledge_base
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// itemValue
	ItemValue map[string]interface{} `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// status
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// updated
	//
	// example:
	//
	// 1729665694
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListCapabilitiesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCapabilitiesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCapabilitiesResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ListCapabilitiesResponseBodyResult) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListCapabilitiesResponseBodyResult) GetItemCategory() *string {
	return s.ItemCategory
}

func (s *ListCapabilitiesResponseBodyResult) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *ListCapabilitiesResponseBodyResult) GetItemName() *string {
	return s.ItemName
}

func (s *ListCapabilitiesResponseBodyResult) GetItemValue() map[string]interface{} {
	return s.ItemValue
}

func (s *ListCapabilitiesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListCapabilitiesResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ListCapabilitiesResponseBodyResult) SetCreated(v int64) *ListCapabilitiesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetIsDefault(v bool) *ListCapabilitiesResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetItemCategory(v string) *ListCapabilitiesResponseBodyResult {
	s.ItemCategory = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetItemDesc(v string) *ListCapabilitiesResponseBodyResult {
	s.ItemDesc = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetItemName(v string) *ListCapabilitiesResponseBodyResult {
	s.ItemName = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetItemValue(v map[string]interface{}) *ListCapabilitiesResponseBodyResult {
	s.ItemValue = v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetStatus(v string) *ListCapabilitiesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) SetUpdated(v int64) *ListCapabilitiesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListCapabilitiesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
