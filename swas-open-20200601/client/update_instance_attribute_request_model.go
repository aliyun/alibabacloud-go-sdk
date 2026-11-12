// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateInstanceAttributeRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceAttributeRequest
	GetInstanceName() *string
	SetPassword(v string) *UpdateInstanceAttributeRequest
	GetPassword() *string
	SetRegionId(v string) *UpdateInstanceAttributeRequest
	GetRegionId() *string
}

type UpdateInstanceAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- parameter supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server. The name must be 2 to 50 characters in length, and must start with a letter, digit, or Chinese character. It can contain digits, underscores (_), and hyphens (-). Domain names are also supported.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The new password of the server. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	// ```
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// ```
	//
	// The password of a Windows instance cannot start with a forward slash (/).
	//
	// > If you specify the `Password` parameter, send the request over HTTPS to prevent password leaks.
	//
	// example:
	//
	// Test123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceAttributeRequest) SetClientToken(v string) *UpdateInstanceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceId(v string) *UpdateInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceName(v string) *UpdateInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetPassword(v string) *UpdateInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetRegionId(v string) *UpdateInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
