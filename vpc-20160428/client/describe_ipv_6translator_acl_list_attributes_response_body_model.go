// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetAclEntries() *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries
	SetAclId(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetAclId() *string
	SetAclName(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetAclName() *string
	SetPageNumber(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody
	GetTotalCount() *int32
}

type DescribeIPv6TranslatorAclListAttributesResponseBody struct {
	AclEntries *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Struct"`
	// The ACL ID.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// acl1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetAclEntries() *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries {
	return s.AclEntries
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetAclName() *string {
	return s.AclName
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetAclEntries(v *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.AclEntries = v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetAclId(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetAclName(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.AclName = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetPageNumber(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetPageSize(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetRequestId(v string) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) SetTotalCount(v int32) *DescribeIPv6TranslatorAclListAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBody) Validate() error {
	if s.AclEntries != nil {
		if err := s.AclEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries struct {
	AclEntry []*DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry `json:"AclEntry,omitempty" xml:"AclEntry,omitempty" type:"Repeated"`
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) GetAclEntry() []*DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry {
	return s.AclEntry
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) SetAclEntry(v []*DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries {
	s.AclEntry = v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntries) Validate() error {
	if s.AclEntry != nil {
		for _, item := range s.AclEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry struct {
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	AclEntryId      *string `json:"AclEntryId,omitempty" xml:"AclEntryId,omitempty"`
	AclEntryIp      *string `json:"AclEntryIp,omitempty" xml:"AclEntryIp,omitempty"`
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) GetAclEntryId() *string {
	return s.AclEntryId
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) GetAclEntryIp() *string {
	return s.AclEntryIp
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) SetAclEntryComment(v string) *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) SetAclEntryId(v string) *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry {
	s.AclEntryId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) SetAclEntryIp(v string) *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry {
	s.AclEntryIp = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponseBodyAclEntriesAclEntry) Validate() error {
	return dara.Validate(s)
}
