// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestinationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetOnFailure(v *Destination) *DestinationConfig
	GetOnFailure() *Destination
	SetOnSuccess(v *Destination) *DestinationConfig
	GetOnSuccess() *Destination
}

type DestinationConfig struct {
	OnFailure *Destination `json:"onFailure,omitempty" xml:"onFailure,omitempty"`
	OnSuccess *Destination `json:"onSuccess,omitempty" xml:"onSuccess,omitempty"`
}

func (s DestinationConfig) String() string {
	return dara.Prettify(s)
}

func (s DestinationConfig) GoString() string {
	return s.String()
}

func (s *DestinationConfig) GetOnFailure() *Destination {
	return s.OnFailure
}

func (s *DestinationConfig) GetOnSuccess() *Destination {
	return s.OnSuccess
}

func (s *DestinationConfig) SetOnFailure(v *Destination) *DestinationConfig {
	s.OnFailure = v
	return s
}

func (s *DestinationConfig) SetOnSuccess(v *Destination) *DestinationConfig {
	s.OnSuccess = v
	return s
}

func (s *DestinationConfig) Validate() error {
	return dara.Validate(s)
}
