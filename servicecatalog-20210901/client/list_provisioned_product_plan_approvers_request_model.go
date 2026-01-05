// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlanApproversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevelFilter(v string) *ListProvisionedProductPlanApproversRequest
	GetAccessLevelFilter() *string
	SetApprovalFilter(v string) *ListProvisionedProductPlanApproversRequest
	GetApprovalFilter() *string
	SetFilters(v []*ListProvisionedProductPlanApproversRequestFilters) *ListProvisionedProductPlanApproversRequest
	GetFilters() []*ListProvisionedProductPlanApproversRequestFilters
}

type ListProvisionedProductPlanApproversRequest struct {
	// The access filter. Valid values:
	//
	// 	- User (default): queries the plans that are created by the current requester.
	//
	// 	- Account: queries the plans that belong to the current Alibaba Cloud account.
	//
	// 	- ResourceDirectory: queries the plans that belong to the current resource directory.
	//
	// >  You must specify one of the `ApprovalFilter` and `AccessLevelFilter` parameters, but not both.
	//
	// example:
	//
	// User
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The access filter of the review dimension. Valid values:
	//
	// 	- AccountRequests: queries all reviewed plans that belong to the current Alibaba Cloud account.
	//
	// 	- ResourceDirectoryRequests: queries all plans that belong to the current resource directory.
	//
	// >  You must specify one of the `ApprovalFilter` and `AccessLevelFilter` parameters, but not both.
	//
	// example:
	//
	// AccountRequests
	ApprovalFilter *string `json:"ApprovalFilter,omitempty" xml:"ApprovalFilter,omitempty"`
	// An array that consists of filter conditions.
	Filters []*ListProvisionedProductPlanApproversRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
}

func (s ListProvisionedProductPlanApproversRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlanApproversRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversRequest) GetAccessLevelFilter() *string {
	return s.AccessLevelFilter
}

func (s *ListProvisionedProductPlanApproversRequest) GetApprovalFilter() *string {
	return s.ApprovalFilter
}

func (s *ListProvisionedProductPlanApproversRequest) GetFilters() []*ListProvisionedProductPlanApproversRequestFilters {
	return s.Filters
}

func (s *ListProvisionedProductPlanApproversRequest) SetAccessLevelFilter(v string) *ListProvisionedProductPlanApproversRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequest) SetApprovalFilter(v string) *ListProvisionedProductPlanApproversRequest {
	s.ApprovalFilter = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequest) SetFilters(v []*ListProvisionedProductPlanApproversRequestFilters) *ListProvisionedProductPlanApproversRequest {
	s.Filters = v
	return s
}

func (s *ListProvisionedProductPlanApproversRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProvisionedProductPlanApproversRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// 	- ProvisionedProductPlanApproverName: performs fuzzy match by reviewer name.
	//
	// example:
	//
	// ProvisionedProductPlanApproverName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// approver
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductPlanApproversRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlanApproversRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListProvisionedProductPlanApproversRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListProvisionedProductPlanApproversRequestFilters) SetKey(v string) *ListProvisionedProductPlanApproversRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequestFilters) SetValue(v string) *ListProvisionedProductPlanApproversRequestFilters {
	s.Value = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequestFilters) Validate() error {
	return dara.Validate(s)
}
