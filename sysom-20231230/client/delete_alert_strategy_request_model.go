// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteAlertStrategyRequest
	GetId() *string
}

type DeleteAlertStrategyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertStrategyRequest) GetId() *string {
	return s.Id
}

func (s *DeleteAlertStrategyRequest) SetId(v string) *DeleteAlertStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertStrategyRequest) Validate() error {
	return dara.Validate(s)
}
