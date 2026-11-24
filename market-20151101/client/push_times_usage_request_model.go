// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTimesUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjust(v bool) *PushTimesUsageRequest
	GetAdjust() *bool
	SetAdjustDate(v int64) *PushTimesUsageRequest
	GetAdjustDate() *int64
	SetClientToken(v string) *PushTimesUsageRequest
	GetClientToken() *string
	SetInstanceId(v string) *PushTimesUsageRequest
	GetInstanceId() *string
	SetMark(v string) *PushTimesUsageRequest
	GetMark() *string
	SetTimes(v int64) *PushTimesUsageRequest
	GetTimes() *int64
}

type PushTimesUsageRequest struct {
	Adjust     *bool  `json:"Adjust,omitempty" xml:"Adjust,omitempty"`
	AdjustDate *int64 `json:"AdjustDate,omitempty" xml:"AdjustDate,omitempty"`
	// The client token used to ensure request idempotence. We recommend that you use a UUID. The same `ClientToken` can be submitted only once within 24 hours. Retries for the same operation must use the same `ClientToken`, and new operations must use a new `ClientToken`.
	//
	// example:
	//
	// 6dff6c70-3484-4a39-b725-164e3ad9b20d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The Alibaba Cloud Marketplace instance ID.
	//
	// example:
	//
	// 1000001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mark       *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The number of uses to deduct. The value must be greater than 0.
	//
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

func (s *PushTimesUsageRequest) GetAdjust() *bool {
	return s.Adjust
}

func (s *PushTimesUsageRequest) GetAdjustDate() *int64 {
	return s.AdjustDate
}

func (s *PushTimesUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PushTimesUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushTimesUsageRequest) GetMark() *string {
	return s.Mark
}

func (s *PushTimesUsageRequest) GetTimes() *int64 {
	return s.Times
}

func (s *PushTimesUsageRequest) SetAdjust(v bool) *PushTimesUsageRequest {
	s.Adjust = &v
	return s
}

func (s *PushTimesUsageRequest) SetAdjustDate(v int64) *PushTimesUsageRequest {
	s.AdjustDate = &v
	return s
}

func (s *PushTimesUsageRequest) SetClientToken(v string) *PushTimesUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *PushTimesUsageRequest) SetInstanceId(v string) *PushTimesUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *PushTimesUsageRequest) SetMark(v string) *PushTimesUsageRequest {
	s.Mark = &v
	return s
}

func (s *PushTimesUsageRequest) SetTimes(v int64) *PushTimesUsageRequest {
	s.Times = &v
	return s
}

func (s *PushTimesUsageRequest) Validate() error {
	return dara.Validate(s)
}
