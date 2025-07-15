// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*DescribeAclEntriesResponseBodyAclEntries) *DescribeAclEntriesResponseBody
	GetAclEntries() []*DescribeAclEntriesResponseBodyAclEntries
	SetNextToken(v string) *DescribeAclEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAclEntriesResponseBody
	GetRequestId() *string
}

type DescribeAclEntriesResponseBody struct {
	// The ACL entries.
	AclEntries []*DescribeAclEntriesResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The token that is used to start the next query. If the value of this parameter is empty, all results are returned.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAclEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclEntriesResponseBody) GetAclEntries() []*DescribeAclEntriesResponseBodyAclEntries {
	return s.AclEntries
}

func (s *DescribeAclEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAclEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclEntriesResponseBody) SetAclEntries(v []*DescribeAclEntriesResponseBodyAclEntries) *DescribeAclEntriesResponseBody {
	s.AclEntries = v
	return s
}

func (s *DescribeAclEntriesResponseBody) SetNextToken(v string) *DescribeAclEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAclEntriesResponseBody) SetRequestId(v string) *DescribeAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAclEntriesResponseBodyAclEntries struct {
	// The ACL type.
	//
	// Valid values:
	//
	// 	- allow: whitelist
	//
	// 	- disable: blacklist
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the instance to which the ACL applies. You can specify an office network ID or a cloud computer ID.
	//
	// example:
	//
	// ecd-fsafeweh***
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The granularity of the ACL.
	//
	// Valid values:
	//
	// 	- desktop: cloud computer
	//
	// 	- vpc: office network
	//
	// example:
	//
	// desktop
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeAclEntriesResponseBodyAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclEntriesResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeAclEntriesResponseBodyAclEntries) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeAclEntriesResponseBodyAclEntries) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeAclEntriesResponseBodyAclEntries) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAclEntriesResponseBodyAclEntries) SetPolicy(v string) *DescribeAclEntriesResponseBodyAclEntries {
	s.Policy = &v
	return s
}

func (s *DescribeAclEntriesResponseBodyAclEntries) SetSourceId(v string) *DescribeAclEntriesResponseBodyAclEntries {
	s.SourceId = &v
	return s
}

func (s *DescribeAclEntriesResponseBodyAclEntries) SetSourceType(v string) *DescribeAclEntriesResponseBodyAclEntries {
	s.SourceType = &v
	return s
}

func (s *DescribeAclEntriesResponseBodyAclEntries) Validate() error {
	return dara.Validate(s)
}
