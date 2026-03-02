// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListEnvInfosRequest
	GetEnterpriseId() *int64
	SetEnv(v string) *ListEnvInfosRequest
	GetEnv() *string
	SetPbcId(v int64) *ListEnvInfosRequest
	GetPbcId() *int64
	SetRegion(v string) *ListEnvInfosRequest
	GetRegion() *string
}

type ListEnvInfosRequest struct {
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Env          *string `json:"env,omitempty" xml:"env,omitempty"`
	PbcId        *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	Region       *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s ListEnvInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvInfosRequest) GoString() string {
	return s.String()
}

func (s *ListEnvInfosRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListEnvInfosRequest) GetEnv() *string {
	return s.Env
}

func (s *ListEnvInfosRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListEnvInfosRequest) GetRegion() *string {
	return s.Region
}

func (s *ListEnvInfosRequest) SetEnterpriseId(v int64) *ListEnvInfosRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListEnvInfosRequest) SetEnv(v string) *ListEnvInfosRequest {
	s.Env = &v
	return s
}

func (s *ListEnvInfosRequest) SetPbcId(v int64) *ListEnvInfosRequest {
	s.PbcId = &v
	return s
}

func (s *ListEnvInfosRequest) SetRegion(v string) *ListEnvInfosRequest {
	s.Region = &v
	return s
}

func (s *ListEnvInfosRequest) Validate() error {
	return dara.Validate(s)
}
