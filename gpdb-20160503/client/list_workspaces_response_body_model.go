// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListWorkspacesResponseBodyItems) *ListWorkspacesResponseBody
	GetItems() []*ListWorkspacesResponseBodyItems
	SetMaxResults(v int32) *ListWorkspacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListWorkspacesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListWorkspacesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListWorkspacesResponseBody
	GetTotalRecordCount() *int32
}

type ListWorkspacesResponseBody struct {
	// The list of workspace details.
	Items []*ListWorkspacesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Default value: 10.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetItems() []*ListWorkspacesResponseBodyItems {
	return s.Items
}

func (s *ListWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkspacesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListWorkspacesResponseBody) SetItems(v []*ListWorkspacesResponseBodyItems) *ListWorkspacesResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetPageNumber(v int32) *ListWorkspacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetPageRecordCount(v int32) *ListWorkspacesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalRecordCount(v int32) *ListWorkspacesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
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

type ListWorkspacesResponseBodyItems struct {
	// The list of API keys for the workspace.
	Apikeys []*ListWorkspacesResponseBodyItemsApikeys `json:"Apikeys,omitempty" xml:"Apikeys,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The services in the workspace.
	Services []*ListWorkspacesResponseBodyItemsServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// gp-ws-wkb4fp3j9u79ha
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// - Maximum length: 50.
	//
	// - Must be unique.
	//
	// - Special characters are not allowed.
	//
	// example:
	//
	// anchashid8FocugQ.oxs.xaliyun.com/oxspopscand8FocugQ#
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItems) GetApikeys() []*ListWorkspacesResponseBodyItemsApikeys {
	return s.Apikeys
}

func (s *ListWorkspacesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyItems) GetServices() []*ListWorkspacesResponseBodyItemsServices {
	return s.Services
}

func (s *ListWorkspacesResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyItems) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyItems) SetApikeys(v []*ListWorkspacesResponseBodyItemsApikeys) *ListWorkspacesResponseBodyItems {
	s.Apikeys = v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetCreateTime(v string) *ListWorkspacesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetServices(v []*ListWorkspacesResponseBodyItemsServices) *ListWorkspacesResponseBodyItems {
	s.Services = v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetWorkspaceId(v string) *ListWorkspacesResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetWorkspaceName(v string) *ListWorkspacesResponseBodyItems {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) Validate() error {
	if s.Apikeys != nil {
		for _, item := range s.Apikeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyItemsApikeys struct {
	// The services authorized for the API key.
	AuthServices []*ListWorkspacesResponseBodyItemsApikeysAuthServices `json:"AuthServices,omitempty" xml:"AuthServices,omitempty" type:"Repeated"`
	// The creation time of the API key.
	//
	// example:
	//
	// 2026-03-09T02:26:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// my api key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API key.
	//
	// example:
	//
	// api-xxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// my-apikey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// sk-xxxxxx
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
}

func (s ListWorkspacesResponseBodyItemsApikeys) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItemsApikeys) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetAuthServices() []*ListWorkspacesResponseBodyItemsApikeysAuthServices {
	return s.AuthServices
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetKeyId() *string {
	return s.KeyId
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetKeyName() *string {
	return s.KeyName
}

func (s *ListWorkspacesResponseBodyItemsApikeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetAuthServices(v []*ListWorkspacesResponseBodyItemsApikeysAuthServices) *ListWorkspacesResponseBodyItemsApikeys {
	s.AuthServices = v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetCreateTime(v string) *ListWorkspacesResponseBodyItemsApikeys {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetDescription(v string) *ListWorkspacesResponseBodyItemsApikeys {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetKeyId(v string) *ListWorkspacesResponseBodyItemsApikeys {
	s.KeyId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetKeyName(v string) *ListWorkspacesResponseBodyItemsApikeys {
	s.KeyName = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) SetKeyPrefix(v string) *ListWorkspacesResponseBodyItemsApikeys {
	s.KeyPrefix = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeys) Validate() error {
	if s.AuthServices != nil {
		for _, item := range s.AuthServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyItemsApikeysAuthServices struct {
	// The service ID.
	//
	// example:
	//
	// agdb-eqxwj5tj5ojx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type. Valid values:
	//
	// - memory
	//
	// - drama.
	//
	// example:
	//
	// drama
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListWorkspacesResponseBodyItemsApikeysAuthServices) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItemsApikeysAuthServices) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItemsApikeysAuthServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListWorkspacesResponseBodyItemsApikeysAuthServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListWorkspacesResponseBodyItemsApikeysAuthServices) SetServiceId(v string) *ListWorkspacesResponseBodyItemsApikeysAuthServices {
	s.ServiceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeysAuthServices) SetServiceType(v string) *ListWorkspacesResponseBodyItemsApikeysAuthServices {
	s.ServiceType = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsApikeysAuthServices) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesResponseBodyItemsServices struct {
	// The creation time.
	//
	// example:
	//
	// 2026-03-09T02:26:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The compute resource.
	//
	// example:
	//
	// 1
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-01-26T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing type. Valid values:
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// - **PREPAY**: subscription.
	//
	// > - If this parameter is not specified, the default value is pay-as-you-go.
	//
	// > - In subscription billing mode, a discount is available when you purchase a duration of one year or longer. Select the billing type as needed.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// [Deprecated].
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
	// - memory
	//
	// - drama.
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service status. Valid values:
	//
	// - creating: The service is being created.
	//
	// - active: The service is running.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkspacesResponseBodyItemsServices) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItemsServices) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItemsServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyItemsServices) GetCu() *string {
	return s.Cu
}

func (s *ListWorkspacesResponseBodyItemsServices) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListWorkspacesResponseBodyItemsServices) GetPayType() *string {
	return s.PayType
}

func (s *ListWorkspacesResponseBodyItemsServices) GetPlan() *string {
	return s.Plan
}

func (s *ListWorkspacesResponseBodyItemsServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListWorkspacesResponseBodyItemsServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListWorkspacesResponseBodyItemsServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListWorkspacesResponseBodyItemsServices) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspacesResponseBodyItemsServices) SetCreateTime(v string) *ListWorkspacesResponseBodyItemsServices {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetCu(v string) *ListWorkspacesResponseBodyItemsServices {
	s.Cu = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetExpireTime(v string) *ListWorkspacesResponseBodyItemsServices {
	s.ExpireTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetPayType(v string) *ListWorkspacesResponseBodyItemsServices {
	s.PayType = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetPlan(v string) *ListWorkspacesResponseBodyItemsServices {
	s.Plan = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetServiceId(v string) *ListWorkspacesResponseBodyItemsServices {
	s.ServiceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetServiceName(v string) *ListWorkspacesResponseBodyItemsServices {
	s.ServiceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetServiceType(v string) *ListWorkspacesResponseBodyItemsServices {
	s.ServiceType = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) SetStatus(v string) *ListWorkspacesResponseBodyItemsServices {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsServices) Validate() error {
	return dara.Validate(s)
}
