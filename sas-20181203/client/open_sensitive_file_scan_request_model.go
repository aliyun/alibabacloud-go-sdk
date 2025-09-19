// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSensitiveFileScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSwitchOn(v string) *OpenSensitiveFileScanRequest
	GetSwitchOn() *string
}

type OpenSensitiveFileScanRequest struct {
	// Specifies whether to enable or disable sensitive file scan. Valid values:
	//
	// 	- **on**: enables sensitive file scan
	//
	// 	- **off**: disables sensitive file scan
	//
	// example:
	//
	// on
	SwitchOn *string `json:"SwitchOn,omitempty" xml:"SwitchOn,omitempty"`
}

func (s OpenSensitiveFileScanRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenSensitiveFileScanRequest) GoString() string {
	return s.String()
}

func (s *OpenSensitiveFileScanRequest) GetSwitchOn() *string {
	return s.SwitchOn
}

func (s *OpenSensitiveFileScanRequest) SetSwitchOn(v string) *OpenSensitiveFileScanRequest {
	s.SwitchOn = &v
	return s
}

func (s *OpenSensitiveFileScanRequest) Validate() error {
	return dara.Validate(s)
}
