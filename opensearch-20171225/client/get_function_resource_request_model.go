// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *GetFunctionResourceRequest
	GetOutput() *string
}

type GetFunctionResourceRequest struct {
	// The output level.
	//
	// Valid values:
	//
	// 	- simple
	//
	// 	- normal
	//
	// 	- detail
	//
	// example:
	//
	// detail
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
}

func (s GetFunctionResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceRequest) GetOutput() *string {
	return s.Output
}

func (s *GetFunctionResourceRequest) SetOutput(v string) *GetFunctionResourceRequest {
	s.Output = &v
	return s
}

func (s *GetFunctionResourceRequest) Validate() error {
	return dara.Validate(s)
}
