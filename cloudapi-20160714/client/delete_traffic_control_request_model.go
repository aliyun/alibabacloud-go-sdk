// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DeleteTrafficControlRequest
	GetSecurityToken() *string
	SetTrafficControlId(v string) *DeleteTrafficControlRequest
	GetTrafficControlId() *string
}

type DeleteTrafficControlRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DeleteTrafficControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteTrafficControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DeleteTrafficControlRequest) SetSecurityToken(v string) *DeleteTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteTrafficControlRequest) SetTrafficControlId(v string) *DeleteTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DeleteTrafficControlRequest) Validate() error {
	return dara.Validate(s)
}
