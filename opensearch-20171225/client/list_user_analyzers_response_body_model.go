// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUserAnalyzersResponseBody
	GetRequestId() *string
	SetResult(v []*ListUserAnalyzersResponseBodyResult) *ListUserAnalyzersResponseBody
	GetResult() []*ListUserAnalyzersResponseBodyResult
	SetTotalCount(v int32) *ListUserAnalyzersResponseBody
	GetTotalCount() *int32
}

type ListUserAnalyzersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// For more information, see [UserAnalyzer](https://help.aliyun.com/document_detail/178934.html).
	Result []*ListUserAnalyzersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserAnalyzersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAnalyzersResponseBody) GetResult() []*ListUserAnalyzersResponseBodyResult {
	return s.Result
}

func (s *ListUserAnalyzersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserAnalyzersResponseBody) SetRequestId(v string) *ListUserAnalyzersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetResult(v []*ListUserAnalyzersResponseBodyResult) *ListUserAnalyzersResponseBody {
	s.Result = v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetTotalCount(v int32) *ListUserAnalyzersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserAnalyzersResponseBody) Validate() error {
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

type ListUserAnalyzersResponseBodyResult struct {
	// Indicates whether the application is available.
	//
	// example:
	//
	// false
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The basic analyzer. Valid values:
	//
	// 	- chn_standard: [a common analyzer in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- chn_scene_name: an analyzer for person names in Chinese
	//
	// 	- chn_ecommerce: [an analyzer for E-commerce in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- chn_it_content: [an analyzer for IT content in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- en_min: a small-granularity analyzer in English
	//
	// 	- th_standard: a common analyzer in Thai
	//
	// 	- th_ecommerce: an analyzer for E-commerce in Thai
	//
	// 	- vn_standard: a common analyzer in Vietnamese
	//
	// 	- chn_community_it: an analyzer for IT community content in Chinese
	//
	// 	- chn_ecommerce_general: a common analyzer for the E-commerce industry in Chinese
	//
	// 	- chn_esports_general: a common analyzer for the gaming industry in Chinese
	//
	// 	- chn_edu_question: an analyzer for question search of the education industry in Chinese
	//
	// example:
	//
	// chn_standard
	Business *string `json:"business,omitempty" xml:"business,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The dictionaries that are used by the custom analyzer.
	//
	// For more information, see [UserDict](https://help.aliyun.com/document_detail/178933.html).
	Dicts []*ListUserAnalyzersResponseBodyResultDicts `json:"dicts,omitempty" xml:"dicts,omitempty" type:"Repeated"`
	// The ID of the custom analyzer.
	//
	// example:
	//
	// 1234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the custom analyzer.
	//
	// example:
	//
	// kevin_test2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1588054131
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResult) GetAvailable() *bool {
	return s.Available
}

func (s *ListUserAnalyzersResponseBodyResult) GetBusiness() *string {
	return s.Business
}

func (s *ListUserAnalyzersResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListUserAnalyzersResponseBodyResult) GetDicts() []*ListUserAnalyzersResponseBodyResultDicts {
	return s.Dicts
}

func (s *ListUserAnalyzersResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListUserAnalyzersResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListUserAnalyzersResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListUserAnalyzersResponseBodyResult) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetBusiness(v string) *ListUserAnalyzersResponseBodyResult {
	s.Business = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetCreated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetDicts(v []*ListUserAnalyzersResponseBodyResultDicts) *ListUserAnalyzersResponseBodyResult {
	s.Dicts = v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetId(v string) *ListUserAnalyzersResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetName(v string) *ListUserAnalyzersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) Validate() error {
	if s.Dicts != nil {
		for _, item := range s.Dicts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserAnalyzersResponseBodyResultDicts struct {
	// Indicates whether the application is available.
	//
	// example:
	//
	// false
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The number of intervention entries.
	//
	// example:
	//
	// -1
	EntriesCount *int32 `json:"entriesCount,omitempty" xml:"entriesCount,omitempty"`
	// The maximum number of intervention entries that can be created in the dictionary.
	//
	// example:
	//
	// 4
	EntriesLimit *int32 `json:"entriesLimit,omitempty" xml:"entriesLimit,omitempty"`
	// The ID of the dictionary.
	//
	// example:
	//
	// 123
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type. Valid value:
	//
	// 	- segment
	//
	// example:
	//
	// segment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1588054131
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResultDicts) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResultDicts) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetAvailable() *bool {
	return s.Available
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetCreated() *int32 {
	return s.Created
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetEntriesCount() *int32 {
	return s.EntriesCount
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetEntriesLimit() *int32 {
	return s.EntriesLimit
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetId() *string {
	return s.Id
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetType() *string {
	return s.Type
}

func (s *ListUserAnalyzersResponseBodyResultDicts) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResultDicts {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetCreated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesCount(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesCount = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesLimit(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesLimit = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetId(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetType(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Type = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Updated = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) Validate() error {
	return dara.Validate(s)
}
