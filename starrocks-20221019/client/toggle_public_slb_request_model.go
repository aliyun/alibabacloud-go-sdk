// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePublicSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePublicSlb(v bool) *TogglePublicSlbRequest
	GetEnablePublicSlb() *bool
	SetGatewayId(v string) *TogglePublicSlbRequest
	GetGatewayId() *string
	SetInstanceId(v string) *TogglePublicSlbRequest
	GetInstanceId() *string
}

type TogglePublicSlbRequest struct {
	// example:
	//
	// true
	EnablePublicSlb *bool `json:"EnablePublicSlb,omitempty" xml:"EnablePublicSlb,omitempty"`
	// example:
	//
	// gw-0002xci9buu68ongixvk
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TogglePublicSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s TogglePublicSlbRequest) GoString() string {
	return s.String()
}

func (s *TogglePublicSlbRequest) GetEnablePublicSlb() *bool {
	return s.EnablePublicSlb
}

func (s *TogglePublicSlbRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *TogglePublicSlbRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TogglePublicSlbRequest) SetEnablePublicSlb(v bool) *TogglePublicSlbRequest {
	s.EnablePublicSlb = &v
	return s
}

func (s *TogglePublicSlbRequest) SetGatewayId(v string) *TogglePublicSlbRequest {
	s.GatewayId = &v
	return s
}

func (s *TogglePublicSlbRequest) SetInstanceId(v string) *TogglePublicSlbRequest {
	s.InstanceId = &v
	return s
}

func (s *TogglePublicSlbRequest) Validate() error {
	return dara.Validate(s)
}
