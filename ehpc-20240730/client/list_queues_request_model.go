// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListQueuesRequest
	GetClusterId() *string
	SetQueueNames(v []*string) *ListQueuesRequest
	GetQueueNames() []*string
}

type ListQueuesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The names of the queues that you want to query. You can specify up to eight names.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
}

func (s ListQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListQueuesRequest) GetQueueNames() []*string {
	return s.QueueNames
}

func (s *ListQueuesRequest) SetClusterId(v string) *ListQueuesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListQueuesRequest) SetQueueNames(v []*string) *ListQueuesRequest {
	s.QueueNames = v
	return s
}

func (s *ListQueuesRequest) Validate() error {
	return dara.Validate(s)
}
