// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPassReleaseStagePipelineValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *PassReleaseStagePipelineValidateResponseBody
	GetSuccess() *bool
}

type PassReleaseStagePipelineValidateResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PassReleaseStagePipelineValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PassReleaseStagePipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *PassReleaseStagePipelineValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PassReleaseStagePipelineValidateResponseBody) SetSuccess(v bool) *PassReleaseStagePipelineValidateResponseBody {
	s.Success = &v
	return s
}

func (s *PassReleaseStagePipelineValidateResponseBody) Validate() error {
	return dara.Validate(s)
}
