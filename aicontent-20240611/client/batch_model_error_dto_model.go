// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModelErrorDTO interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *BatchModelErrorDTO
	GetErrorMsg() *string
	SetModelId(v string) *BatchModelErrorDTO
	GetModelId() *string
	SetName(v string) *BatchModelErrorDTO
	GetName() *string
}

type BatchModelErrorDTO struct {
	// The failure cause.
	//
	// example:
	//
	// The model type cannot be empty
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The model ID.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-max
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BatchModelErrorDTO) String() string {
	return dara.Prettify(s)
}

func (s BatchModelErrorDTO) GoString() string {
	return s.String()
}

func (s *BatchModelErrorDTO) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BatchModelErrorDTO) GetModelId() *string {
	return s.ModelId
}

func (s *BatchModelErrorDTO) GetName() *string {
	return s.Name
}

func (s *BatchModelErrorDTO) SetErrorMsg(v string) *BatchModelErrorDTO {
	s.ErrorMsg = &v
	return s
}

func (s *BatchModelErrorDTO) SetModelId(v string) *BatchModelErrorDTO {
	s.ModelId = &v
	return s
}

func (s *BatchModelErrorDTO) SetName(v string) *BatchModelErrorDTO {
	s.Name = &v
	return s
}

func (s *BatchModelErrorDTO) Validate() error {
	return dara.Validate(s)
}
