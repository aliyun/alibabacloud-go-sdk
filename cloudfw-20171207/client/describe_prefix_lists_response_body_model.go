// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixList(v []*DescribePrefixListsResponseBodyPrefixList) *DescribePrefixListsResponseBody
	GetPrefixList() []*DescribePrefixListsResponseBodyPrefixList
	SetRequestId(v string) *DescribePrefixListsResponseBody
	GetRequestId() *string
}

type DescribePrefixListsResponseBody struct {
	// Details about the prefix lists.
	PrefixList []*DescribePrefixListsResponseBodyPrefixList `json:"PrefixList,omitempty" xml:"PrefixList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 67FD76C2-C493-5815-8107-643FD7AB77C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrefixListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBody) GetPrefixList() []*DescribePrefixListsResponseBodyPrefixList {
	return s.PrefixList
}

func (s *DescribePrefixListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrefixListsResponseBody) SetPrefixList(v []*DescribePrefixListsResponseBodyPrefixList) *DescribePrefixListsResponseBody {
	s.PrefixList = v
	return s
}

func (s *DescribePrefixListsResponseBody) SetRequestId(v string) *DescribePrefixListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrefixListsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePrefixListsResponseBodyPrefixList struct {
	// The IP address family of the prefix list. Valid values:
	//
	// 	- IPv4
	//
	// 	- IPv6
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The number of associated resources.
	//
	// example:
	//
	// 2
	AssociationCount *int32 `json:"AssociationCount,omitempty" xml:"AssociationCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-10-16T08:31:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// TCP_14900-14911
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries in the prefix list.
	//
	// example:
	//
	// 20
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-uf64nco3ujjqchx6aaji
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list.
	//
	// example:
	//
	// participant-name
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
}

func (s DescribePrefixListsResponseBodyPrefixList) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixList) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetAssociationCount() *int32 {
	return s.AssociationCount
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetDescription() *string {
	return s.Description
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DescribePrefixListsResponseBodyPrefixList) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetAddressFamily(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.AddressFamily = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetAssociationCount(v int32) *DescribePrefixListsResponseBodyPrefixList {
	s.AssociationCount = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetCreationTime(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.CreationTime = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetDescription(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.Description = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetMaxEntries(v int32) *DescribePrefixListsResponseBodyPrefixList {
	s.MaxEntries = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetPrefixListId(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetPrefixListName(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.PrefixListName = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) Validate() error {
	return dara.Validate(s)
}
