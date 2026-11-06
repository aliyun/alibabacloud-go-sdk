// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskStepRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ListTaskStepRequest
	GetType() *string
}

type ListTaskStepRequest struct {
	// The task type to which the step belongs. Valid values:
	//
	// - **ADAPTOR**: application adaptation steps
	//
	// - **IMAGE**: image creation steps
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTaskStepRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskStepRequest) GoString() string {
	return s.String()
}

func (s *ListTaskStepRequest) GetType() *string {
	return s.Type
}

func (s *ListTaskStepRequest) SetType(v string) *ListTaskStepRequest {
	s.Type = &v
	return s
}

func (s *ListTaskStepRequest) Validate() error {
	return dara.Validate(s)
}
