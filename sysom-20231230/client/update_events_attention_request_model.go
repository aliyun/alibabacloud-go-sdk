// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventsAttentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v int32) *UpdateEventsAttentionRequest
	GetMode() *int32
	SetRange(v string) *UpdateEventsAttentionRequest
	GetRange() *string
	SetUuid(v string) *UpdateEventsAttentionRequest
	GetUuid() *string
}

type UpdateEventsAttentionRequest struct {
	// The sensitivity of the anomaly event. Valid values: -1 to 3. A value of -1 indicates that the sensitivity is decreased by 1. A value of 0 indicates that the sensitivity is increased by 1.
	//
	// example:
	//
	// -1
	Mode *int32 `json:"mode,omitempty" xml:"mode,omitempty"`
	// The scope in which the update takes effect. Valid values: cluster and node.
	//
	// example:
	//
	// cluster
	Range *string `json:"range,omitempty" xml:"range,omitempty"`
	// The UUID of the anomaly event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03de78af-f49f-433d-b5b1-0f6a70c493ba
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s UpdateEventsAttentionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventsAttentionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionRequest) GetMode() *int32 {
	return s.Mode
}

func (s *UpdateEventsAttentionRequest) GetRange() *string {
	return s.Range
}

func (s *UpdateEventsAttentionRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateEventsAttentionRequest) SetMode(v int32) *UpdateEventsAttentionRequest {
	s.Mode = &v
	return s
}

func (s *UpdateEventsAttentionRequest) SetRange(v string) *UpdateEventsAttentionRequest {
	s.Range = &v
	return s
}

func (s *UpdateEventsAttentionRequest) SetUuid(v string) *UpdateEventsAttentionRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateEventsAttentionRequest) Validate() error {
	return dara.Validate(s)
}
