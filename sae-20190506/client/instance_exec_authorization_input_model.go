// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceExecAuthorizationInput interface {
	dara.Model
	String() string
	GoString() string
	SetOptions(v *InstanceExecAuthorizationInputOptions) *InstanceExecAuthorizationInput
	GetOptions() *InstanceExecAuthorizationInputOptions
}

type InstanceExecAuthorizationInput struct {
	Options *InstanceExecAuthorizationInputOptions `json:"options,omitempty" xml:"options,omitempty"`
}

func (s InstanceExecAuthorizationInput) String() string {
	return dara.Prettify(s)
}

func (s InstanceExecAuthorizationInput) GoString() string {
	return s.String()
}

func (s *InstanceExecAuthorizationInput) GetOptions() *InstanceExecAuthorizationInputOptions {
	return s.Options
}

func (s *InstanceExecAuthorizationInput) SetOptions(v *InstanceExecAuthorizationInputOptions) *InstanceExecAuthorizationInput {
	s.Options = v
	return s
}

func (s *InstanceExecAuthorizationInput) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}
