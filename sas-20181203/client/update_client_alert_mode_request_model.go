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
	// Specifies the protection mode. Valid values:
	//
	// - **strict**: Strict mode. Defense mode has a certain risk of false positives and is recommended during critical event protection.
	//
	// - **balance**: Balanced mode. Defense mode detects more suspicious risks while reducing false positives.
	//
	// example:
	//
	// balance
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The list of server UUIDs.
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
