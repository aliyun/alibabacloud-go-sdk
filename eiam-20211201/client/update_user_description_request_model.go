// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateUserDescriptionRequest
	GetInstanceId() *string
	SetUserId(v string) *UpdateUserDescriptionRequest
	GetUserId() *string
}

type UpdateUserDescriptionRequest struct {
	// The description of the account. The value can be up to 256 characters in length.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserDescriptionRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserDescriptionRequest) SetDescription(v string) *UpdateUserDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserDescriptionRequest) SetInstanceId(v string) *UpdateUserDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserDescriptionRequest) SetUserId(v string) *UpdateUserDescriptionRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
