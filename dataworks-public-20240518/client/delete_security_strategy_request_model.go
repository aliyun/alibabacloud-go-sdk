// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteSecurityStrategyRequest
	GetId() *int64
}

type DeleteSecurityStrategyRequest struct {
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSecurityStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSecurityStrategyRequest) SetId(v int64) *DeleteSecurityStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteSecurityStrategyRequest) Validate() error {
	return dara.Validate(s)
}
