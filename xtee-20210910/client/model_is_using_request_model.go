// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelIsUsingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelCode(v string) *ModelIsUsingRequest
	GetModelCode() *string
	SetModelId(v string) *ModelIsUsingRequest
	GetModelId() *string
	SetModelName(v string) *ModelIsUsingRequest
	GetModelName() *string
	SetRegId(v string) *ModelIsUsingRequest
	GetRegId() *string
	SetStatus(v string) *ModelIsUsingRequest
	GetStatus() *string
}

type ModelIsUsingRequest struct {
	// Model code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ebgdsa
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// Model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27269
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// Model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v1
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Model status.
	//
	// This parameter is required.
	//
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModelIsUsingRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelIsUsingRequest) GoString() string {
	return s.String()
}

func (s *ModelIsUsingRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelIsUsingRequest) GetModelId() *string {
	return s.ModelId
}

func (s *ModelIsUsingRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModelIsUsingRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModelIsUsingRequest) GetStatus() *string {
	return s.Status
}

func (s *ModelIsUsingRequest) SetModelCode(v string) *ModelIsUsingRequest {
	s.ModelCode = &v
	return s
}

func (s *ModelIsUsingRequest) SetModelId(v string) *ModelIsUsingRequest {
	s.ModelId = &v
	return s
}

func (s *ModelIsUsingRequest) SetModelName(v string) *ModelIsUsingRequest {
	s.ModelName = &v
	return s
}

func (s *ModelIsUsingRequest) SetRegId(v string) *ModelIsUsingRequest {
	s.RegId = &v
	return s
}

func (s *ModelIsUsingRequest) SetStatus(v string) *ModelIsUsingRequest {
	s.Status = &v
	return s
}

func (s *ModelIsUsingRequest) Validate() error {
	return dara.Validate(s)
}
