// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorIdsShrink(v string) *DeleteJobsShrinkRequest
	GetExecutorIdsShrink() *string
	SetJobScheduler(v string) *DeleteJobsShrinkRequest
	GetJobScheduler() *string
	SetJobSpecShrink(v string) *DeleteJobsShrinkRequest
	GetJobSpecShrink() *string
}

type DeleteJobsShrinkRequest struct {
	ExecutorIdsShrink *string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty"`
	JobScheduler      *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	JobSpecShrink     *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
}

func (s DeleteJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsShrinkRequest) GetExecutorIdsShrink() *string {
	return s.ExecutorIdsShrink
}

func (s *DeleteJobsShrinkRequest) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *DeleteJobsShrinkRequest) GetJobSpecShrink() *string {
	return s.JobSpecShrink
}

func (s *DeleteJobsShrinkRequest) SetExecutorIdsShrink(v string) *DeleteJobsShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobScheduler(v string) *DeleteJobsShrinkRequest {
	s.JobScheduler = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobSpecShrink(v string) *DeleteJobsShrinkRequest {
	s.JobSpecShrink = &v
	return s
}

func (s *DeleteJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
