// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNames(v string) *GetSecretParametersRequest
	GetNames() *string
	SetRegionId(v string) *GetSecretParametersRequest
	GetRegionId() *string
	SetWithDecryption(v bool) *GetSecretParametersRequest
	GetWithDecryption() *bool
}

type GetSecretParametersRequest struct {
	// The name of the encryption parameter. Multiple encryption parameters can form a JSON array in the format of ["xxxxxxxxx", "yyyyyyyyy", … "zzzzzzzzz"]. Each JSON array can contain a maximum of 10 encryption parameters. Multiple encryption parameters in the array are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["secretParameter","secretParameter1"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersRequest) GetNames() *string {
	return s.Names
}

func (s *GetSecretParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSecretParametersRequest) GetWithDecryption() *bool {
	return s.WithDecryption
}

func (s *GetSecretParametersRequest) SetNames(v string) *GetSecretParametersRequest {
	s.Names = &v
	return s
}

func (s *GetSecretParametersRequest) SetRegionId(v string) *GetSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersRequest) SetWithDecryption(v bool) *GetSecretParametersRequest {
	s.WithDecryption = &v
	return s
}

func (s *GetSecretParametersRequest) Validate() error {
	return dara.Validate(s)
}
