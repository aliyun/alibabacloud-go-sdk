// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeSnatTableEntriesRequest
	GetMaxResults() *int32
	SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest
	GetNatGatewayId() *string
	SetNextToken(v string) *DescribeSnatTableEntriesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeSnatTableEntriesRequest
	GetRegionId() *string
	SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest
	GetSnatEntryId() *string
	SetSnatTableId(v string) *DescribeSnatTableEntriesRequest
	GetSnatTableId() *string
}

type DescribeSnatTableEntriesRequest struct {
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
}

func (s DescribeSnatTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnatTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnatTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnatTableEntriesRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DescribeSnatTableEntriesRequest) SetMaxResults(v int32) *DescribeSnatTableEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetNextToken(v string) *DescribeSnatTableEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetRegionId(v string) *DescribeSnatTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatTableId(v string) *DescribeSnatTableEntriesRequest {
	s.SnatTableId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
