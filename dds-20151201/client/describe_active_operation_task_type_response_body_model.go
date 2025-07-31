// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeActiveOperationTaskTypeResponseBody
	GetRequestId() *string
	SetTypeList(v []*DescribeActiveOperationTaskTypeResponseBodyTypeList) *DescribeActiveOperationTaskTypeResponseBody
	GetTypeList() []*DescribeActiveOperationTaskTypeResponseBodyTypeList
}

type DescribeActiveOperationTaskTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C7EE83BF-7BA8-5087-BAC9-ED85ED54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The O\\&M tasks.
	TypeList []*DescribeActiveOperationTaskTypeResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeActiveOperationTaskTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTaskTypeResponseBody) GetTypeList() []*DescribeActiveOperationTaskTypeResponseBodyTypeList {
	return s.TypeList
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetTypeList(v []*DescribeActiveOperationTaskTypeResponseBodyTypeList) *DescribeActiveOperationTaskTypeResponseBody {
	s.TypeList = v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveOperationTaskTypeResponseBodyTypeList struct {
	// The number of pending tasks.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **rds_apsaradb_transfer**: data migration
	//
	// 	- **rds_apsaradb_upgrade**: minor version update
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The task type in English.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskTypeInfoEn *string `json:"TaskTypeInfoEn,omitempty" xml:"TaskTypeInfoEn,omitempty"`
	// The task type in Chinese.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskTypeInfoZh *string `json:"TaskTypeInfoZh,omitempty" xml:"TaskTypeInfoZh,omitempty"`
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) GetTaskTypeInfoEn() *string {
	return s.TaskTypeInfoEn
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) GetTaskTypeInfoZh() *string {
	return s.TaskTypeInfoZh
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetCount(v int32) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.Count = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskType(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskTypeInfoEn(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskTypeInfoEn = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskTypeInfoZh(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskTypeInfoZh = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) Validate() error {
	return dara.Validate(s)
}
