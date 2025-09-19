// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionTargetDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetId(v int64) *GetInterceptionTargetDetailRequest
	GetTargetId() *int64
}

type GetInterceptionTargetDetailRequest struct {
	// The ID of the network object.
	//
	// > You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the IDs of network objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// 402008
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s GetInterceptionTargetDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionTargetDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInterceptionTargetDetailRequest) GetTargetId() *int64 {
	return s.TargetId
}

func (s *GetInterceptionTargetDetailRequest) SetTargetId(v int64) *GetInterceptionTargetDetailRequest {
	s.TargetId = &v
	return s
}

func (s *GetInterceptionTargetDetailRequest) Validate() error {
	return dara.Validate(s)
}
