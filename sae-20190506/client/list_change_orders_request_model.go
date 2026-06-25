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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145341c-9708-4967-b3ec-24933767****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The status of the change order. Valid values:
	//
	// - **0**: Preparing.
	//
	// - **1**: In progress.
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// - **6**: Stopped.
	//
	// - **8**: Paused for manual confirmation.
	//
	// - **9**: Paused for automatic confirmation.
	//
	// - **10**: Failed due to a system exception.
	//
	// - **11**: Pending approval.
	//
	// - **12**: Approved and pending execution.
	//
	// example:
	//
	// 2
	CoStatus *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	// The type of the change order. Valid values:
	//
	// - **CoBindSlb**: Attach an SLB instance.
	//
	// - **CoUnbindSlb**: Detach an SLB instance.
	//
	// - **CoCreateApp**: Create an application.
	//
	// - **CoDeleteApp**: Delete an application.
	//
	// - **CoDeploy**: Deploy an application.
	//
	// - **CoRestartApplication**: Restart an application.
	//
	// - **CoRollback**: Roll back an application.
	//
	// - **CoScaleIn**: Scale in an application.
	//
	// - **CoScaleOut**: Scale out an application.
	//
	// - **CoStartApplication**: Start an application.
	//
	// - **CoStopApplication**: Stop an application.
	//
	// - **CoRescaleApplicationVertically**: Change the instance type.
	//
	// - **CoDeployHistroy**: Roll back to a previous version.
	//
	// - **CoBindNas**: Attach a NAS file system.
	//
	// - **CoUnbindNas**: Detach a NAS file system.
	//
	// - **CoBatchStartApplication**: Batch start applications.
	//
	// - **CoBatchStopApplication**: Batch stop applications.
	//
	// - **CoRestartInstances**: Restart instances.
	//
	// - **CoDeleteInstances**: Delete instances.
	//
	// - **CoScaleInAppWithInstances**: Scale in an application by specifying instances.
	//
	// example:
	//
	// CoCreateApp
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword for a fuzzy search of change order descriptions. The operation returns only the change orders whose descriptions contain the **keyword**.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The field by which to sort the results.
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order for the field specified by the **OrderBy*	- parameter. Valid values:
	//
	// - **true**: The results are sorted in ascending order.
	//
	// - **false**: The results are sorted in descending order.
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
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
