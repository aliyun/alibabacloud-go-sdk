// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemoteADBDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceItem(v *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) *ModifyRemoteADBDataSourceResponseBody
	GetDataSourceItem() *ModifyRemoteADBDataSourceResponseBodyDataSourceItem
	SetRequestId(v string) *ModifyRemoteADBDataSourceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ModifyRemoteADBDataSourceResponseBody
	GetTaskId() *int32
}

type ModifyRemoteADBDataSourceResponseBody struct {
	// Returns the successfully modified data sharing service data.
	DataSourceItem *ModifyRemoteADBDataSourceResponseBodyDataSourceItem `json:"DataSourceItem,omitempty" xml:"DataSourceItem,omitempty" type:"Struct"`
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

func (s ModifyRemoteADBDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemoteADBDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRemoteADBDataSourceResponseBody) GetDataSourceItem() *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	return s.DataSourceItem
}

func (s *ModifyRemoteADBDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRemoteADBDataSourceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyRemoteADBDataSourceResponseBody) SetDataSourceItem(v *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) *ModifyRemoteADBDataSourceResponseBody {
	s.DataSourceItem = v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBody) SetRequestId(v string) *ModifyRemoteADBDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBody) SetTaskId(v int32) *ModifyRemoteADBDataSourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyRemoteADBDataSourceResponseBodyDataSourceItem struct {
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
	// test
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
	// Manager user name
	//
	// example:
	//
	// test
	ManagerUserName *string `json:"ManagerUserName,omitempty" xml:"ManagerUserName,omitempty"`
	// Region ID where the instance is located.
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

func (s ModifyRemoteADBDataSourceResponseBodyDataSourceItem) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GoString() string {
	return s.String()
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetDescription() *string {
	return s.Description
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetId() *int64 {
	return s.Id
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetLocalDatabase() *string {
	return s.LocalDatabase
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetLocalInstanceName() *string {
	return s.LocalInstanceName
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetManagerUserName() *string {
	return s.ManagerUserName
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetRemoteDatabase() *string {
	return s.RemoteDatabase
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetRemoteInstanceName() *string {
	return s.RemoteInstanceName
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetStatus() *string {
	return s.Status
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) GetUserName() *string {
	return s.UserName
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetDataSourceName(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.DataSourceName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetDescription(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Description = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetId(v int64) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Id = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetLocalDatabase(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.LocalDatabase = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetLocalInstanceName(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.LocalInstanceName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetManagerUserName(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.ManagerUserName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetRegionId(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RegionId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetRemoteDatabase(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RemoteDatabase = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetRemoteInstanceName(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.RemoteInstanceName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetStatus(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.Status = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) SetUserName(v string) *ModifyRemoteADBDataSourceResponseBodyDataSourceItem {
	s.UserName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceResponseBodyDataSourceItem) Validate() error {
	return dara.Validate(s)
}
