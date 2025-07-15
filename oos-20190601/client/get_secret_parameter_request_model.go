// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetSecretParameterRequest
	GetName() *string
	SetParameterVersion(v int32) *GetSecretParameterRequest
	GetParameterVersion() *int32
	SetRegionId(v string) *GetSecretParameterRequest
	GetRegionId() *string
	SetWithDecryption(v bool) *GetSecretParameterRequest
	GetWithDecryption() *bool
}

type GetSecretParameterRequest struct {
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. The decrypted parameter value is returned only if this parameter is set to true. Otherwise, null is returned.
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParameterRequest) GetName() *string {
	return s.Name
}

func (s *GetSecretParameterRequest) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetSecretParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSecretParameterRequest) GetWithDecryption() *bool {
	return s.WithDecryption
}

func (s *GetSecretParameterRequest) SetName(v string) *GetSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *GetSecretParameterRequest) SetParameterVersion(v int32) *GetSecretParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterRequest) SetRegionId(v string) *GetSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParameterRequest) SetWithDecryption(v bool) *GetSecretParameterRequest {
	s.WithDecryption = &v
	return s
}

func (s *GetSecretParameterRequest) Validate() error {
	return dara.Validate(s)
}
