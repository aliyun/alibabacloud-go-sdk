// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResidentConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ResidentConfig
	GetCount() *int64
	SetPoolId(v string) *ResidentConfig
	GetPoolId() *string
}

type ResidentConfig struct {
	Count  *int64  `json:"count,omitempty" xml:"count,omitempty"`
	PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty"`
}

func (s ResidentConfig) String() string {
	return dara.Prettify(s)
}

func (s ResidentConfig) GoString() string {
	return s.String()
}

func (s *ResidentConfig) GetCount() *int64 {
	return s.Count
}

func (s *ResidentConfig) GetPoolId() *string {
	return s.PoolId
}

func (s *ResidentConfig) SetCount(v int64) *ResidentConfig {
	s.Count = &v
	return s
}

func (s *ResidentConfig) SetPoolId(v string) *ResidentConfig {
	s.PoolId = &v
	return s
}

func (s *ResidentConfig) Validate() error {
	return dara.Validate(s)
}
