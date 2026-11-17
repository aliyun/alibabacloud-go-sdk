// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueueName(v string) *StartSessionClusterRequest
	GetQueueName() *string
	SetSessionClusterId(v string) *StartSessionClusterRequest
	GetSessionClusterId() *string
	SetRegionId(v string) *StartSessionClusterRequest
	GetRegionId() *string
}

type StartSessionClusterRequest struct {
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
	// sc-xxxxxxxxxxx
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *StartSessionClusterRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *StartSessionClusterRequest) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *StartSessionClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartSessionClusterRequest) SetQueueName(v string) *StartSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *StartSessionClusterRequest) SetSessionClusterId(v string) *StartSessionClusterRequest {
	s.SessionClusterId = &v
	return s
}

func (s *StartSessionClusterRequest) SetRegionId(v string) *StartSessionClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StartSessionClusterRequest) Validate() error {
	return dara.Validate(s)
}
