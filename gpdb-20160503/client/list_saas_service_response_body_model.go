// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListSaasServiceResponseBodyItems) *ListSaasServiceResponseBody
	GetItems() []*ListSaasServiceResponseBodyItems
	SetMaxResults(v int32) *ListSaasServiceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSaasServiceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSaasServiceResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListSaasServiceResponseBody
	GetTotalRecordCount() *int32
}

type ListSaasServiceResponseBody struct {
	// The list of instance details.
	Items []*ListSaasServiceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Default value: 10.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaasServiceResponseBody) GetItems() []*ListSaasServiceResponseBodyItems {
	return s.Items
}

func (s *ListSaasServiceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSaasServiceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSaasServiceResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListSaasServiceResponseBody) SetItems(v []*ListSaasServiceResponseBodyItems) *ListSaasServiceResponseBody {
	s.Items = v
	return s
}

func (s *ListSaasServiceResponseBody) SetMaxResults(v int32) *ListSaasServiceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSaasServiceResponseBody) SetNextToken(v string) *ListSaasServiceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSaasServiceResponseBody) SetRequestId(v string) *ListSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSaasServiceResponseBody) SetTotalRecordCount(v int32) *ListSaasServiceResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListSaasServiceResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSaasServiceResponseBodyItems struct {
	// The list of service subcomponents.
	Components []*ListSaasServiceResponseBodyItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The compute resource.
	//
	// example:
	//
	// 1
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// Indicates whether the release protection feature is enabled. Valid values:
	//
	// 	- **true**: Enabled.
	//
	// 	- **false**: Disabled.
	//
	// example:
	//
	// True
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2021-10-15T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing type. Valid values:
	//
	// - **POSTPAY**: Pay-as-you-go.
	//
	// - **PREPAY**: Subscription.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// deprecated
	Plan *string `json:"Plan,omitempty" xml:"Plan,omitempty"`
	// The service ID.
	//
	// example:
	//
	// agdb-xxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// agdb-xxxx
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service type. Valid values:
	//
	// - **memory**
	//
	// - **drama**
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service status. Valid values:
	//
	// - active: Running.
	//
	// - creating: Being created.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSaasServiceResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSaasServiceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSaasServiceResponseBodyItems) GetComponents() []*ListSaasServiceResponseBodyItemsComponents {
	return s.Components
}

func (s *ListSaasServiceResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSaasServiceResponseBodyItems) GetCu() *int32 {
	return s.Cu
}

func (s *ListSaasServiceResponseBodyItems) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListSaasServiceResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListSaasServiceResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *ListSaasServiceResponseBodyItems) GetPlan() *string {
	return s.Plan
}

func (s *ListSaasServiceResponseBodyItems) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListSaasServiceResponseBodyItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListSaasServiceResponseBodyItems) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListSaasServiceResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListSaasServiceResponseBodyItems) SetComponents(v []*ListSaasServiceResponseBodyItemsComponents) *ListSaasServiceResponseBodyItems {
	s.Components = v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetCreateTime(v string) *ListSaasServiceResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetCu(v int32) *ListSaasServiceResponseBodyItems {
	s.Cu = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetDeletionProtection(v bool) *ListSaasServiceResponseBodyItems {
	s.DeletionProtection = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetExpireTime(v string) *ListSaasServiceResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetPayType(v string) *ListSaasServiceResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetPlan(v string) *ListSaasServiceResponseBodyItems {
	s.Plan = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetServiceId(v string) *ListSaasServiceResponseBodyItems {
	s.ServiceId = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetServiceName(v string) *ListSaasServiceResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetServiceType(v string) *ListSaasServiceResponseBodyItems {
	s.ServiceType = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) SetStatus(v string) *ListSaasServiceResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListSaasServiceResponseBodyItems) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSaasServiceResponseBodyItemsComponents struct {
	// The component ID.
	//
	// example:
	//
	// 0644c5aa-5306-478b-ac39-bb4660cdc9f7
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The type of the subcomponent.
	//
	// example:
	//
	// gamestudio
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The compute resource of the component.
	//
	// example:
	//
	// 2
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The release protection status.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The service status. Valid values:
	//
	// - active: Running.
	//
	// - creating: Being created.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSaasServiceResponseBodyItemsComponents) String() string {
	return dara.Prettify(s)
}

func (s ListSaasServiceResponseBodyItemsComponents) GoString() string {
	return s.String()
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetCu() *string {
	return s.Cu
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListSaasServiceResponseBodyItemsComponents) GetStatus() *string {
	return s.Status
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetComponentId(v string) *ListSaasServiceResponseBodyItemsComponents {
	s.ComponentId = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetComponentType(v string) *ListSaasServiceResponseBodyItemsComponents {
	s.ComponentType = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetCreateTime(v string) *ListSaasServiceResponseBodyItemsComponents {
	s.CreateTime = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetCu(v string) *ListSaasServiceResponseBodyItemsComponents {
	s.Cu = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetDeletionProtection(v bool) *ListSaasServiceResponseBodyItemsComponents {
	s.DeletionProtection = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) SetStatus(v string) *ListSaasServiceResponseBodyItemsComponents {
	s.Status = &v
	return s
}

func (s *ListSaasServiceResponseBodyItemsComponents) Validate() error {
	return dara.Validate(s)
}
