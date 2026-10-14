// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest
	GetDeletionProtection() *bool
	SetHostName(v string) *ModifyInstanceAttributeRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeRequest
	GetInstanceName() *string
	SetPassword(v string) *ModifyInstanceAttributeRequest
	GetPassword() *string
	SetUserData(v string) *ModifyInstanceAttributeRequest
	GetUserData() *string
}

type ModifyInstanceAttributeRequest struct {
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The hostname of the ECS instance. The hostname must be 2 to 64 characters in length. You can use periods (.) to separate the hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-). The hostname cannot contain consecutive periods (.) or hyphens (-), and cannot start or end with a period (.) or a hyphen (-).
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the instance to modify. Only a single instance is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instanc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// i-instanceidname
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The password of the instance.
	//
	// example:
	//
	// yourPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The instance user data, which must be Base64-encoded. The size of the raw data before encoding cannot exceed 16 KB. Do not pass in sensitive information such as passwords or private keys in plaintext. If you must pass in sensitive information, encode it before Base64-encoding it, and decode it in the same way inside the instance.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgK****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyInstanceAttributeRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceAttributeRequest) GetUserData() *string {
	return s.UserData
}

func (s *ModifyInstanceAttributeRequest) SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetHostName(v string) *ModifyInstanceAttributeRequest {
	s.HostName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPassword(v string) *ModifyInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetUserData(v string) *ModifyInstanceAttributeRequest {
	s.UserData = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
