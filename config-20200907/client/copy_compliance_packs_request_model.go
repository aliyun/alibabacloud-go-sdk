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
	// The IDs of the destination account groups into which the compliance packages are replicated. Separate multiple account group IDs with commas (,).
	//
	// > If this parameter is left empty, the compliance packages are replicated to the current account group.
	//
	// example:
	//
	// ca-c73c626622af00f8****
	DesAggregatorIds *string `json:"DesAggregatorIds,omitempty" xml:"DesAggregatorIds,omitempty"`
	// The ID of the account group to which the compliance packages belong.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// example:
	//
	// ca-05e6626622af0050****
	SrcAggregatorId *string `json:"SrcAggregatorId,omitempty" xml:"SrcAggregatorId,omitempty"`
	// The IDs of the compliance packages. Separate multiple compliance package IDs with commas (,).
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-4c02626622af0050****,cp-47c1626622af0050****
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
