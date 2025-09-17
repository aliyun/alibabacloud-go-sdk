// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceTypesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceTypesResponseBody
	GetRequestId() *string
	SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody
	GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes
	SetTotalCount(v int32) *ListResourceTypesResponseBody
	GetTotalCount() *int32
}

type ListResourceTypesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// ECS::Disk
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of a request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the resource types.
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypesResponseBody) GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes {
	return s.ResourceTypes
}

func (s *ListResourceTypesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceTypesResponseBody) SetMaxResults(v int32) *ListResourceTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetNextToken(v string) *ListResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) SetTotalCount(v int32) *ListResourceTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypes struct {
	// The properties that are specific to the create operation. You need to specify these properties when you create the resource. These properties are not returned when you query the resource.
	CreateOnlyProperties []*string `json:"createOnlyProperties,omitempty" xml:"createOnlyProperties,omitempty" type:"Repeated"`
	// The properties that are specific to the delete operation. You need to specify these properties when you delete the resource. These properties are not returned when you query the resource.
	DeleteOnlyProperties []*string `json:"deleteOnlyProperties,omitempty" xml:"deleteOnlyProperties,omitempty" type:"Repeated"`
	// The properties that can be used to filter the resource when you list the resource.
	FilterProperties []*string `json:"filterProperties,omitempty" xml:"filterProperties,omitempty" type:"Repeated"`
	// The properties that are specific to the query operation. You need to specify these properties when you query the resource. These properties are not returned in the query result.
	GetOnlyProperties []*string `json:"getOnlyProperties,omitempty" xml:"getOnlyProperties,omitempty" type:"Repeated"`
	// The properties that are returned when you query the resource.
	GetResponseProperties []*string `json:"getResponseProperties,omitempty" xml:"getResponseProperties,omitempty" type:"Repeated"`
	// The information about the operation, including the required Resource Access Management (RAM) permissions.
	Handlers *ListResourceTypesResponseBodyResourceTypesHandlers `json:"handlers,omitempty" xml:"handlers,omitempty" type:"Struct"`
	// The information about the resource type.
	Info *ListResourceTypesResponseBodyResourceTypesInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	// The properties that are specific to the list operation. You need to specify these properties when you list the resource. These properties are not returned when you query the resource.
	ListOnlyProperties []*string `json:"listOnlyProperties,omitempty" xml:"listOnlyProperties,omitempty" type:"Repeated"`
	// The properties that are returned when you list the resource.
	ListResponseProperties []*string `json:"listResponseProperties,omitempty" xml:"listResponseProperties,omitempty" type:"Repeated"`
	// The ID of the resource.
	//
	// example:
	//
	// /properties/InstanceId
	PrimaryIdentifier *string `json:"primaryIdentifier,omitempty" xml:"primaryIdentifier,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The resource properties. The key specifies the property name and the value specifies the details of the property.
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The common properties of the resource. The common properties are not operation-specific.
	PublicProperties []*string `json:"publicProperties,omitempty" xml:"publicProperties,omitempty" type:"Repeated"`
	// The read-only properties. These properties are returned only when you list or query the resource. You do not need to specify these properties when you create or update the resource.
	ReadOnlyProperties []*string `json:"readOnlyProperties,omitempty" xml:"readOnlyProperties,omitempty" type:"Repeated"`
	// The properties that must be specified when you create the resource.
	Required []*string `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The sensitive properties, such as the password.
	SensitiveInfoProperties []*string `json:"sensitiveInfoProperties,omitempty" xml:"sensitiveInfoProperties,omitempty" type:"Repeated"`
	// The properties that are specific to the update operation. You need to specify these properties when you update the resource. These properties are not returned when you query the resource.
	UpdateOnlyProperties []*string `json:"updateOnlyProperties,omitempty" xml:"updateOnlyProperties,omitempty" type:"Repeated"`
	// The properties that can be modified.
	UpdateTypeProperties []*string `json:"updateTypeProperties,omitempty" xml:"updateTypeProperties,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetCreateOnlyProperties() []*string {
	return s.CreateOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetDeleteOnlyProperties() []*string {
	return s.DeleteOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetFilterProperties() []*string {
	return s.FilterProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetGetOnlyProperties() []*string {
	return s.GetOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetGetResponseProperties() []*string {
	return s.GetResponseProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetHandlers() *ListResourceTypesResponseBodyResourceTypesHandlers {
	return s.Handlers
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetInfo() *ListResourceTypesResponseBodyResourceTypesInfo {
	return s.Info
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetListOnlyProperties() []*string {
	return s.ListOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetListResponseProperties() []*string {
	return s.ListResponseProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetPrimaryIdentifier() *string {
	return s.PrimaryIdentifier
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetProduct() *string {
	return s.Product
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetPublicProperties() []*string {
	return s.PublicProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetReadOnlyProperties() []*string {
	return s.ReadOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetRequired() []*string {
	return s.Required
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetSensitiveInfoProperties() []*string {
	return s.SensitiveInfoProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetUpdateOnlyProperties() []*string {
	return s.UpdateOnlyProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetUpdateTypeProperties() []*string {
	return s.UpdateTypeProperties
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCreateOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.CreateOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetDeleteOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.DeleteOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetGetOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.GetOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetGetResponseProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.GetResponseProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetHandlers(v *ListResourceTypesResponseBodyResourceTypesHandlers) *ListResourceTypesResponseBodyResourceTypes {
	s.Handlers = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetInfo(v *ListResourceTypesResponseBodyResourceTypesInfo) *ListResourceTypesResponseBodyResourceTypes {
	s.Info = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetListOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ListOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetListResponseProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ListResponseProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetPrimaryIdentifier(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.PrimaryIdentifier = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProduct(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Product = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProperties(v map[string]interface{}) *ListResourceTypesResponseBodyResourceTypes {
	s.Properties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetPublicProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.PublicProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetReadOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ReadOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetRequired(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.Required = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetSensitiveInfoProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.SensitiveInfoProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetUpdateOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.UpdateOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetUpdateTypeProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.UpdateTypeProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlers struct {
	// The information about the create operation.
	Create *ListResourceTypesResponseBodyResourceTypesHandlersCreate `json:"create,omitempty" xml:"create,omitempty" type:"Struct"`
	// The information about the delete operation.
	Delete *ListResourceTypesResponseBodyResourceTypesHandlersDelete `json:"delete,omitempty" xml:"delete,omitempty" type:"Struct"`
	// The information about the query operation.
	Get *ListResourceTypesResponseBodyResourceTypesHandlersGet `json:"get,omitempty" xml:"get,omitempty" type:"Struct"`
	// The information about the list operation.
	List *ListResourceTypesResponseBodyResourceTypesHandlersList `json:"list,omitempty" xml:"list,omitempty" type:"Struct"`
	// The information about the update operation.
	Update *ListResourceTypesResponseBodyResourceTypesHandlersUpdate `json:"update,omitempty" xml:"update,omitempty" type:"Struct"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlers) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlers) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) GetCreate() *ListResourceTypesResponseBodyResourceTypesHandlersCreate {
	return s.Create
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) GetDelete() *ListResourceTypesResponseBodyResourceTypesHandlersDelete {
	return s.Delete
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) GetGet() *ListResourceTypesResponseBodyResourceTypesHandlersGet {
	return s.Get
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) GetList() *ListResourceTypesResponseBodyResourceTypesHandlersList {
	return s.List
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) GetUpdate() *ListResourceTypesResponseBodyResourceTypesHandlersUpdate {
	return s.Update
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetCreate(v *ListResourceTypesResponseBodyResourceTypesHandlersCreate) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Create = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetDelete(v *ListResourceTypesResponseBodyResourceTypesHandlersDelete) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Delete = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetGet(v *ListResourceTypesResponseBodyResourceTypesHandlersGet) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Get = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetList(v *ListResourceTypesResponseBodyResourceTypesHandlersList) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.List = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetUpdate(v *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Update = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlersCreate struct {
	// The RAM permissions required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersCreate) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersCreate) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersCreate) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersCreate) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersCreate {
	s.Permissions = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersCreate) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlersDelete struct {
	// The RAM permissions required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersDelete) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersDelete) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersDelete) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersDelete) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersDelete {
	s.Permissions = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersDelete) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlersGet struct {
	// The RAM permissions required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersGet) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersGet) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersGet) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersGet) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersGet {
	s.Permissions = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersGet) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlersList struct {
	// The RAM permissions required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersList) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersList) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersList) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersList {
	s.Permissions = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersList) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesHandlersUpdate struct {
	// The RAM permissions required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersUpdate) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersUpdate) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersUpdate {
	s.Permissions = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypesInfo struct {
	// Billing method\\
	//
	// paid free
	//
	// example:
	//
	// paid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The deployment level of the resource.
	//
	// center
	//
	// region
	//
	// zone
	//
	// example:
	//
	// region
	DeliveryScope *string `json:"deliveryScope,omitempty" xml:"deliveryScope,omitempty"`
	// The description of the resource type.
	//
	// example:
	//
	// An ECS instance is equivalent to a virtual machine, including the most basic computing components such as CPU, memory, operating system, network, and disk. You can easily customize and change the configuration of the instance. You have full control over the virtual machine.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// Instance
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) GetDeliveryScope() *string {
	return s.DeliveryScope
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) GetDescription() *string {
	return s.Description
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) GetTitle() *string {
	return s.Title
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetChargeType(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.ChargeType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDeliveryScope(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.DeliveryScope = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetTitle(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Title = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) Validate() error {
	return dara.Validate(s)
}
