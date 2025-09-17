// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWorkgroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyWorkgroupAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyWorkgroupAttributeRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyWorkgroupAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyWorkgroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetWorkgroupId(v string) *ModifyWorkgroupAttributeRequest
	GetWorkgroupId() *string
}

type ModifyWorkgroupAttributeRequest struct {
	// The new description of the workgroup. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workgroup. The name must meet the following requirements:
	//
	// 	- The name must be unique.
	//
	// 	- The name must be 2 to 64 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testMigrationTaskName
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the workgroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-***
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s ModifyWorkgroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWorkgroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWorkgroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyWorkgroupAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyWorkgroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyWorkgroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyWorkgroupAttributeRequest) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *ModifyWorkgroupAttributeRequest) SetDescription(v string) *ModifyWorkgroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyWorkgroupAttributeRequest) SetName(v string) *ModifyWorkgroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyWorkgroupAttributeRequest) SetOwnerId(v int64) *ModifyWorkgroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyWorkgroupAttributeRequest) SetResourceOwnerAccount(v string) *ModifyWorkgroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyWorkgroupAttributeRequest) SetWorkgroupId(v string) *ModifyWorkgroupAttributeRequest {
	s.WorkgroupId = &v
	return s
}

func (s *ModifyWorkgroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
