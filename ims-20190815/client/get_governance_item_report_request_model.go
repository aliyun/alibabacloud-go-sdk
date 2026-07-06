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
	// The check item. For more information, see [Identity and access governance check items](https://help.aliyun.com/zh/ram/user-guide/overview-of-cloud-governance-for-ram?spm=a2c4g.11174283.0.0.88b3de53tfL5XG#section-q06-p9p-8vl).
	//
	// example:
	//
	// SSOLoginEnabled
	GovernanceItemType *string `json:"GovernanceItemType,omitempty" xml:"GovernanceItemType,omitempty"`
	// If the response is truncated, use the `Marker` to retrieve the subsequent content.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If the response is truncated because the number of entries exceeds the value of `MaxItems`, the value of the `IsTruncated` parameter is true.
	//
	// Valid values: 1 to 1000. Default value: 1000.
	//
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
