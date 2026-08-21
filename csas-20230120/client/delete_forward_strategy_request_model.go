// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardId(v string) *DeleteForwardStrategyRequest
	GetForwardId() *string
}

type DeleteForwardStrategyRequest struct {
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fs-3fb9b5ae28ee5416
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
}

func (s DeleteForwardStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardStrategyRequest) GetForwardId() *string {
	return s.ForwardId
}

func (s *DeleteForwardStrategyRequest) SetForwardId(v string) *DeleteForwardStrategyRequest {
	s.ForwardId = &v
	return s
}

func (s *DeleteForwardStrategyRequest) Validate() error {
	return dara.Validate(s)
}
