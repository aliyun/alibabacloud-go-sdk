// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListQueuesShrinkRequest
	GetClusterId() *string
	SetQueueNamesShrink(v string) *ListQueuesShrinkRequest
	GetQueueNamesShrink() *string
}

type ListQueuesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The names of the queues that you want to query. You can specify up to eight names.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
}

func (s ListQueuesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListQueuesShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *ListQueuesShrinkRequest) SetClusterId(v string) *ListQueuesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListQueuesShrinkRequest) SetQueueNamesShrink(v string) *ListQueuesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *ListQueuesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
