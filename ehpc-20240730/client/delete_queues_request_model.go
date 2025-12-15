// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteQueuesRequest
	GetClusterId() *string
	SetQueueNames(v []*string) *DeleteQueuesRequest
	GetQueueNames() []*string
}

type DeleteQueuesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queues that you want to delete.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
}

func (s DeleteQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueuesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueuesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteQueuesRequest) GetQueueNames() []*string {
	return s.QueueNames
}

func (s *DeleteQueuesRequest) SetClusterId(v string) *DeleteQueuesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueuesRequest) SetQueueNames(v []*string) *DeleteQueuesRequest {
	s.QueueNames = v
	return s
}

func (s *DeleteQueuesRequest) Validate() error {
	return dara.Validate(s)
}
