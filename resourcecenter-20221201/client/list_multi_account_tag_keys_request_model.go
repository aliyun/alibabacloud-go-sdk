// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListMultiAccountTagKeysRequest
	GetMatchType() *string
	SetMaxResults(v int32) *ListMultiAccountTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountTagKeysRequest
	GetNextToken() *string
	SetScope(v string) *ListMultiAccountTagKeysRequest
	GetScope() *string
	SetTagKey(v string) *ListMultiAccountTagKeysRequest
	GetTagKey() *string
}

type ListMultiAccountTagKeysRequest struct {
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
	// The search scope. The value of this parameter can be one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID of the resource directory. The ID is indicated by the `ResourceDirectoryId` parameter in the response of the operation.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID of the Root folder. The ID is indicated by the `RootFolderId` parameter in the response of the operation.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID of the folder. The ID is indicated by the `FolderId` parameter in the response of the operation.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID of the member. The ID is indicated by the `AccountId` parameter in the response of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListMultiAccountTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListMultiAccountTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountTagKeysRequest) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountTagKeysRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListMultiAccountTagKeysRequest) SetMatchType(v string) *ListMultiAccountTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetMaxResults(v int32) *ListMultiAccountTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetNextToken(v string) *ListMultiAccountTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetScope(v string) *ListMultiAccountTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetTagKey(v string) *ListMultiAccountTagKeysRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
