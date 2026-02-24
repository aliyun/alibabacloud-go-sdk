// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesAggregatorIds(v string) *CopyCompliancePacksRequest
	GetDesAggregatorIds() *string
	SetSrcAggregatorId(v string) *CopyCompliancePacksRequest
	GetSrcAggregatorId() *string
	SetSrcCompliancePackIds(v string) *CopyCompliancePacksRequest
	GetSrcCompliancePackIds() *string
}

type CopyCompliancePacksRequest struct {
	DesAggregatorIds *string `json:"DesAggregatorIds,omitempty" xml:"DesAggregatorIds,omitempty"`
	SrcAggregatorId  *string `json:"SrcAggregatorId,omitempty" xml:"SrcAggregatorId,omitempty"`
	// This parameter is required.
	SrcCompliancePackIds *string `json:"SrcCompliancePackIds,omitempty" xml:"SrcCompliancePackIds,omitempty"`
}

func (s CopyCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *CopyCompliancePacksRequest) GetDesAggregatorIds() *string {
	return s.DesAggregatorIds
}

func (s *CopyCompliancePacksRequest) GetSrcAggregatorId() *string {
	return s.SrcAggregatorId
}

func (s *CopyCompliancePacksRequest) GetSrcCompliancePackIds() *string {
	return s.SrcCompliancePackIds
}

func (s *CopyCompliancePacksRequest) SetDesAggregatorIds(v string) *CopyCompliancePacksRequest {
	s.DesAggregatorIds = &v
	return s
}

func (s *CopyCompliancePacksRequest) SetSrcAggregatorId(v string) *CopyCompliancePacksRequest {
	s.SrcAggregatorId = &v
	return s
}

func (s *CopyCompliancePacksRequest) SetSrcCompliancePackIds(v string) *CopyCompliancePacksRequest {
	s.SrcCompliancePackIds = &v
	return s
}

func (s *CopyCompliancePacksRequest) Validate() error {
	return dara.Validate(s)
}
