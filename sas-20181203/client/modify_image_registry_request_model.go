// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyImageRegistryRequest
	GetId() *int64
	SetPassword(v string) *ModifyImageRegistryRequest
	GetPassword() *string
	SetTransPerHour(v int32) *ModifyImageRegistryRequest
	GetTransPerHour() *int32
	SetUserName(v string) *ModifyImageRegistryRequest
	GetUserName() *string
}

type ModifyImageRegistryRequest struct {
	// The ID of the image repository. You can call the listImageRegistry operation to query the ID of the image repository.
	//
	// example:
	//
	// 390103286
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The password.
	//
	// example:
	//
	// ********************
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The number of images that are scanned per hour.
	//
	// example:
	//
	// 10
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username.
	//
	// example:
	//
	// xxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyImageRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageRegistryRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageRegistryRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyImageRegistryRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyImageRegistryRequest) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *ModifyImageRegistryRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyImageRegistryRequest) SetId(v int64) *ModifyImageRegistryRequest {
	s.Id = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetPassword(v string) *ModifyImageRegistryRequest {
	s.Password = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetTransPerHour(v int32) *ModifyImageRegistryRequest {
	s.TransPerHour = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetUserName(v string) *ModifyImageRegistryRequest {
	s.UserName = &v
	return s
}

func (s *ModifyImageRegistryRequest) Validate() error {
	return dara.Validate(s)
}
