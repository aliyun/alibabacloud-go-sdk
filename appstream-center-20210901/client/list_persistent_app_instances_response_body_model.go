// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersistentAppInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPersistentAppInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPersistentAppInstancesResponseBody
	GetPageSize() *int32
	SetPersistentAppInstanceModels(v []*ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) *ListPersistentAppInstancesResponseBody
	GetPersistentAppInstanceModels() []*ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels
	SetRequestId(v string) *ListPersistentAppInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPersistentAppInstancesResponseBody
	GetTotalCount() *int32
}

type ListPersistentAppInstancesResponseBody struct {
	// The current page number of the query results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of persistent session application instances.
	PersistentAppInstanceModels []*ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels `json:"PersistentAppInstanceModels,omitempty" xml:"PersistentAppInstanceModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPersistentAppInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPersistentAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersistentAppInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPersistentAppInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPersistentAppInstancesResponseBody) GetPersistentAppInstanceModels() []*ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	return s.PersistentAppInstanceModels
}

func (s *ListPersistentAppInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPersistentAppInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPersistentAppInstancesResponseBody) SetPageNumber(v int32) *ListPersistentAppInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBody) SetPageSize(v int32) *ListPersistentAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBody) SetPersistentAppInstanceModels(v []*ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) *ListPersistentAppInstancesResponseBody {
	s.PersistentAppInstanceModels = v
	return s
}

func (s *ListPersistentAppInstancesResponseBody) SetRequestId(v string) *ListPersistentAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBody) SetTotalCount(v int32) *ListPersistentAppInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBody) Validate() error {
	if s.PersistentAppInstanceModels != nil {
		for _, item := range s.PersistentAppInstanceModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels struct {
	// The delivery group ID.
	//
	// example:
	//
	// aig-0bxls9m9arax6****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application instance ID.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The persistent session ID.
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The name of the persistent session application instance.
	//
	// example:
	//
	// test-persistent-name
	AppInstancePersistentName *string `json:"AppInstancePersistentName,omitempty" xml:"AppInstancePersistentName,omitempty"`
	// The instance status of the persistent session application.
	//
	// example:
	//
	// RUNNING
	AppInstancePersistentStatus *string `json:"AppInstancePersistentStatus,omitempty" xml:"AppInstancePersistentStatus,omitempty"`
	// The application instance status.
	//
	// example:
	//
	// RUNNING
	AppInstanceStatus *string `json:"AppInstanceStatus,omitempty" xml:"AppInstanceStatus,omitempty"`
	// The list of authorized usernames.
	AuthorizedUsers []*string `json:"AuthorizedUsers,omitempty" xml:"AuthorizedUsers,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2025-03-13T03:22:18.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GoString() string {
	return s.String()
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstancePersistentName() *string {
	return s.AppInstancePersistentName
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstancePersistentStatus() *string {
	return s.AppInstancePersistentStatus
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAppInstanceStatus() *string {
	return s.AppInstanceStatus
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetAuthorizedUsers() []*string {
	return s.AuthorizedUsers
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstanceGroupId(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstanceId(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstancePersistentId(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstancePersistentId = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstancePersistentName(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstancePersistentName = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstancePersistentStatus(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstancePersistentStatus = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAppInstanceStatus(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AppInstanceStatus = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetAuthorizedUsers(v []*string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.AuthorizedUsers = v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) SetGmtCreate(v string) *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels {
	s.GmtCreate = &v
	return s
}

func (s *ListPersistentAppInstancesResponseBodyPersistentAppInstanceModels) Validate() error {
	return dara.Validate(s)
}
