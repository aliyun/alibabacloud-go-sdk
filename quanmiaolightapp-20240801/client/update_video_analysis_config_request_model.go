// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncConcurrency(v int32) *UpdateVideoAnalysisConfigRequest
	GetAsyncConcurrency() *int32
}

type UpdateVideoAnalysisConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s UpdateVideoAnalysisConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigRequest) GetAsyncConcurrency() *int32 {
	return s.AsyncConcurrency
}

func (s *UpdateVideoAnalysisConfigRequest) SetAsyncConcurrency(v int32) *UpdateVideoAnalysisConfigRequest {
	s.AsyncConcurrency = &v
	return s
}

func (s *UpdateVideoAnalysisConfigRequest) Validate() error {
	return dara.Validate(s)
}
