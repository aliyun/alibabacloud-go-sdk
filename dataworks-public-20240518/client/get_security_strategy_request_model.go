// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetSecurityStrategyRequest
	GetId() *int64
}

type GetSecurityStrategyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSecurityStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityStrategyRequest) SetId(v int64) *GetSecurityStrategyRequest {
	s.Id = &v
	return s
}

func (s *GetSecurityStrategyRequest) Validate() error {
	return dara.Validate(s)
}
