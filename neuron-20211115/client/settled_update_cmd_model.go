// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSettledUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *SettledUpdateCmd
	GetStatus() *string
}

type SettledUpdateCmd struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SettledUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s SettledUpdateCmd) GoString() string {
	return s.String()
}

func (s *SettledUpdateCmd) GetStatus() *string {
	return s.Status
}

func (s *SettledUpdateCmd) SetStatus(v string) *SettledUpdateCmd {
	s.Status = &v
	return s
}

func (s *SettledUpdateCmd) Validate() error {
	return dara.Validate(s)
}
