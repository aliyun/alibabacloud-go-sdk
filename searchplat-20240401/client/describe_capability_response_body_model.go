// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpCode(v int64) *DescribeCapabilityResponseBody
	GetHttpCode() *int64
	SetRequestId(v string) *DescribeCapabilityResponseBody
	GetRequestId() *string
	SetResult(v *DescribeCapabilityResponseBodyResult) *DescribeCapabilityResponseBody
	GetResult() *DescribeCapabilityResponseBodyResult
	SetStatus(v string) *DescribeCapabilityResponseBody
	GetStatus() *string
}

type DescribeCapabilityResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Request ID
	//
	// example:
	//
	// 0E3D5E2B-B63A-4445-B359-329CC07255EA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response result
	Result *DescribeCapabilityResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// Request status
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapabilityResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DescribeCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCapabilityResponseBody) GetResult() *DescribeCapabilityResponseBodyResult {
	return s.Result
}

func (s *DescribeCapabilityResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCapabilityResponseBody) SetHttpCode(v int64) *DescribeCapabilityResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DescribeCapabilityResponseBody) SetRequestId(v string) *DescribeCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapabilityResponseBody) SetResult(v *DescribeCapabilityResponseBodyResult) *DescribeCapabilityResponseBody {
	s.Result = v
	return s
}

func (s *DescribeCapabilityResponseBody) SetStatus(v string) *DescribeCapabilityResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCapabilityResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapabilityResponseBodyResult struct {
	// Timestamp of creation time
	//
	// example:
	//
	// 1745893195510
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// Whether it is the default configuration
	//
	// example:
	//
	// false
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Configuration category
	//
	// example:
	//
	// ai_search_agent
	ItemCategory *string `json:"itemCategory,omitempty" xml:"itemCategory,omitempty"`
	// Configuration description
	//
	// example:
	//
	// 描述
	ItemDesc *string `json:"itemDesc,omitempty" xml:"itemDesc,omitempty"`
	// Configuration name
	//
	// example:
	//
	// es_knowledge_base
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// An object containing information such as endpoint and function, which describes the detailed configuration of the knowledge base.
	ItemValue map[string]interface{} `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// Status
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Update timestamp
	//
	// example:
	//
	// 1729665694
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeCapabilityResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapabilityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeCapabilityResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *DescribeCapabilityResponseBodyResult) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeCapabilityResponseBodyResult) GetItemCategory() *string {
	return s.ItemCategory
}

func (s *DescribeCapabilityResponseBodyResult) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *DescribeCapabilityResponseBodyResult) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeCapabilityResponseBodyResult) GetItemValue() map[string]interface{} {
	return s.ItemValue
}

func (s *DescribeCapabilityResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeCapabilityResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *DescribeCapabilityResponseBodyResult) SetCreated(v int64) *DescribeCapabilityResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetIsDefault(v bool) *DescribeCapabilityResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetItemCategory(v string) *DescribeCapabilityResponseBodyResult {
	s.ItemCategory = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetItemDesc(v string) *DescribeCapabilityResponseBodyResult {
	s.ItemDesc = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetItemName(v string) *DescribeCapabilityResponseBodyResult {
	s.ItemName = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetItemValue(v map[string]interface{}) *DescribeCapabilityResponseBodyResult {
	s.ItemValue = v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetStatus(v string) *DescribeCapabilityResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) SetUpdated(v int64) *DescribeCapabilityResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeCapabilityResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
