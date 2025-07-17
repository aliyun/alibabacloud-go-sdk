// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishGroupTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *CreatePublishGroupTaskRequest
	GetDbId() *int32
	SetLogic(v bool) *CreatePublishGroupTaskRequest
	GetLogic() *bool
	SetOrderId(v int64) *CreatePublishGroupTaskRequest
	GetOrderId() *int64
	SetPlanTime(v string) *CreatePublishGroupTaskRequest
	GetPlanTime() *string
	SetPublishStrategy(v string) *CreatePublishGroupTaskRequest
	GetPublishStrategy() *string
	SetTid(v int64) *CreatePublishGroupTaskRequest
	GetTid() *int64
}

type CreatePublishGroupTaskRequest struct {
	// The ID of the database for which the schema design is executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Indicates whether the database is a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The ID of the ticket.
	//
	// > : You can create a schema design ticket in the DMS console. For more information, see [Design schemas](https://help.aliyun.com/document_detail/69711.html). You can also create a schema design ticket by calling the [CreateOrder](https://help.aliyun.com/document_detail/144649.html) operation and obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 142435
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The time to execute the schema design ticket.
	//
	// example:
	//
	// 2019-10-10 00:00:00
	PlanTime *string `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	// The policy to execute the schema design ticket. Valid values:
	//
	// 	- IMMEDIATELY: immediately executes the schema design ticket.
	//
	// 	- REGULARLY: executes the schema design ticket at a scheduled time.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMMEDIATELY
	PublishStrategy *string `json:"PublishStrategy,omitempty" xml:"PublishStrategy,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreatePublishGroupTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishGroupTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *CreatePublishGroupTaskRequest) GetLogic() *bool {
	return s.Logic
}

func (s *CreatePublishGroupTaskRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreatePublishGroupTaskRequest) GetPlanTime() *string {
	return s.PlanTime
}

func (s *CreatePublishGroupTaskRequest) GetPublishStrategy() *string {
	return s.PublishStrategy
}

func (s *CreatePublishGroupTaskRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreatePublishGroupTaskRequest) SetDbId(v int32) *CreatePublishGroupTaskRequest {
	s.DbId = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetLogic(v bool) *CreatePublishGroupTaskRequest {
	s.Logic = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetOrderId(v int64) *CreatePublishGroupTaskRequest {
	s.OrderId = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetPlanTime(v string) *CreatePublishGroupTaskRequest {
	s.PlanTime = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetPublishStrategy(v string) *CreatePublishGroupTaskRequest {
	s.PublishStrategy = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetTid(v int64) *CreatePublishGroupTaskRequest {
	s.Tid = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) Validate() error {
	return dara.Validate(s)
}
