// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnatEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnatEntries(v *DescribeDnatEntriesResponseBodyDnatEntries) *DescribeDnatEntriesResponseBody
	GetDnatEntries() *DescribeDnatEntriesResponseBodyDnatEntries
	SetPageNumber(v int32) *DescribeDnatEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnatEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDnatEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDnatEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeDnatEntriesResponseBody struct {
	DnatEntries *DescribeDnatEntriesResponseBodyDnatEntries `json:"DnatEntries,omitempty" xml:"DnatEntries,omitempty" type:"Struct"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 635640CA-2335-4856-A9CB-1CB5C444DC5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDnatEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnatEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBody) GetDnatEntries() *DescribeDnatEntriesResponseBodyDnatEntries {
	return s.DnatEntries
}

func (s *DescribeDnatEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnatEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnatEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnatEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDnatEntriesResponseBody) SetDnatEntries(v *DescribeDnatEntriesResponseBodyDnatEntries) *DescribeDnatEntriesResponseBody {
	s.DnatEntries = v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetPageNumber(v int32) *DescribeDnatEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetPageSize(v int32) *DescribeDnatEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetRequestId(v string) *DescribeDnatEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetTotalCount(v int32) *DescribeDnatEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) Validate() error {
	if s.DnatEntries != nil {
		if err := s.DnatEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnatEntriesResponseBodyDnatEntries struct {
	DnatEntry []*DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry `json:"DnatEntry,omitempty" xml:"DnatEntry,omitempty" type:"Repeated"`
}

func (s DescribeDnatEntriesResponseBodyDnatEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnatEntriesResponseBodyDnatEntries) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBodyDnatEntries) GetDnatEntry() []*DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	return s.DnatEntry
}

func (s *DescribeDnatEntriesResponseBodyDnatEntries) SetDnatEntry(v []*DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) *DescribeDnatEntriesResponseBodyDnatEntries {
	s.DnatEntry = v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntries) Validate() error {
	if s.DnatEntry != nil {
		for _, item := range s.DnatEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry struct {
	DnatEntryId  *string `json:"DnatEntryId,omitempty" xml:"DnatEntryId,omitempty"`
	ExternalIp   *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	InternalIp   *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	SagId        *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetDnatEntryId() *string {
	return s.DnatEntryId
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetSagId() *string {
	return s.SagId
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) GetType() *string {
	return s.Type
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetDnatEntryId(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.DnatEntryId = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetExternalIp(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.ExternalIp = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetExternalPort(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.ExternalPort = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetInternalIp(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.InternalIp = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetInternalPort(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.InternalPort = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetIpProtocol(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.IpProtocol = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetSagId(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.SagId = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) SetType(v string) *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry {
	s.Type = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyDnatEntriesDnatEntry) Validate() error {
	return dara.Validate(s)
}
