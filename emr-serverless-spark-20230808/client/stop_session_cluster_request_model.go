// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueueName(v string) *StopSessionClusterRequest
	GetQueueName() *string
	SetSessionClusterId(v string) *StopSessionClusterRequest
	GetSessionClusterId() *string
	SetRegionId(v string) *StopSessionClusterRequest
	GetRegionId() *string
}

type StopSessionClusterRequest struct {
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The session ID.
	//
	// example:
	//
	// sc-xxxxxxxxxxxx
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *StopSessionClusterRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *StopSessionClusterRequest) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *StopSessionClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopSessionClusterRequest) SetQueueName(v string) *StopSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *StopSessionClusterRequest) SetSessionClusterId(v string) *StopSessionClusterRequest {
	s.SessionClusterId = &v
	return s
}

func (s *StopSessionClusterRequest) SetRegionId(v string) *StopSessionClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StopSessionClusterRequest) Validate() error {
	return dara.Validate(s)
}
