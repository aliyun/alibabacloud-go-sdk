// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllTrafficSpecialControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DeleteAllTrafficSpecialControlRequest
	GetSecurityToken() *string
	SetTrafficControlId(v string) *DeleteAllTrafficSpecialControlRequest
	GetTrafficControlId() *string
}

type DeleteAllTrafficSpecialControlRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// fa876ffb-caab-4f0a-93b3-3409f2fa5199
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

func (s DeleteAllTrafficSpecialControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAllTrafficSpecialControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DeleteAllTrafficSpecialControlRequest) SetSecurityToken(v string) *DeleteAllTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlRequest) SetTrafficControlId(v string) *DeleteAllTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlRequest) Validate() error {
	return dara.Validate(s)
}
