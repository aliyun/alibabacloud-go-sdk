// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourceUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppResourceUserPermissionsResponseBody
	GetCode() *string
	SetItems(v []*ListAppResourceUserPermissionsResponseBodyItems) *ListAppResourceUserPermissionsResponseBody
	GetItems() []*ListAppResourceUserPermissionsResponseBodyItems
	SetMessage(v string) *ListAppResourceUserPermissionsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListAppResourceUserPermissionsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourceUserPermissionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAppResourceUserPermissionsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAppResourceUserPermissionsResponseBody
	GetTotal() *int64
}

type ListAppResourceUserPermissionsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The resources and permissions on the current page.
	Items []*ListAppResourceUserPermissionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error description. This parameter is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAppResourceUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourceUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResourceUserPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppResourceUserPermissionsResponseBody) GetItems() []*ListAppResourceUserPermissionsResponseBodyItems {
	return s.Items
}

func (s *ListAppResourceUserPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppResourceUserPermissionsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourceUserPermissionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourceUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppResourceUserPermissionsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAppResourceUserPermissionsResponseBody) SetCode(v string) *ListAppResourceUserPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetItems(v []*ListAppResourceUserPermissionsResponseBodyItems) *ListAppResourceUserPermissionsResponseBody {
	s.Items = v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetMessage(v string) *ListAppResourceUserPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetPage(v int64) *ListAppResourceUserPermissionsResponseBody {
	s.Page = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetPageSize(v int64) *ListAppResourceUserPermissionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetRequestId(v string) *ListAppResourceUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) SetTotal(v int64) *ListAppResourceUserPermissionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBody) Validate() error {
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

type ListAppResourceUserPermissionsResponseBodyItems struct {
	// The stored and valid permission points.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// The resource key.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListAppResourceUserPermissionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourceUserPermissionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) SetPermissions(v []*string) *ListAppResourceUserPermissionsResponseBodyItems {
	s.Permissions = v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) SetResourceKey(v string) *ListAppResourceUserPermissionsResponseBodyItems {
	s.ResourceKey = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) SetResourceType(v string) *ListAppResourceUserPermissionsResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
