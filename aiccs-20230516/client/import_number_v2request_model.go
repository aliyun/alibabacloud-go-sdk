// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCustomers(v []*ImportNumberV2RequestCustomers) *ImportNumberV2Request
	GetCustomers() []*ImportNumberV2RequestCustomers
	SetFailReturn(v int64) *ImportNumberV2Request
	GetFailReturn() *int64
	SetOutId(v string) *ImportNumberV2Request
	GetOutId() *string
	SetOwnerId(v int64) *ImportNumberV2Request
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ImportNumberV2Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportNumberV2Request
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportNumberV2Request
	GetTaskId() *int64
}

type ImportNumberV2Request struct {
	Customers []*ImportNumberV2RequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// example:
	//
	// 示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberV2Request) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2Request) GoString() string {
	return s.String()
}

func (s *ImportNumberV2Request) GetCustomers() []*ImportNumberV2RequestCustomers {
	return s.Customers
}

func (s *ImportNumberV2Request) GetFailReturn() *int64 {
	return s.FailReturn
}

func (s *ImportNumberV2Request) GetOutId() *string {
	return s.OutId
}

func (s *ImportNumberV2Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportNumberV2Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportNumberV2Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportNumberV2Request) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportNumberV2Request) SetCustomers(v []*ImportNumberV2RequestCustomers) *ImportNumberV2Request {
	s.Customers = v
	return s
}

func (s *ImportNumberV2Request) SetFailReturn(v int64) *ImportNumberV2Request {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberV2Request) SetOutId(v string) *ImportNumberV2Request {
	s.OutId = &v
	return s
}

func (s *ImportNumberV2Request) SetOwnerId(v int64) *ImportNumberV2Request {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberV2Request) SetResourceOwnerAccount(v string) *ImportNumberV2Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberV2Request) SetResourceOwnerId(v int64) *ImportNumberV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberV2Request) SetTaskId(v int64) *ImportNumberV2Request {
	s.TaskId = &v
	return s
}

func (s *ImportNumberV2Request) Validate() error {
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

type ImportNumberV2RequestCustomers struct {
	// example:
	//
	// 示例值示例值
	ClientUrl *string `json:"ClientUrl,omitempty" xml:"ClientUrl,omitempty"`
	// example:
	//
	// 示例值示例值
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 示例值示例值
	NumberMD5  *string                `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 示例值示例值
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ImportNumberV2RequestCustomers) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2RequestCustomers) GoString() string {
	return s.String()
}

func (s *ImportNumberV2RequestCustomers) GetClientUrl() *string {
	return s.ClientUrl
}

func (s *ImportNumberV2RequestCustomers) GetNumber() *string {
	return s.Number
}

func (s *ImportNumberV2RequestCustomers) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *ImportNumberV2RequestCustomers) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ImportNumberV2RequestCustomers) GetTag() *string {
	return s.Tag
}

func (s *ImportNumberV2RequestCustomers) SetClientUrl(v string) *ImportNumberV2RequestCustomers {
	s.ClientUrl = &v
	return s
}

func (s *ImportNumberV2RequestCustomers) SetNumber(v string) *ImportNumberV2RequestCustomers {
	s.Number = &v
	return s
}

func (s *ImportNumberV2RequestCustomers) SetNumberMD5(v string) *ImportNumberV2RequestCustomers {
	s.NumberMD5 = &v
	return s
}

func (s *ImportNumberV2RequestCustomers) SetProperties(v map[string]interface{}) *ImportNumberV2RequestCustomers {
	s.Properties = v
	return s
}

func (s *ImportNumberV2RequestCustomers) SetTag(v string) *ImportNumberV2RequestCustomers {
	s.Tag = &v
	return s
}

func (s *ImportNumberV2RequestCustomers) Validate() error {
	return dara.Validate(s)
}
