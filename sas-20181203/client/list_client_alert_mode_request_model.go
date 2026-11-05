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
	// Specifies the protection mode. Valid values:
	//
	// - **strict**: Strict mode. Defense mode has a certain risk of false positives and is recommended for use during critical event protection.
	//
	// - **balance**: Balance mode. Defense mode detects more suspicious risks while reducing false positives.
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
