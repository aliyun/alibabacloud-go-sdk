// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionReleaseStageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *CancelExecutionReleaseStageResponseBody
	GetSuccess() *bool
}

type CancelExecutionReleaseStageResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelExecutionReleaseStageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionReleaseStageResponseBody) GoString() string {
	return s.String()
}

func (s *CancelExecutionReleaseStageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelExecutionReleaseStageResponseBody) SetSuccess(v bool) *CancelExecutionReleaseStageResponseBody {
	s.Success = &v
	return s
}

func (s *CancelExecutionReleaseStageResponseBody) Validate() error {
	return dara.Validate(s)
}
