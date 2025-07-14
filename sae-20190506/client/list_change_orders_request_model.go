// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListChangeOrdersRequest
	GetAppId() *string
	SetCoStatus(v string) *ListChangeOrdersRequest
	GetCoStatus() *string
	SetCoType(v string) *ListChangeOrdersRequest
	GetCoType() *string
	SetCurrentPage(v int32) *ListChangeOrdersRequest
	GetCurrentPage() *int32
	SetKey(v string) *ListChangeOrdersRequest
	GetKey() *string
	SetOrderBy(v string) *ListChangeOrdersRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListChangeOrdersRequest
	GetPageSize() *int32
	SetReverse(v bool) *ListChangeOrdersRequest
	GetReverse() *bool
}

type ListChangeOrdersRequest struct {
	// 1
	//
	// This parameter is required.
	//
	// example:
	//
	// 145341c-9708-4967-b3ec-24933767****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2
	CoStatus *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	// The type of the change order. Valid values:
	//
	// 	- **CoBindSlb**: associates the Server Load Balancer (SLB) instance with the application.
	//
	// 	- **CoUnbindSlb**: disassociates an SLB instance from the application.
	//
	// 	- **CoCreateApp**: creates the application.
	//
	// 	- **CoDeleteApp**: deletes the application.
	//
	// 	- **CoDeploy**: deploys the application.
	//
	// 	- **CoRestartApplication**: restarts the application.
	//
	// 	- **CoRollback**: rolls back the application.
	//
	// 	- **CoScaleIn**: scales in the application.
	//
	// 	- **CoScaleOut**: scales out the application.
	//
	// 	- **CoStartApplication**: starts the application.
	//
	// 	- **CoStopApplication**: stops the application.
	//
	// 	- **CoRescaleApplicationVertically**: modifies the instance type.
	//
	// 	- **CoDeployHistroy**: rolls back the application to an earlier version.
	//
	// 	- **CoBindNas**: associates a network-attached storage (NAS) file system with the application.
	//
	// 	- **CoUnbindNas**: disassociates a NAS file system from the application.
	//
	// 	- **CoBatchStartApplication**: starts multiple applications concurrently.
	//
	// 	- **CoBatchStopApplication**: stops multiple applications concurrently.
	//
	// 	- **CoRestartInstances**: restarts the instance.
	//
	// 	- **CoDeleteInstances**: deletes the instance.
	//
	// 	- **CoScaleInAppWithInstances**: reduces the specified number of application instances.
	//
	// example:
	//
	// CoCreateApp
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// 20
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// CoCreateApp
	//
	// example:
	//
	// test
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// test
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse  *bool  `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s ListChangeOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChangeOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListChangeOrdersRequest) GetCoStatus() *string {
	return s.CoStatus
}

func (s *ListChangeOrdersRequest) GetCoType() *string {
	return s.CoType
}

func (s *ListChangeOrdersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListChangeOrdersRequest) GetKey() *string {
	return s.Key
}

func (s *ListChangeOrdersRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListChangeOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChangeOrdersRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListChangeOrdersRequest) SetAppId(v string) *ListChangeOrdersRequest {
	s.AppId = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCoStatus(v string) *ListChangeOrdersRequest {
	s.CoStatus = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCoType(v string) *ListChangeOrdersRequest {
	s.CoType = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCurrentPage(v int32) *ListChangeOrdersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListChangeOrdersRequest) SetKey(v string) *ListChangeOrdersRequest {
	s.Key = &v
	return s
}

func (s *ListChangeOrdersRequest) SetOrderBy(v string) *ListChangeOrdersRequest {
	s.OrderBy = &v
	return s
}

func (s *ListChangeOrdersRequest) SetPageSize(v int32) *ListChangeOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListChangeOrdersRequest) SetReverse(v bool) *ListChangeOrdersRequest {
	s.Reverse = &v
	return s
}

func (s *ListChangeOrdersRequest) Validate() error {
	return dara.Validate(s)
}
