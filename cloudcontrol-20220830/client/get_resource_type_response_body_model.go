// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceTypeResponseBody
	GetRequestId() *string
	SetResourceType(v *GetResourceTypeResponseBodyResourceType) *GetResourceTypeResponseBody
	GetResourceType() *GetResourceTypeResponseBodyResourceType
}

type GetResourceTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// No parent resource:
	//
	// Instance
	//
	// Has parent resource:
	//
	// DBInstance/Account
	ResourceType *GetResourceTypeResponseBodyResourceType `json:"resourceType,omitempty" xml:"resourceType,omitempty" type:"Struct"`
}

func (s GetResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceTypeResponseBody) GetResourceType() *GetResourceTypeResponseBodyResourceType {
	return s.ResourceType
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v *GetResourceTypeResponseBodyResourceType) *GetResourceTypeResponseBody {
	s.ResourceType = v
	return s
}

func (s *GetResourceTypeResponseBody) Validate() error {
	if s.ResourceType != nil {
		if err := s.ResourceType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceTypeResponseBodyResourceType struct {
	// The properties exclusive to the create operation. These properties are not returned in resource query operations but are required as input parameters for the create operation.
	CreateOnlyProperties []*string `json:"createOnlyProperties,omitempty" xml:"createOnlyProperties,omitempty" type:"Repeated"`
	// The properties exclusive to the delete operation. These properties are not returned in resource query operations but are required as input parameters for the delete operation.
	DeleteOnlyProperties []*string `json:"deleteOnlyProperties,omitempty" xml:"deleteOnlyProperties,omitempty" type:"Repeated"`
	// The properties that can be used as filter parameters in the list operation.
	FilterProperties []*string `json:"filterProperties,omitempty" xml:"filterProperties,omitempty" type:"Repeated"`
	// The properties exclusive to the get operation. These properties are not returned in resource query operations but are required as input parameters for the get operation.
	GetOnlyProperties []*string `json:"getOnlyProperties,omitempty" xml:"getOnlyProperties,omitempty" type:"Repeated"`
	// The properties returned by the get operation.
	GetResponseProperties []*string `json:"getResponseProperties,omitempty" xml:"getResponseProperties,omitempty" type:"Repeated"`
	// The supported resource operations, including RAM permissions.
	Handlers *GetResourceTypeResponseBodyResourceTypeHandlers `json:"handlers,omitempty" xml:"handlers,omitempty" type:"Struct"`
	// The basic information about the resource type.
	Info *GetResourceTypeResponseBodyResourceTypeInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	// The properties exclusive to the list operation. These properties are not returned in resource query operations but are required as input parameters for the list operation.
	ListOnlyProperties []*string `json:"listOnlyProperties,omitempty" xml:"listOnlyProperties,omitempty" type:"Repeated"`
	// The properties returned by the list operation.
	ListResponseProperties []*string `json:"listResponseProperties,omitempty" xml:"listResponseProperties,omitempty" type:"Repeated"`
	// The resource ID.
	//
	// example:
	//
	// /properties/InstanceId
	PrimaryIdentifier *string `json:"primaryIdentifier,omitempty" xml:"primaryIdentifier,omitempty"`
	// The product code.
	//
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The resource property definitions. The key is the property name, and the value is the detailed property information.
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The common properties that represent basic resource attributes. These are not operation-specific properties.
	PublicProperties []*string `json:"publicProperties,omitempty" xml:"publicProperties,omitempty" type:"Repeated"`
	// The read-only properties. These properties are returned only in list or get operations and cannot be used as input parameters for create or update operations.
	ReadOnlyProperties []*string `json:"readOnlyProperties,omitempty" xml:"readOnlyProperties,omitempty" type:"Repeated"`
	// The required parameters for resource creation.
	Required []*string `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
	// The resource type. If the resource has a parent resource, the format is {parentResourceTypeCode/resourceTypeCode}.
	//
	// example:
	//
	// 无父资源：
	//
	// Instance
	//
	// 有父资源：
	//
	// DBInstance/Account
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The sensitive properties, such as passwords.
	SensitiveInfoProperties []*string `json:"sensitiveInfoProperties,omitempty" xml:"sensitiveInfoProperties,omitempty" type:"Repeated"`
	// The properties exclusive to the update operation. These properties are not returned in resource query operations but are required as input parameters for the update operation.
	UpdateOnlyProperties []*string `json:"updateOnlyProperties,omitempty" xml:"updateOnlyProperties,omitempty" type:"Repeated"`
	// The properties that can be modified.
	UpdateTypeProperties []*string `json:"updateTypeProperties,omitempty" xml:"updateTypeProperties,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceType) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceType) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceType) GetCreateOnlyProperties() []*string {
	return s.CreateOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetDeleteOnlyProperties() []*string {
	return s.DeleteOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetFilterProperties() []*string {
	return s.FilterProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetGetOnlyProperties() []*string {
	return s.GetOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetGetResponseProperties() []*string {
	return s.GetResponseProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetHandlers() *GetResourceTypeResponseBodyResourceTypeHandlers {
	return s.Handlers
}

func (s *GetResourceTypeResponseBodyResourceType) GetInfo() *GetResourceTypeResponseBodyResourceTypeInfo {
	return s.Info
}

func (s *GetResourceTypeResponseBodyResourceType) GetListOnlyProperties() []*string {
	return s.ListOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetListResponseProperties() []*string {
	return s.ListResponseProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetPrimaryIdentifier() *string {
	return s.PrimaryIdentifier
}

func (s *GetResourceTypeResponseBodyResourceType) GetProduct() *string {
	return s.Product
}

func (s *GetResourceTypeResponseBodyResourceType) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetResourceTypeResponseBodyResourceType) GetPublicProperties() []*string {
	return s.PublicProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetReadOnlyProperties() []*string {
	return s.ReadOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetRequired() []*string {
	return s.Required
}

func (s *GetResourceTypeResponseBodyResourceType) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceTypeResponseBodyResourceType) GetSensitiveInfoProperties() []*string {
	return s.SensitiveInfoProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetUpdateOnlyProperties() []*string {
	return s.UpdateOnlyProperties
}

func (s *GetResourceTypeResponseBodyResourceType) GetUpdateTypeProperties() []*string {
	return s.UpdateTypeProperties
}

func (s *GetResourceTypeResponseBodyResourceType) SetCreateOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.CreateOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetDeleteOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.DeleteOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetFilterProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.FilterProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetGetOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.GetOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetGetResponseProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.GetResponseProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetHandlers(v *GetResourceTypeResponseBodyResourceTypeHandlers) *GetResourceTypeResponseBodyResourceType {
	s.Handlers = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetInfo(v *GetResourceTypeResponseBodyResourceTypeInfo) *GetResourceTypeResponseBodyResourceType {
	s.Info = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetListOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ListOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetListResponseProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ListResponseProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetPrimaryIdentifier(v string) *GetResourceTypeResponseBodyResourceType {
	s.PrimaryIdentifier = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProduct(v string) *GetResourceTypeResponseBodyResourceType {
	s.Product = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBodyResourceType {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetPublicProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.PublicProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetReadOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ReadOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetRequired(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.Required = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetResourceType(v string) *GetResourceTypeResponseBodyResourceType {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetSensitiveInfoProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.SensitiveInfoProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetUpdateOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.UpdateOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetUpdateTypeProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.UpdateTypeProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) Validate() error {
	if s.Handlers != nil {
		if err := s.Handlers.Validate(); err != nil {
			return err
		}
	}
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceTypeResponseBodyResourceTypeHandlers struct {
	// The information associated with the create operation.
	Create *GetResourceTypeResponseBodyResourceTypeHandlersCreate `json:"create,omitempty" xml:"create,omitempty" type:"Struct"`
	// The information associated with the delete operation.
	Delete *GetResourceTypeResponseBodyResourceTypeHandlersDelete `json:"delete,omitempty" xml:"delete,omitempty" type:"Struct"`
	// The information associated with the get operation.
	Get *GetResourceTypeResponseBodyResourceTypeHandlersGet `json:"get,omitempty" xml:"get,omitempty" type:"Struct"`
	// The information associated with the list operation.
	List *GetResourceTypeResponseBodyResourceTypeHandlersList `json:"list,omitempty" xml:"list,omitempty" type:"Struct"`
	// The information associated with the update operation.
	Update *GetResourceTypeResponseBodyResourceTypeHandlersUpdate `json:"update,omitempty" xml:"update,omitempty" type:"Struct"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlers) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlers) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) GetCreate() *GetResourceTypeResponseBodyResourceTypeHandlersCreate {
	return s.Create
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) GetDelete() *GetResourceTypeResponseBodyResourceTypeHandlersDelete {
	return s.Delete
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) GetGet() *GetResourceTypeResponseBodyResourceTypeHandlersGet {
	return s.Get
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) GetList() *GetResourceTypeResponseBodyResourceTypeHandlersList {
	return s.List
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) GetUpdate() *GetResourceTypeResponseBodyResourceTypeHandlersUpdate {
	return s.Update
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetCreate(v *GetResourceTypeResponseBodyResourceTypeHandlersCreate) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Create = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetDelete(v *GetResourceTypeResponseBodyResourceTypeHandlersDelete) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Delete = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetGet(v *GetResourceTypeResponseBodyResourceTypeHandlersGet) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Get = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetList(v *GetResourceTypeResponseBodyResourceTypeHandlersList) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.List = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetUpdate(v *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Update = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) Validate() error {
	if s.Create != nil {
		if err := s.Create.Validate(); err != nil {
			return err
		}
	}
	if s.Delete != nil {
		if err := s.Delete.Validate(); err != nil {
			return err
		}
	}
	if s.Get != nil {
		if err := s.Get.Validate(); err != nil {
			return err
		}
	}
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	if s.Update != nil {
		if err := s.Update.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceTypeResponseBodyResourceTypeHandlersCreate struct {
	// The required RAM permissions.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersCreate) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersCreate) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersCreate) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersCreate) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersCreate {
	s.Permissions = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersCreate) Validate() error {
	return dara.Validate(s)
}

type GetResourceTypeResponseBodyResourceTypeHandlersDelete struct {
	// The required RAM permissions.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersDelete) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersDelete) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersDelete) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersDelete) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersDelete {
	s.Permissions = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersDelete) Validate() error {
	return dara.Validate(s)
}

type GetResourceTypeResponseBodyResourceTypeHandlersGet struct {
	// The required RAM permissions.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersGet) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersGet) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersGet) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersGet) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersGet {
	s.Permissions = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersGet) Validate() error {
	return dara.Validate(s)
}

type GetResourceTypeResponseBodyResourceTypeHandlersList struct {
	// The required RAM permissions.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersList) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersList) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersList) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersList {
	s.Permissions = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersList) Validate() error {
	return dara.Validate(s)
}

type GetResourceTypeResponseBodyResourceTypeHandlersUpdate struct {
	// The required RAM permissions.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersUpdate) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersUpdate) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersUpdate {
	s.Permissions = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) Validate() error {
	return dara.Validate(s)
}

type GetResourceTypeResponseBodyResourceTypeInfo struct {
	// The billing method. Valid values:
	//
	// paid: paid.
	//
	// free: free.
	//
	// example:
	//
	// paid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The delivery scope. Valid values:
	//
	// center: centralized deployment.
	//
	// region: region-level deployment.
	//
	// zone: zone-level deployment.
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

func (s GetResourceTypeResponseBodyResourceTypeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeInfo) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) GetDeliveryScope() *string {
	return s.DeliveryScope
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) GetTitle() *string {
	return s.Title
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetChargeType(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.ChargeType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetDeliveryScope(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.DeliveryScope = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetDescription(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.Description = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetTitle(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.Title = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) Validate() error {
	return dara.Validate(s)
}
