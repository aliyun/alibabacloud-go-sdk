// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *ListAclEntriesRequest
	GetAclId() *string
	SetMaxResults(v int32) *ListAclEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAclEntriesRequest
	GetNextToken() *string
}

type ListAclEntriesRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAclEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListAclEntriesRequest) GetAclId() *string {
	return s.AclId
}

func (s *ListAclEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAclEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAclEntriesRequest) SetAclId(v string) *ListAclEntriesRequest {
	s.AclId = &v
	return s
}

func (s *ListAclEntriesRequest) SetMaxResults(v int32) *ListAclEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesRequest) SetNextToken(v string) *ListAclEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclEntriesRequest) Validate() error {
	return dara.Validate(s)
}
