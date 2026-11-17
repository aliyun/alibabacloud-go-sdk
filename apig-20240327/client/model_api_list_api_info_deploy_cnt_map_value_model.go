// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiListApiInfoDeployCntMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetDeployedCnt(v int64) *ModelApiListApiInfoDeployCntMapValue
	GetDeployedCnt() *int64
	SetCnt(v int64) *ModelApiListApiInfoDeployCntMapValue
	GetCnt() *int64
}

type ModelApiListApiInfoDeployCntMapValue struct {
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

func (s ModelApiListApiInfoDeployCntMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoDeployCntMapValue) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoDeployCntMapValue) GetDeployedCnt() *int64 {
	return s.DeployedCnt
}

func (s *ModelApiListApiInfoDeployCntMapValue) GetCnt() *int64 {
	return s.Cnt
}

func (s *ModelApiListApiInfoDeployCntMapValue) SetDeployedCnt(v int64) *ModelApiListApiInfoDeployCntMapValue {
	s.DeployedCnt = &v
	return s
}

func (s *ModelApiListApiInfoDeployCntMapValue) SetCnt(v int64) *ModelApiListApiInfoDeployCntMapValue {
	s.Cnt = &v
	return s
}

func (s *ModelApiListApiInfoDeployCntMapValue) Validate() error {
	return dara.Validate(s)
}
