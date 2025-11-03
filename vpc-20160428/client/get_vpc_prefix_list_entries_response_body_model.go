// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GetVpcPrefixListEntriesResponseBody
	GetCount() *int64
	SetNextToken(v string) *GetVpcPrefixListEntriesResponseBody
	GetNextToken() *string
	SetPrefixListEntry(v []*GetVpcPrefixListEntriesResponseBodyPrefixListEntry) *GetVpcPrefixListEntriesResponseBody
	GetPrefixListEntry() []*GetVpcPrefixListEntriesResponseBodyPrefixListEntry
	SetRequestId(v string) *GetVpcPrefixListEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetVpcPrefixListEntriesResponseBody
	GetTotalCount() *int64
}

type GetVpcPrefixListEntriesResponseBody struct {
	// The number of entries.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value indicates the token that is used for the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the prefix list.
	PrefixListEntry []*GetVpcPrefixListEntriesResponseBodyPrefixListEntry `json:"PrefixListEntry,omitempty" xml:"PrefixListEntry,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVpcPrefixListEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListEntriesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *GetVpcPrefixListEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVpcPrefixListEntriesResponseBody) GetPrefixListEntry() []*GetVpcPrefixListEntriesResponseBodyPrefixListEntry {
	return s.PrefixListEntry
}

func (s *GetVpcPrefixListEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcPrefixListEntriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetVpcPrefixListEntriesResponseBody) SetCount(v int64) *GetVpcPrefixListEntriesResponseBody {
	s.Count = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBody) SetNextToken(v string) *GetVpcPrefixListEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBody) SetPrefixListEntry(v []*GetVpcPrefixListEntriesResponseBodyPrefixListEntry) *GetVpcPrefixListEntriesResponseBody {
	s.PrefixListEntry = v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBody) SetRequestId(v string) *GetVpcPrefixListEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBody) SetTotalCount(v int64) *GetVpcPrefixListEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBody) Validate() error {
	if s.PrefixListEntry != nil {
		for _, item := range s.PrefixListEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVpcPrefixListEntriesResponseBodyPrefixListEntry struct {
	// The CIDR blocks specified in the prefix list.
	//
	// example:
	//
	// 192.168.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the prefix list.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVpcPrefixListEntriesResponseBodyPrefixListEntry) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListEntriesResponseBodyPrefixListEntry) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) GetCidr() *string {
	return s.Cidr
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) GetDescription() *string {
	return s.Description
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) SetCidr(v string) *GetVpcPrefixListEntriesResponseBodyPrefixListEntry {
	s.Cidr = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) SetDescription(v string) *GetVpcPrefixListEntriesResponseBodyPrefixListEntry {
	s.Description = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) SetPrefixListId(v string) *GetVpcPrefixListEntriesResponseBodyPrefixListEntry {
	s.PrefixListId = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) SetRegionId(v string) *GetVpcPrefixListEntriesResponseBodyPrefixListEntry {
	s.RegionId = &v
	return s
}

func (s *GetVpcPrefixListEntriesResponseBodyPrefixListEntry) Validate() error {
	return dara.Validate(s)
}
