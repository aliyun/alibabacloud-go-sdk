// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientAlertModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *UpdateClientAlertModeRequest
	GetMode() *string
	SetUuids(v []*string) *UpdateClientAlertModeRequest
	GetUuids() []*string
}

type UpdateClientAlertModeRequest struct {
	// The protection mode. Valid values:
	//
	// 	- **strict**: The strict mode. False positives may be generated. We recommend that you enable this mode during major events.
	//
	// 	- **balance**: The balanced mode. More risks can be detected with less false positives in this mode.
	//
	// example:
	//
	// balance
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The UUIDs of servers.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s UpdateClientAlertModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientAlertModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientAlertModeRequest) GetMode() *string {
	return s.Mode
}

func (s *UpdateClientAlertModeRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *UpdateClientAlertModeRequest) SetMode(v string) *UpdateClientAlertModeRequest {
	s.Mode = &v
	return s
}

func (s *UpdateClientAlertModeRequest) SetUuids(v []*string) *UpdateClientAlertModeRequest {
	s.Uuids = v
	return s
}

func (s *UpdateClientAlertModeRequest) Validate() error {
	return dara.Validate(s)
}
