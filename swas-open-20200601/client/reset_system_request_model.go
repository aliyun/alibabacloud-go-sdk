// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResetSystemRequest
	GetClientToken() *string
	SetImageId(v string) *ResetSystemRequest
	GetImageId() *string
	SetInstanceId(v string) *ResetSystemRequest
	GetInstanceId() *string
	SetLoginCredentials(v *ResetSystemRequestLoginCredentials) *ResetSystemRequest
	GetLoginCredentials() *ResetSystemRequestLoginCredentials
	SetRegionId(v string) *ResetSystemRequest
	GetRegionId() *string
}

type ResetSystemRequest struct {
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the destination image. If you do not specify this parameter, the current image is reset.
	//
	// example:
	//
	// 794c230fd3e64ea19f83f4d7a0ad****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The information about the logon credentials.
	LoginCredentials *ResetSystemRequestLoginCredentials `json:"LoginCredentials,omitempty" xml:"LoginCredentials,omitempty" type:"Struct"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetSystemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResetSystemRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ResetSystemRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetSystemRequest) GetLoginCredentials() *ResetSystemRequestLoginCredentials {
	return s.LoginCredentials
}

func (s *ResetSystemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSystemRequest) SetClientToken(v string) *ResetSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetSystemRequest) SetImageId(v string) *ResetSystemRequest {
	s.ImageId = &v
	return s
}

func (s *ResetSystemRequest) SetInstanceId(v string) *ResetSystemRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetSystemRequest) SetLoginCredentials(v *ResetSystemRequestLoginCredentials) *ResetSystemRequest {
	s.LoginCredentials = v
	return s
}

func (s *ResetSystemRequest) SetRegionId(v string) *ResetSystemRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSystemRequest) Validate() error {
	if s.LoginCredentials != nil {
		if err := s.LoginCredentials.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetSystemRequestLoginCredentials struct {
	// The name of the key pair that you want to bind to the simple application server. Only Linux simple application servers support key pairs.
	//
	// example:
	//
	// aliyunkey
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password that corresponds to the administrator.
	//
	// - For Windows servers, the default administrator account is Administrator.
	//
	// - Linux servers, the default administrator account is root.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ResetSystemRequestLoginCredentials) String() string {
	return dara.Prettify(s)
}

func (s ResetSystemRequestLoginCredentials) GoString() string {
	return s.String()
}

func (s *ResetSystemRequestLoginCredentials) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ResetSystemRequestLoginCredentials) GetPassword() *string {
	return s.Password
}

func (s *ResetSystemRequestLoginCredentials) SetKeyPairName(v string) *ResetSystemRequestLoginCredentials {
	s.KeyPairName = &v
	return s
}

func (s *ResetSystemRequestLoginCredentials) SetPassword(v string) *ResetSystemRequestLoginCredentials {
	s.Password = &v
	return s
}

func (s *ResetSystemRequestLoginCredentials) Validate() error {
	return dara.Validate(s)
}
