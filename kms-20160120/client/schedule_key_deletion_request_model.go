// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleKeyDeletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *ScheduleKeyDeletionRequest
	GetKeyId() *string
	SetPendingWindowInDays(v int32) *ScheduleKeyDeletionRequest
	GetPendingWindowInDays() *int32
}

type ScheduleKeyDeletionRequest struct {
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The scheduled period after which the CMK is deleted. During this period, the CMK is in the PendingDeletion state. After this period ends, you cannot cancel the key deletion task.
	//
	// Valid values: 7 to 366.
	//
	// Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PendingWindowInDays *int32 `json:"PendingWindowInDays,omitempty" xml:"PendingWindowInDays,omitempty"`
}

func (s ScheduleKeyDeletionRequest) String() string {
	return dara.Prettify(s)
}

func (s ScheduleKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ScheduleKeyDeletionRequest) GetPendingWindowInDays() *int32 {
	return s.PendingWindowInDays
}

func (s *ScheduleKeyDeletionRequest) SetKeyId(v string) *ScheduleKeyDeletionRequest {
	s.KeyId = &v
	return s
}

func (s *ScheduleKeyDeletionRequest) SetPendingWindowInDays(v int32) *ScheduleKeyDeletionRequest {
	s.PendingWindowInDays = &v
	return s
}

func (s *ScheduleKeyDeletionRequest) Validate() error {
	return dara.Validate(s)
}
