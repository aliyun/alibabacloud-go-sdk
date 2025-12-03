// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefuseReleaseStagePipelineValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *RefuseReleaseStagePipelineValidateResponseBody
	GetSuccess() *bool
}

type RefuseReleaseStagePipelineValidateResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefuseReleaseStagePipelineValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefuseReleaseStagePipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseReleaseStagePipelineValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefuseReleaseStagePipelineValidateResponseBody) SetSuccess(v bool) *RefuseReleaseStagePipelineValidateResponseBody {
	s.Success = &v
	return s
}

func (s *RefuseReleaseStagePipelineValidateResponseBody) Validate() error {
	return dara.Validate(s)
}
