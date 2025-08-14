// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelName(v string) *ModelDeleteRequest
	GetModelName() *string
	SetRegId(v string) *ModelDeleteRequest
	GetRegId() *string
}

type ModelDeleteRequest struct {
	// Model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s ModelDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelDeleteRequest) GoString() string {
	return s.String()
}

func (s *ModelDeleteRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModelDeleteRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModelDeleteRequest) SetModelName(v string) *ModelDeleteRequest {
	s.ModelName = &v
	return s
}

func (s *ModelDeleteRequest) SetRegId(v string) *ModelDeleteRequest {
	s.RegId = &v
	return s
}

func (s *ModelDeleteRequest) Validate() error {
	return dara.Validate(s)
}
