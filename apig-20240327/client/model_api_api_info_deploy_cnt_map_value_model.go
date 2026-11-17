// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiApiInfoDeployCntMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetDeployedCnt(v int64) *ModelApiApiInfoDeployCntMapValue
	GetDeployedCnt() *int64
	SetCnt(v int64) *ModelApiApiInfoDeployCntMapValue
	GetCnt() *int64
}

type ModelApiApiInfoDeployCntMapValue struct {
	// The number of deployed APIs.
	//
	// example:
	//
	// 1
	DeployedCnt *int64 `json:"deployedCnt,omitempty" xml:"deployedCnt,omitempty"`
	// The total number of APIs.
	//
	// example:
	//
	// 1
	Cnt *int64 `json:"cnt,omitempty" xml:"cnt,omitempty"`
}

func (s ModelApiApiInfoDeployCntMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoDeployCntMapValue) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoDeployCntMapValue) GetDeployedCnt() *int64 {
	return s.DeployedCnt
}

func (s *ModelApiApiInfoDeployCntMapValue) GetCnt() *int64 {
	return s.Cnt
}

func (s *ModelApiApiInfoDeployCntMapValue) SetDeployedCnt(v int64) *ModelApiApiInfoDeployCntMapValue {
	s.DeployedCnt = &v
	return s
}

func (s *ModelApiApiInfoDeployCntMapValue) SetCnt(v int64) *ModelApiApiInfoDeployCntMapValue {
	s.Cnt = &v
	return s
}

func (s *ModelApiApiInfoDeployCntMapValue) Validate() error {
	return dara.Validate(s)
}
