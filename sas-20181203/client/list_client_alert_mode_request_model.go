// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientAlertModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *ListClientAlertModeRequest
	GetMode() *string
}

type ListClientAlertModeRequest struct {
	// The protection mode. Valid values:
	//
	// 	- **strict**: The strict mode. False positives may be generated. We recommend that you enable this mode during major events.
	//
	// 	- **balance**: The balanced mode. More risks can be detected with less false positives in this mode.
	//
	// example:
	//
	// strict
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ListClientAlertModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientAlertModeRequest) GoString() string {
	return s.String()
}

func (s *ListClientAlertModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ListClientAlertModeRequest) SetMode(v string) *ListClientAlertModeRequest {
	s.Mode = &v
	return s
}

func (s *ListClientAlertModeRequest) Validate() error {
	return dara.Validate(s)
}
