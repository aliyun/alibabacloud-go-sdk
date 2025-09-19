// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetAlertStrategyRequest
	GetId() *int64
}

type GetAlertStrategyRequest struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *GetAlertStrategyRequest) SetId(v int64) *GetAlertStrategyRequest {
	s.Id = &v
	return s
}

func (s *GetAlertStrategyRequest) Validate() error {
	return dara.Validate(s)
}
