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
	// - **0**: preparing
	//
	// - **1**: executing
	//
	// - **2**: executed successfully
	//
	// - **3**: execution failed
	//
	// - **6**: terminated
	//
	// - **8**: waiting for manual batch confirmation
	//
	// - **9**: waiting for automatic batch confirmation
	//
	// - **10**: execution failed due to a system exception
	//
	// - **11**: waiting for approval
	//
	// - **12**: approved and waiting for execution
	//
	// example:
	//
	// 2
	CoStatus *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	// The type of the change order. Valid values:
	//
	// - **CoBindSlb**: bind an SLB instance
	//
	// - **CoUnbindSlb**: unbind an SLB instance
	//
	// - **CoCreateApp**: create an application
	//
	// - **CoDeleteApp**: delete an application
	//
	// - **CoDeploy**: deploy an application
	//
	// - **CoRestartApplication**: restart an application
	//
	// - **CoRollback**: roll back an application
	//
	// - **CoScaleIn**: scale in an application
	//
	// - **CoScaleOut**: scale out an application
	//
	// - **CoStartApplication**: start an application
	//
	// - **CoStopApplication**: stop an application
	//
	// - **CoRescaleApplicationVertically**: change instance specifications
	//
	// - **CoDeployHistroy**: roll back to a historical version
	//
	// - **CoBindNas**: bind a NAS file system
	//
	// - **CoUnbindNas**: unbind a NAS file system
	//
	// - **CoBatchStartApplication**: start applications in batches
	//
	// - **CoBatchStopApplication**: stop applications in batches
	//
	// - **CoRestartInstances**: restart instances
	//
	// - **CoDeleteInstances**: delete instances
	//
	// - **CoScaleInAppWithInstances**: scale in specified instances
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
	// The keyword used for fuzzy match of the change order description. Change orders whose descriptions contain the keyword are returned.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The sorting condition.
	//
	// example:
	//
	// create_time
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to sort the results in ascending order based on the field specified by the **OrderBy*	- parameter. Valid values:
	//
	//   - **true**: sort in ascending order
	//
	//   - **false**: sort in descending order
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
