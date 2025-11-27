// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIdsShrink(v string) *DeleteJobRecordsShrinkRequest
	GetJobIdsShrink() *string
}

type DeleteJobRecordsShrinkRequest struct {
	// The list of job IDs.
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s DeleteJobRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRecordsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *DeleteJobRecordsShrinkRequest) SetJobIdsShrink(v string) *DeleteJobRecordsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *DeleteJobRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
