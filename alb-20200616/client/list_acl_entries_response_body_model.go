// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*ListAclEntriesResponseBodyAclEntries) *ListAclEntriesResponseBody
	GetAclEntries() []*ListAclEntriesResponseBodyAclEntries
	SetMaxResults(v int32) *ListAclEntriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAclEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAclEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAclEntriesResponseBody
	GetTotalCount() *int32
}

type ListAclEntriesResponseBody struct {
	// The ACL entries.
	AclEntries []*ListAclEntriesResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBody) GetAclEntries() []*ListAclEntriesResponseBodyAclEntries {
	return s.AclEntries
}

func (s *ListAclEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAclEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAclEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAclEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAclEntriesResponseBody) SetAclEntries(v []*ListAclEntriesResponseBodyAclEntries) *ListAclEntriesResponseBody {
	s.AclEntries = v
	return s
}

func (s *ListAclEntriesResponseBody) SetMaxResults(v int32) *ListAclEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetNextToken(v string) *ListAclEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetRequestId(v string) *ListAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetTotalCount(v int32) *ListAclEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAclEntriesResponseBody) Validate() error {
	if s.AclEntries != nil {
		for _, item := range s.AclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclEntriesResponseBodyAclEntries struct {
	// The description of the ACL entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// test-entry
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The CIDR block for the ACL entry.
	//
	// example:
	//
	// 10.0.1.1/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The status of the ACL entry. Valid values:
	//
	// 	- **Adding**: The ACL entry is being added.
	//
	// 	- **Available**: The ACL entry is added and available.
	//
	// 	- **Removing**: The ACL entry is being removed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclEntriesResponseBodyAclEntries) String() string {
	return dara.Prettify(s)
}

func (s ListAclEntriesResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBodyAclEntries) GetDescription() *string {
	return s.Description
}

func (s *ListAclEntriesResponseBodyAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *ListAclEntriesResponseBodyAclEntries) GetStatus() *string {
	return s.Status
}

func (s *ListAclEntriesResponseBodyAclEntries) SetDescription(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Description = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetEntry(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetStatus(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Status = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) Validate() error {
	return dara.Validate(s)
}
