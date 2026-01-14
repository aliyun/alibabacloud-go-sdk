// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomers(v []*ImportNumberRequestCustomers) *ImportNumberRequest
	GetCustomers() []*ImportNumberRequestCustomers
	SetFailReturn(v int64) *ImportNumberRequest
	GetFailReturn() *int64
	SetOutId(v string) *ImportNumberRequest
	GetOutId() *string
	SetOwnerId(v int64) *ImportNumberRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ImportNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportNumberRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportNumberRequest
	GetTaskId() *int64
}

type ImportNumberRequest struct {
	// This parameter is required.
	Customers []*ImportNumberRequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberRequest) GoString() string {
	return s.String()
}

func (s *ImportNumberRequest) GetCustomers() []*ImportNumberRequestCustomers {
	return s.Customers
}

func (s *ImportNumberRequest) GetFailReturn() *int64 {
	return s.FailReturn
}

func (s *ImportNumberRequest) GetOutId() *string {
	return s.OutId
}

func (s *ImportNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportNumberRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportNumberRequest) SetCustomers(v []*ImportNumberRequestCustomers) *ImportNumberRequest {
	s.Customers = v
	return s
}

func (s *ImportNumberRequest) SetFailReturn(v int64) *ImportNumberRequest {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberRequest) SetOutId(v string) *ImportNumberRequest {
	s.OutId = &v
	return s
}

func (s *ImportNumberRequest) SetOwnerId(v int64) *ImportNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberRequest) SetResourceOwnerAccount(v string) *ImportNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberRequest) SetResourceOwnerId(v int64) *ImportNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberRequest) SetTaskId(v int64) *ImportNumberRequest {
	s.TaskId = &v
	return s
}

func (s *ImportNumberRequest) Validate() error {
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

type ImportNumberRequestCustomers struct {
	// example:
	//
	// http://test.com
	ClientUrl *string `json:"ClientUrl,omitempty" xml:"ClientUrl,omitempty"`
	// example:
	//
	// 13541251222,18665214444
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 示例值
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// example:
	//
	// {"testt":"123"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ImportNumberRequestCustomers) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberRequestCustomers) GoString() string {
	return s.String()
}

func (s *ImportNumberRequestCustomers) GetClientUrl() *string {
	return s.ClientUrl
}

func (s *ImportNumberRequestCustomers) GetNumber() *string {
	return s.Number
}

func (s *ImportNumberRequestCustomers) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *ImportNumberRequestCustomers) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ImportNumberRequestCustomers) GetTag() *string {
	return s.Tag
}

func (s *ImportNumberRequestCustomers) SetClientUrl(v string) *ImportNumberRequestCustomers {
	s.ClientUrl = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetNumber(v string) *ImportNumberRequestCustomers {
	s.Number = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetNumberMD5(v string) *ImportNumberRequestCustomers {
	s.NumberMD5 = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetProperties(v map[string]interface{}) *ImportNumberRequestCustomers {
	s.Properties = v
	return s
}

func (s *ImportNumberRequestCustomers) SetTag(v string) *ImportNumberRequestCustomers {
	s.Tag = &v
	return s
}

func (s *ImportNumberRequestCustomers) Validate() error {
	return dara.Validate(s)
}
