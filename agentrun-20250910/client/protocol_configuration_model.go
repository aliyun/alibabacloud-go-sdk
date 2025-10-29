// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProtocolConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ProtocolConfiguration
	GetType() *string
}

type ProtocolConfiguration struct {
	// example:
	//
	// HTTP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ProtocolConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ProtocolConfiguration) GoString() string {
	return s.String()
}

func (s *ProtocolConfiguration) GetType() *string {
	return s.Type
}

func (s *ProtocolConfiguration) SetType(v string) *ProtocolConfiguration {
	s.Type = &v
	return s
}

func (s *ProtocolConfiguration) Validate() error {
	return dara.Validate(s)
}
