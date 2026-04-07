// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealtimeInstanceStatsRequest
	GetInstanceId() *string
}

type GetRealtimeInstanceStatsRequest struct {
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealtimeInstanceStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeInstanceStatsRequest) SetInstanceId(v string) *GetRealtimeInstanceStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeInstanceStatsRequest) Validate() error {
	return dara.Validate(s)
}
