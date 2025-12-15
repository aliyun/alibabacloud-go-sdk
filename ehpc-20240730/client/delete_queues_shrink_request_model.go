// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueuesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteQueuesShrinkRequest
	GetClusterId() *string
	SetQueueNamesShrink(v string) *DeleteQueuesShrinkRequest
	GetQueueNamesShrink() *string
}

type DeleteQueuesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queues that you want to delete.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
}

func (s DeleteQueuesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueuesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueuesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteQueuesShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *DeleteQueuesShrinkRequest) SetClusterId(v string) *DeleteQueuesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueuesShrinkRequest) SetQueueNamesShrink(v string) *DeleteQueuesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *DeleteQueuesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
