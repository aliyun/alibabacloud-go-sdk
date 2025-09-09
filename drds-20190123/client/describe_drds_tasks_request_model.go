// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsTasksRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsTasksRequest
	GetDrdsInstanceId() *string
	SetTaskType(v string) *DescribeDrdsTasksRequest
	GetTaskType() *string
}

type DescribeDrdsTasksRequest struct {
	// The name of the database.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsxxxxxxxxxxxx
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of tasks.
	//
	// example:
	//
	// test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsTasksRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDrdsTasksRequest) SetDbName(v string) *DescribeDrdsTasksRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsTasksRequest) SetTaskType(v string) *DescribeDrdsTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeDrdsTasksRequest) Validate() error {
	return dara.Validate(s)
}
