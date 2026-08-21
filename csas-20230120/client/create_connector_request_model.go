// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *CreateConnectorRequest
	GetBandwidth() *int32
	SetName(v string) *CreateConnectorRequest
	GetName() *string
	SetRegion(v string) *CreateConnectorRequest
	GetRegion() *string
	SetSwitchStatus(v string) *CreateConnectorRequest
	GetSwitchStatus() *string
}

type CreateConnectorRequest struct {
	// The bandwidth value (Mbit/s).
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The connector name. The name must be 1 to 128 characters in length and can contain letters, digits, Chinese characters, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// auto-dr-connector-cq-dl3e4j
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The connector instance status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Shutdown.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s CreateConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectorRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateConnectorRequest) GetName() *string {
	return s.Name
}

func (s *CreateConnectorRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateConnectorRequest) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *CreateConnectorRequest) SetBandwidth(v int32) *CreateConnectorRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateConnectorRequest) SetName(v string) *CreateConnectorRequest {
	s.Name = &v
	return s
}

func (s *CreateConnectorRequest) SetRegion(v string) *CreateConnectorRequest {
	s.Region = &v
	return s
}

func (s *CreateConnectorRequest) SetSwitchStatus(v string) *CreateConnectorRequest {
	s.SwitchStatus = &v
	return s
}

func (s *CreateConnectorRequest) Validate() error {
	return dara.Validate(s)
}
