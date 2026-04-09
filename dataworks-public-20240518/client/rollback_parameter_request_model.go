// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RollbackParameterRequest
	GetId() *int64
	SetRollbackVersion(v int32) *RollbackParameterRequest
	GetRollbackVersion() *int32
}

type RollbackParameterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RollbackVersion *int32 `json:"RollbackVersion,omitempty" xml:"RollbackVersion,omitempty"`
}

func (s RollbackParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackParameterRequest) GoString() string {
	return s.String()
}

func (s *RollbackParameterRequest) GetId() *int64 {
	return s.Id
}

func (s *RollbackParameterRequest) GetRollbackVersion() *int32 {
	return s.RollbackVersion
}

func (s *RollbackParameterRequest) SetId(v int64) *RollbackParameterRequest {
	s.Id = &v
	return s
}

func (s *RollbackParameterRequest) SetRollbackVersion(v int32) *RollbackParameterRequest {
	s.RollbackVersion = &v
	return s
}

func (s *RollbackParameterRequest) Validate() error {
	return dara.Validate(s)
}
