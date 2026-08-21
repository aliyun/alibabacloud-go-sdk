// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForwardStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardId(v string) *GetForwardStrategyRequest
	GetForwardId() *string
}

type GetForwardStrategyRequest struct {
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fs-651b975a22aa019c
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
}

func (s GetForwardStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetForwardStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetForwardStrategyRequest) GetForwardId() *string {
	return s.ForwardId
}

func (s *GetForwardStrategyRequest) SetForwardId(v string) *GetForwardStrategyRequest {
	s.ForwardId = &v
	return s
}

func (s *GetForwardStrategyRequest) Validate() error {
	return dara.Validate(s)
}
