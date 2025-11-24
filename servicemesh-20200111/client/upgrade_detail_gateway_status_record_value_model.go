// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDetailGatewayStatusRecordValue interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpgradeDetailGatewayStatusRecordValue
	GetStatus() *string
	SetMessage(v string) *UpgradeDetailGatewayStatusRecordValue
	GetMessage() *string
	SetVersion(v string) *UpgradeDetailGatewayStatusRecordValue
	GetVersion() *string
}

type UpgradeDetailGatewayStatusRecordValue struct {
	// The upgrade status of the ingress gateway. Valid values:
	//
	// 	- `upgrading`: The ingress gateway is being upgraded.
	//
	// 	- `pending`: The ingress gateway waits to be upgraded.
	//
	// 	- `finished`: The ingress gateway upgrade is complete.
	//
	// 	- `notStart`: The ingress gateway upgrade does not start.
	//
	// 	- `failed`: The ingress gateway upgrade fails.
	//
	// 	- `unknown`: The upgrade status of the ingress gateway is unknown.
	//
	// example:
	//
	// upgrading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Additional status information of the ingress gateway.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The version of the ingress gateway.
	//
	// example:
	//
	// 1.9.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpgradeDetailGatewayStatusRecordValue) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDetailGatewayStatusRecordValue) GoString() string {
	return s.String()
}

func (s *UpgradeDetailGatewayStatusRecordValue) GetStatus() *string {
	return s.Status
}

func (s *UpgradeDetailGatewayStatusRecordValue) GetMessage() *string {
	return s.Message
}

func (s *UpgradeDetailGatewayStatusRecordValue) GetVersion() *string {
	return s.Version
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetStatus(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Status = &v
	return s
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetMessage(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Message = &v
	return s
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetVersion(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Version = &v
	return s
}

func (s *UpgradeDetailGatewayStatusRecordValue) Validate() error {
	return dara.Validate(s)
}
