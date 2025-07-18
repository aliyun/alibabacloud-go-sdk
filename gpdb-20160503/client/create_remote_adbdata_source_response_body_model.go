// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemoteADBDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceItem(v *CreateRemoteADBDataSourceResponseBodyDataSourceItem) *CreateRemoteADBDataSourceResponseBody
	GetDataSourceItem() *CreateRemoteADBDataSourceResponseBodyDataSourceItem
	SetRequestId(v string) *CreateRemoteADBDataSourceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *CreateRemoteADBDataSourceResponseBody
	GetTaskId() *int32
}

type CreateRemoteADBDataSourceResponseBody struct {
	// Returns the successfully added data sharing service data.
	DataSourceItem *CreateRemoteADBDataSourceResponseBodyDataSourceItem `json:"DataSourceItem,omitempty" xml:"DataSourceItem,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// e9d60eb1-e90d-4bc6-a470-c8b767460858
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Workflow task ID.
	//
	// example:
	//
	// 90000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateRemoteADBDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRemoteADBDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemoteADBDataSourceResponseBody) GetDataSourceItem() *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	return s.DataSourceItem
}

func (s *CreateRemoteADBDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRemoteADBDataSourceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateRemoteADBDataSourceResponseBody) SetDataSourceItem(v *CreateRemoteADBDataSourceResponseBodyDataSourceItem) *CreateRemoteADBDataSourceResponseBody {
	s.DataSourceItem = v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBody) SetRequestId(v string) *CreateRemoteADBDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBody) SetTaskId(v int32) *CreateRemoteADBDataSourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRemoteADBDataSourceResponseBodyDataSourceItem struct {
	// Data source name.
	//
	// example:
	//
	// db1_gptest1_to_db2_gp-test2
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Description information.
	//
	// example:
	//
	// userName
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
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
	// Management account user name
	//
	// example:
	//
	// manager_user
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
	// Synchronization status
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

func (s CreateRemoteADBDataSourceResponseBodyDataSourceItem) String() string {
	return dara.Prettify(s)
}

func (s CreateRemoteADBDataSourceResponseBodyDataSourceItem) GoString() string {
	return s.String()
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetDescription() *string {
	return s.Description
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetId() *int64 {
	return s.Id
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetLocalDatabase() *string {
	return s.LocalDatabase
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetLocalInstanceName() *string {
	return s.LocalInstanceName
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetManagerUserName() *string {
	return s.ManagerUserName
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetRemoteDatabase() *string {
	return s.RemoteDatabase
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetRemoteInstanceName() *string {
	return s.RemoteInstanceName
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetStatus() *string {
	return s.Status
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) GetUserName() *string {
	return s.UserName
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetDataSourceName(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.DataSourceName = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetDescription(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Description = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetId(v int64) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Id = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetLocalDatabase(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.LocalDatabase = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetLocalInstanceName(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.LocalInstanceName = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetManagerUserName(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.ManagerUserName = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetRegionId(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RegionId = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetRemoteDatabase(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RemoteDatabase = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetRemoteInstanceName(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RemoteInstanceName = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetStatus(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Status = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) SetUserName(v string) *CreateRemoteADBDataSourceResponseBodyDataSourceItem {
	s.UserName = &v
	return s
}

func (s *CreateRemoteADBDataSourceResponseBodyDataSourceItem) Validate() error {
	return dara.Validate(s)
}
