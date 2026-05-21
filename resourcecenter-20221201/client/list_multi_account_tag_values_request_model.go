// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListMultiAccountTagValuesRequest
	GetMatchType() *string
	SetMaxResults(v int32) *ListMultiAccountTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountTagValuesRequest
	GetNextToken() *string
	SetScope(v string) *ListMultiAccountTagValuesRequest
	GetScope() *string
	SetTagKey(v string) *ListMultiAccountTagValuesRequest
	GetTagKey() *string
	SetTagValue(v string) *ListMultiAccountTagValuesRequest
	GetTagValue() *string
}

type ListMultiAccountTagValuesRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListMultiAccountTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListMultiAccountTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountTagValuesRequest) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListMultiAccountTagValuesRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *ListMultiAccountTagValuesRequest) SetMatchType(v string) *ListMultiAccountTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetMaxResults(v int32) *ListMultiAccountTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetNextToken(v string) *ListMultiAccountTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetScope(v string) *ListMultiAccountTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagKey(v string) *ListMultiAccountTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagValue(v string) *ListMultiAccountTagValuesRequest {
	s.TagValue = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
