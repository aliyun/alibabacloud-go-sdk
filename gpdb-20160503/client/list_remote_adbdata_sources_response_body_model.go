// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemoteADBDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceItems(v *ListRemoteADBDataSourcesResponseBodyDataSourceItems) *ListRemoteADBDataSourcesResponseBody
	GetDataSourceItems() *ListRemoteADBDataSourcesResponseBodyDataSourceItems
	SetRequestId(v string) *ListRemoteADBDataSourcesResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ListRemoteADBDataSourcesResponseBody
	GetTaskId() *int32
}

type ListRemoteADBDataSourcesResponseBody struct {
	// Returns the successfully added data sharing service data.
	DataSourceItems *ListRemoteADBDataSourcesResponseBodyDataSourceItems `json:"DataSourceItems,omitempty" xml:"DataSourceItems,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// e9d60eb1-e90d-4bc6-a470-c8b767460858
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 90000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListRemoteADBDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemoteADBDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemoteADBDataSourcesResponseBody) GetDataSourceItems() *ListRemoteADBDataSourcesResponseBodyDataSourceItems {
	return s.DataSourceItems
}

func (s *ListRemoteADBDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemoteADBDataSourcesResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ListRemoteADBDataSourcesResponseBody) SetDataSourceItems(v *ListRemoteADBDataSourcesResponseBodyDataSourceItems) *ListRemoteADBDataSourcesResponseBody {
	s.DataSourceItems = v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBody) SetRequestId(v string) *ListRemoteADBDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBody) SetTaskId(v int32) *ListRemoteADBDataSourcesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBody) Validate() error {
	if s.DataSourceItems != nil {
		if err := s.DataSourceItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRemoteADBDataSourcesResponseBodyDataSourceItems struct {
	RemoteDataSources []*ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources `json:"RemoteDataSources,omitempty" xml:"RemoteDataSources,omitempty" type:"Repeated"`
}

func (s ListRemoteADBDataSourcesResponseBodyDataSourceItems) String() string {
	return dara.Prettify(s)
}

func (s ListRemoteADBDataSourcesResponseBodyDataSourceItems) GoString() string {
	return s.String()
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItems) GetRemoteDataSources() []*ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	return s.RemoteDataSources
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItems) SetRemoteDataSources(v []*ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) *ListRemoteADBDataSourcesResponseBodyDataSourceItems {
	s.RemoteDataSources = v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItems) Validate() error {
	if s.RemoteDataSources != nil {
		for _, item := range s.RemoteDataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources struct {
	// Data source name
	//
	// example:
	//
	// db1_gptest1_to_db2_gp-test2
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Local database name
	//
	// example:
	//
	// db1
	LocalDatabase *string `json:"LocalDatabase,omitempty" xml:"LocalDatabase,omitempty"`
	// Local instance name
	//
	// example:
	//
	// gp-test1
	LocalInstanceName *string `json:"LocalInstanceName,omitempty" xml:"LocalInstanceName,omitempty"`
	// Manager user name
	//
	// example:
	//
	// admin
	ManagerUserName *string `json:"ManagerUserName,omitempty" xml:"ManagerUserName,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) API to view available region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Remote database name
	//
	// example:
	//
	// db2
	RemoteDatabase *string `json:"RemoteDatabase,omitempty" xml:"RemoteDatabase,omitempty"`
	// Remote instance name
	//
	// example:
	//
	// gp-test2
	RemoteInstanceName *string `json:"RemoteInstanceName,omitempty" xml:"RemoteInstanceName,omitempty"`
	// Data source status
	//
	// example:
	//
	// creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// User name
	//
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GoString() string {
	return s.String()
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetDescription() *string {
	return s.Description
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetId() *int64 {
	return s.Id
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetLocalDatabase() *string {
	return s.LocalDatabase
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetLocalInstanceName() *string {
	return s.LocalInstanceName
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetManagerUserName() *string {
	return s.ManagerUserName
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetRemoteDatabase() *string {
	return s.RemoteDatabase
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetRemoteInstanceName() *string {
	return s.RemoteInstanceName
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetStatus() *string {
	return s.Status
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) GetUserName() *string {
	return s.UserName
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetDataSourceName(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.DataSourceName = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetDescription(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.Description = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetId(v int64) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.Id = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetLocalDatabase(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.LocalDatabase = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetLocalInstanceName(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.LocalInstanceName = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetManagerUserName(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.ManagerUserName = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetRegionId(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.RegionId = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetRemoteDatabase(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.RemoteDatabase = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetRemoteInstanceName(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.RemoteInstanceName = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetStatus(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.Status = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) SetUserName(v string) *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources {
	s.UserName = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponseBodyDataSourceItemsRemoteDataSources) Validate() error {
	return dara.Validate(s)
}
