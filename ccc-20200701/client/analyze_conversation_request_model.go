// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *AnalyzeConversationRequest
	GetContactId() *string
	SetFieldListJson(v string) *AnalyzeConversationRequest
	GetFieldListJson() *string
	SetInstanceId(v string) *AnalyzeConversationRequest
	GetInstanceId() *string
	SetTaskListJson(v string) *AnalyzeConversationRequest
	GetTaskListJson() *string
}

type AnalyzeConversationRequest struct {
	// Call ID.
	//
	// example:
	//
	// job-10963442671187****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// When the TaskListJson parameter is set to "fields", this parameter specifies the field content to fetch.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "code": "OrderNo",
	//
	//     "name": "订单号",
	//
	//     "description": "订单号"
	//
	//   },
	//
	//   {
	//
	//     "code": "ProductName",
	//
	//     "name": "商品名称",
	//
	//     "description": "商品名称"
	//
	//   }
	//
	// ]
	FieldListJson *string `json:"FieldListJson,omitempty" xml:"FieldListJson,omitempty"`
	// Cloud Contact Center instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of Task Types for analysis. Valid values include:
	//
	// - title: title
	//
	// - summary: summary
	//
	// - keywords: keywords
	//
	// - fields: fields
	//
	// Currently, only "summary" and "title" can be passed together. All other parameters must be passed individually.
	//
	// example:
	//
	// ["keywords"]
	TaskListJson *string `json:"TaskListJson,omitempty" xml:"TaskListJson,omitempty"`
}

func (s AnalyzeConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequest) GetContactId() *string {
	return s.ContactId
}

func (s *AnalyzeConversationRequest) GetFieldListJson() *string {
	return s.FieldListJson
}

func (s *AnalyzeConversationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AnalyzeConversationRequest) GetTaskListJson() *string {
	return s.TaskListJson
}

func (s *AnalyzeConversationRequest) SetContactId(v string) *AnalyzeConversationRequest {
	s.ContactId = &v
	return s
}

func (s *AnalyzeConversationRequest) SetFieldListJson(v string) *AnalyzeConversationRequest {
	s.FieldListJson = &v
	return s
}

func (s *AnalyzeConversationRequest) SetInstanceId(v string) *AnalyzeConversationRequest {
	s.InstanceId = &v
	return s
}

func (s *AnalyzeConversationRequest) SetTaskListJson(v string) *AnalyzeConversationRequest {
	s.TaskListJson = &v
	return s
}

func (s *AnalyzeConversationRequest) Validate() error {
	return dara.Validate(s)
}
