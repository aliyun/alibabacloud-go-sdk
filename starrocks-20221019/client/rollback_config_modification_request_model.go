// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackConfigModificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigHistoryId(v int64) *RollbackConfigModificationRequest
	GetConfigHistoryId() *int64
	SetInstanceId(v string) *RollbackConfigModificationRequest
	GetInstanceId() *string
	SetRestart(v bool) *RollbackConfigModificationRequest
	GetRestart() *bool
}

type RollbackConfigModificationRequest struct {
	// example:
	//
	// 187291
	ConfigHistoryId *int64 `json:"ConfigHistoryId,omitempty" xml:"ConfigHistoryId,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s RollbackConfigModificationRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackConfigModificationRequest) GoString() string {
	return s.String()
}

func (s *RollbackConfigModificationRequest) GetConfigHistoryId() *int64 {
	return s.ConfigHistoryId
}

func (s *RollbackConfigModificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RollbackConfigModificationRequest) GetRestart() *bool {
	return s.Restart
}

func (s *RollbackConfigModificationRequest) SetConfigHistoryId(v int64) *RollbackConfigModificationRequest {
	s.ConfigHistoryId = &v
	return s
}

func (s *RollbackConfigModificationRequest) SetInstanceId(v string) *RollbackConfigModificationRequest {
	s.InstanceId = &v
	return s
}

func (s *RollbackConfigModificationRequest) SetRestart(v bool) *RollbackConfigModificationRequest {
	s.Restart = &v
	return s
}

func (s *RollbackConfigModificationRequest) Validate() error {
	return dara.Validate(s)
}
