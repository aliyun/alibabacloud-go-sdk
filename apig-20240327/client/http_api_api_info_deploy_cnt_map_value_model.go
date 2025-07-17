// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiApiInfoDeployCntMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetDeployedCnt(v int64) *HttpApiApiInfoDeployCntMapValue
	GetDeployedCnt() *int64
	SetCnt(v int64) *HttpApiApiInfoDeployCntMapValue
	GetCnt() *int64
}

type HttpApiApiInfoDeployCntMapValue struct {
	DeployedCnt *int64 `json:"deployedCnt,omitempty" xml:"deployedCnt,omitempty"`
	Cnt         *int64 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
}

func (s HttpApiApiInfoDeployCntMapValue) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoDeployCntMapValue) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoDeployCntMapValue) GetDeployedCnt() *int64 {
	return s.DeployedCnt
}

func (s *HttpApiApiInfoDeployCntMapValue) GetCnt() *int64 {
	return s.Cnt
}

func (s *HttpApiApiInfoDeployCntMapValue) SetDeployedCnt(v int64) *HttpApiApiInfoDeployCntMapValue {
	s.DeployedCnt = &v
	return s
}

func (s *HttpApiApiInfoDeployCntMapValue) SetCnt(v int64) *HttpApiApiInfoDeployCntMapValue {
	s.Cnt = &v
	return s
}

func (s *HttpApiApiInfoDeployCntMapValue) Validate() error {
	return dara.Validate(s)
}
