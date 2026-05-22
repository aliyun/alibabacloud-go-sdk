// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTimesUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PushTimesUsageRequest
	GetClientToken() *string
	SetInstanceId(v string) *PushTimesUsageRequest
	GetInstanceId() *string
	SetTimes(v int64) *PushTimesUsageRequest
	GetTimes() *int64
}

type PushTimesUsageRequest struct {
	// example:
	//
	// 6dff6c70-3484-4a39-b725-164e3ad9b20d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1000001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	Times *int64 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PushTimesUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s PushTimesUsageRequest) GoString() string {
	return s.String()
}

func (s *PushTimesUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PushTimesUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushTimesUsageRequest) GetTimes() *int64 {
	return s.Times
}

func (s *PushTimesUsageRequest) SetClientToken(v string) *PushTimesUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *PushTimesUsageRequest) SetInstanceId(v string) *PushTimesUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *PushTimesUsageRequest) SetTimes(v int64) *PushTimesUsageRequest {
	s.Times = &v
	return s
}

func (s *PushTimesUsageRequest) Validate() error {
	return dara.Validate(s)
}
