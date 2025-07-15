// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteSecretParameterRequest
	GetName() *string
	SetRegionId(v string) *DeleteSecretParameterRequest
	GetRegionId() *string
}

type DeleteSecretParameterRequest struct {
	// The name of the encryption parameter. The name must be 1 to 180 characters in length and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSecretParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterRequest) GetName() *string {
	return s.Name
}

func (s *DeleteSecretParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSecretParameterRequest) SetName(v string) *DeleteSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *DeleteSecretParameterRequest) SetRegionId(v string) *DeleteSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecretParameterRequest) Validate() error {
	return dara.Validate(s)
}
