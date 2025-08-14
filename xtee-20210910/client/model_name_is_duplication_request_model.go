// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelNameIsDuplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelName(v string) *ModelNameIsDuplicationRequest
	GetModelName() *string
	SetRegId(v string) *ModelNameIsDuplicationRequest
	GetRegId() *string
}

type ModelNameIsDuplicationRequest struct {
	// Model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v2
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s ModelNameIsDuplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelNameIsDuplicationRequest) GoString() string {
	return s.String()
}

func (s *ModelNameIsDuplicationRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModelNameIsDuplicationRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModelNameIsDuplicationRequest) SetModelName(v string) *ModelNameIsDuplicationRequest {
	s.ModelName = &v
	return s
}

func (s *ModelNameIsDuplicationRequest) SetRegId(v string) *ModelNameIsDuplicationRequest {
	s.RegId = &v
	return s
}

func (s *ModelNameIsDuplicationRequest) Validate() error {
	return dara.Validate(s)
}
