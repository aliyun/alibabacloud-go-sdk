// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *BatchGetYikeAIAppJobRequest
	GetJobIds() *string
}

type BatchGetYikeAIAppJobRequest struct {
	// example:
	//
	// cec2c13da767c090,ca3d3c9737f04586
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s BatchGetYikeAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *BatchGetYikeAIAppJobRequest) SetJobIds(v string) *BatchGetYikeAIAppJobRequest {
	s.JobIds = &v
	return s
}

func (s *BatchGetYikeAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
