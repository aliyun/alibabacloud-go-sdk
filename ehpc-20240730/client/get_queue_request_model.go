// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetQueueRequest
	GetClusterId() *string
	SetQueueName(v string) *GetQueueRequest
	GetQueueName() *string
}

type GetQueueRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queue name.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s GetQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRequest) GoString() string {
	return s.String()
}

func (s *GetQueueRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueRequest) SetClusterId(v string) *GetQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *GetQueueRequest) SetQueueName(v string) *GetQueueRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueRequest) Validate() error {
	return dara.Validate(s)
}
