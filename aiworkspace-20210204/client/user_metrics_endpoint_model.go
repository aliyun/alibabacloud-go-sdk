// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserMetricsEndpoint interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *UserMetricsEndpoint
	GetPath() *string
	SetPort(v int32) *UserMetricsEndpoint
	GetPort() *int32
}

type UserMetricsEndpoint struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UserMetricsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UserMetricsEndpoint) GoString() string {
	return s.String()
}

func (s *UserMetricsEndpoint) GetPath() *string {
	return s.Path
}

func (s *UserMetricsEndpoint) GetPort() *int32 {
	return s.Port
}

func (s *UserMetricsEndpoint) SetPath(v string) *UserMetricsEndpoint {
	s.Path = &v
	return s
}

func (s *UserMetricsEndpoint) SetPort(v int32) *UserMetricsEndpoint {
	s.Port = &v
	return s
}

func (s *UserMetricsEndpoint) Validate() error {
	return dara.Validate(s)
}
