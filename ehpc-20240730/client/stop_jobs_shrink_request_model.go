// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StopJobsShrinkRequest
	GetClusterId() *string
	SetJobIdsShrink(v string) *StopJobsShrinkRequest
	GetJobIdsShrink() *string
}

type StopJobsShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the jobs that you want to stop.
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s StopJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopJobsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StopJobsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *StopJobsShrinkRequest) SetClusterId(v string) *StopJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsShrinkRequest) SetJobIdsShrink(v string) *StopJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *StopJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
