// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsolateLeaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *IsolateLeaderRequest
	GetInstanceId() *string
	SetIsolateLeader(v bool) *IsolateLeaderRequest
	GetIsolateLeader() *bool
}

type IsolateLeaderRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsolateLeader *bool `json:"IsolateLeader,omitempty" xml:"IsolateLeader,omitempty"`
}

func (s IsolateLeaderRequest) String() string {
	return dara.Prettify(s)
}

func (s IsolateLeaderRequest) GoString() string {
	return s.String()
}

func (s *IsolateLeaderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *IsolateLeaderRequest) GetIsolateLeader() *bool {
	return s.IsolateLeader
}

func (s *IsolateLeaderRequest) SetInstanceId(v string) *IsolateLeaderRequest {
	s.InstanceId = &v
	return s
}

func (s *IsolateLeaderRequest) SetIsolateLeader(v bool) *IsolateLeaderRequest {
	s.IsolateLeader = &v
	return s
}

func (s *IsolateLeaderRequest) Validate() error {
	return dara.Validate(s)
}
