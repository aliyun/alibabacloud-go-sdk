// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectionStatus interface {
	dara.Model
	String() string
	GoString() string
	SetInstallation(v *Installation) *ConnectionStatus
	GetInstallation() *Installation
}

type ConnectionStatus struct {
	Installation *Installation `json:"installation,omitempty" xml:"installation,omitempty"`
}

func (s ConnectionStatus) String() string {
	return dara.Prettify(s)
}

func (s ConnectionStatus) GoString() string {
	return s.String()
}

func (s *ConnectionStatus) GetInstallation() *Installation {
	return s.Installation
}

func (s *ConnectionStatus) SetInstallation(v *Installation) *ConnectionStatus {
	s.Installation = v
	return s
}

func (s *ConnectionStatus) Validate() error {
	return dara.Validate(s)
}
