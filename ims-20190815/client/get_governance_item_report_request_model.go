// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceItemReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGovernanceItemType(v string) *GetGovernanceItemReportRequest
	GetGovernanceItemType() *string
	SetMarker(v string) *GetGovernanceItemReportRequest
	GetMarker() *string
	SetMaxItems(v string) *GetGovernanceItemReportRequest
	GetMaxItems() *string
}

type GetGovernanceItemReportRequest struct {
	// example:
	//
	// SSOLoginEnabled
	GovernanceItemType *string `json:"GovernanceItemType,omitempty" xml:"GovernanceItemType,omitempty"`
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// 1000
	MaxItems *string `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s GetGovernanceItemReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportRequest) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportRequest) GetGovernanceItemType() *string {
	return s.GovernanceItemType
}

func (s *GetGovernanceItemReportRequest) GetMarker() *string {
	return s.Marker
}

func (s *GetGovernanceItemReportRequest) GetMaxItems() *string {
	return s.MaxItems
}

func (s *GetGovernanceItemReportRequest) SetGovernanceItemType(v string) *GetGovernanceItemReportRequest {
	s.GovernanceItemType = &v
	return s
}

func (s *GetGovernanceItemReportRequest) SetMarker(v string) *GetGovernanceItemReportRequest {
	s.Marker = &v
	return s
}

func (s *GetGovernanceItemReportRequest) SetMaxItems(v string) *GetGovernanceItemReportRequest {
	s.MaxItems = &v
	return s
}

func (s *GetGovernanceItemReportRequest) Validate() error {
	return dara.Validate(s)
}
