// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListGatewayRequest
	GetInstanceId() *string
}

type ListGatewayRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGatewayRequest) SetInstanceId(v string) *ListGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayRequest) Validate() error {
	return dara.Validate(s)
}
