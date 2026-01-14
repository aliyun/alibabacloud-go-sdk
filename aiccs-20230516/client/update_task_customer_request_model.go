// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomers(v []*UpdateTaskCustomerRequestCustomers) *UpdateTaskCustomerRequest
	GetCustomers() []*UpdateTaskCustomerRequestCustomers
	SetOwnerId(v int64) *UpdateTaskCustomerRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTaskCustomerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTaskCustomerRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *UpdateTaskCustomerRequest
	GetTaskId() *int64
}

type UpdateTaskCustomerRequest struct {
	// 外呼客户
	//
	// This parameter is required.
	Customers            []*UpdateTaskCustomerRequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
	OwnerId              *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 59
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTaskCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerRequest) GetCustomers() []*UpdateTaskCustomerRequestCustomers {
	return s.Customers
}

func (s *UpdateTaskCustomerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTaskCustomerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTaskCustomerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTaskCustomerRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateTaskCustomerRequest) SetCustomers(v []*UpdateTaskCustomerRequestCustomers) *UpdateTaskCustomerRequest {
	s.Customers = v
	return s
}

func (s *UpdateTaskCustomerRequest) SetOwnerId(v int64) *UpdateTaskCustomerRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetResourceOwnerAccount(v string) *UpdateTaskCustomerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetResourceOwnerId(v int64) *UpdateTaskCustomerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetTaskId(v int64) *UpdateTaskCustomerRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTaskCustomerRequest) Validate() error {
	if s.Customers != nil {
		for _, item := range s.Customers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTaskCustomerRequestCustomers struct {
	// 是否取消外呼 0否，1是
	//
	// example:
	//
	// 0
	Cancel *int64 `json:"Cancel,omitempty" xml:"Cancel,omitempty"`
	// 电话号码
	//
	// example:
	//
	// 13661109390
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 需根据具体任务参数要求传输
	//
	// example:
	//
	// {"test":"234"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s UpdateTaskCustomerRequestCustomers) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerRequestCustomers) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerRequestCustomers) GetCancel() *int64 {
	return s.Cancel
}

func (s *UpdateTaskCustomerRequestCustomers) GetNumber() *string {
	return s.Number
}

func (s *UpdateTaskCustomerRequestCustomers) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *UpdateTaskCustomerRequestCustomers) GetTag() *string {
	return s.Tag
}

func (s *UpdateTaskCustomerRequestCustomers) SetCancel(v int64) *UpdateTaskCustomerRequestCustomers {
	s.Cancel = &v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetNumber(v string) *UpdateTaskCustomerRequestCustomers {
	s.Number = &v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetProperties(v map[string]interface{}) *UpdateTaskCustomerRequestCustomers {
	s.Properties = v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetTag(v string) *UpdateTaskCustomerRequestCustomers {
	s.Tag = &v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) Validate() error {
	return dara.Validate(s)
}
