// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppServicesPageRequest
	GetAppId() *string
	SetPageNumber(v int32) *ListAppServicesPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppServicesPageRequest
	GetPageSize() *int32
	SetServiceType(v string) *ListAppServicesPageRequest
	GetServiceType() *string
}

type ListAppServicesPageRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6dcc8c9e-d3da-478a-a066-86dcf820****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Valid values: 0 to 9999.
	//
	// example:
	//
	// 9999
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service type. Valid values:
	//
	// 	- **dubbo**
	//
	// 	- **springCloud**
	//
	// This parameter is required.
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListAppServicesPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesPageRequest) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppServicesPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppServicesPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppServicesPageRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListAppServicesPageRequest) SetAppId(v string) *ListAppServicesPageRequest {
	s.AppId = &v
	return s
}

func (s *ListAppServicesPageRequest) SetPageNumber(v int32) *ListAppServicesPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppServicesPageRequest) SetPageSize(v int32) *ListAppServicesPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppServicesPageRequest) SetServiceType(v string) *ListAppServicesPageRequest {
	s.ServiceType = &v
	return s
}

func (s *ListAppServicesPageRequest) Validate() error {
	return dara.Validate(s)
}
