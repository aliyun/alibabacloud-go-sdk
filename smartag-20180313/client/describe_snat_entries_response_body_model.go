// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSnatEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnatEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnatEntriesResponseBody
	GetRequestId() *string
	SetSnatEntries(v *DescribeSnatEntriesResponseBodySnatEntries) *DescribeSnatEntriesResponseBody
	GetSnatEntries() *DescribeSnatEntriesResponseBodySnatEntries
	SetTotalCount(v int32) *DescribeSnatEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeSnatEntriesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 674BC3B2-5828-41D5-830E-148EE6CF86C2
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnatEntries *DescribeSnatEntriesResponseBodySnatEntries `json:"SnatEntries,omitempty" xml:"SnatEntries,omitempty" type:"Struct"`
	// The total number of SNAT entries.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnatEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnatEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnatEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnatEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnatEntriesResponseBody) GetSnatEntries() *DescribeSnatEntriesResponseBodySnatEntries {
	return s.SnatEntries
}

func (s *DescribeSnatEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnatEntriesResponseBody) SetPageNumber(v int32) *DescribeSnatEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnatEntriesResponseBody) SetPageSize(v int32) *DescribeSnatEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnatEntriesResponseBody) SetRequestId(v string) *DescribeSnatEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnatEntriesResponseBody) SetSnatEntries(v *DescribeSnatEntriesResponseBodySnatEntries) *DescribeSnatEntriesResponseBody {
	s.SnatEntries = v
	return s
}

func (s *DescribeSnatEntriesResponseBody) SetTotalCount(v int32) *DescribeSnatEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnatEntriesResponseBody) Validate() error {
	if s.SnatEntries != nil {
		if err := s.SnatEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnatEntriesResponseBodySnatEntries struct {
	SnatEntry []*DescribeSnatEntriesResponseBodySnatEntriesSnatEntry `json:"SnatEntry,omitempty" xml:"SnatEntry,omitempty" type:"Repeated"`
}

func (s DescribeSnatEntriesResponseBodySnatEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatEntriesResponseBodySnatEntries) GoString() string {
	return s.String()
}

func (s *DescribeSnatEntriesResponseBodySnatEntries) GetSnatEntry() []*DescribeSnatEntriesResponseBodySnatEntriesSnatEntry {
	return s.SnatEntry
}

func (s *DescribeSnatEntriesResponseBodySnatEntries) SetSnatEntry(v []*DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) *DescribeSnatEntriesResponseBodySnatEntries {
	s.SnatEntry = v
	return s
}

func (s *DescribeSnatEntriesResponseBodySnatEntries) Validate() error {
	if s.SnatEntry != nil {
		for _, item := range s.SnatEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnatEntriesResponseBodySnatEntriesSnatEntry struct {
	CidrBlock  *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SnatIp     *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) GoString() string {
	return s.String()
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) SetCidrBlock(v string) *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) SetCreateTime(v int64) *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry {
	s.CreateTime = &v
	return s
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) SetInstanceId(v string) *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) SetSnatIp(v string) *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatEntriesResponseBodySnatEntriesSnatEntry) Validate() error {
	return dara.Validate(s)
}
