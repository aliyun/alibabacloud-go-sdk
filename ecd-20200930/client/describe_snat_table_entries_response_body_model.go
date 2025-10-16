// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeSnatTableEntriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnatTableEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSnatTableEntriesResponseBody
	GetRequestId() *string
	SetSnatTableEntries(v []*DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody
	GetSnatTableEntries() []*DescribeSnatTableEntriesResponseBodySnatTableEntries
}

type DescribeSnatTableEntriesResponseBody struct {
	MaxResults       *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnatTableEntries []*DescribeSnatTableEntriesResponseBodySnatTableEntries `json:"SnatTableEntries,omitempty" xml:"SnatTableEntries,omitempty" type:"Repeated"`
}

func (s DescribeSnatTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnatTableEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnatTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnatTableEntriesResponseBody) GetSnatTableEntries() []*DescribeSnatTableEntriesResponseBodySnatTableEntries {
	return s.SnatTableEntries
}

func (s *DescribeSnatTableEntriesResponseBody) SetMaxResults(v int32) *DescribeSnatTableEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetNextToken(v string) *DescribeSnatTableEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetRequestId(v string) *DescribeSnatTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetSnatTableEntries(v []*DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody {
	s.SnatTableEntries = v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) Validate() error {
	if s.SnatTableEntries != nil {
		for _, item := range s.SnatTableEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnatTableEntriesResponseBodySnatTableEntries struct {
	EipAffinity     *string `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	NatGatewayId    *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SnatEntryId     *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	SnatEntryName   *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	SnatIp          *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	SnatTableId     *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	SourceCIDR      *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	SourceVSwitchId *string `json:"SourceVSwitchId,omitempty" xml:"SourceVSwitchId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetEipAffinity() *string {
	return s.EipAffinity
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSourceVSwitchId() *string {
	return s.SourceVSwitchId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetEipAffinity(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.EipAffinity = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetNatGatewayId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatEntryId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatEntryName(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatIp(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatTableId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatTableId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSourceCIDR(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSourceVSwitchId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SourceVSwitchId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetStatus(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.Status = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) Validate() error {
	return dara.Validate(s)
}
