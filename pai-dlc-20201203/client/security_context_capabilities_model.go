// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecurityContextCapabilities interface {
	dara.Model
	String() string
	GoString() string
	SetAdd(v []*string) *SecurityContextCapabilities
	GetAdd() []*string
	SetDrop(v []*string) *SecurityContextCapabilities
	GetDrop() []*string
}

type SecurityContextCapabilities struct {
	Add  []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
	Drop []*string `json:"Drop,omitempty" xml:"Drop,omitempty" type:"Repeated"`
}

func (s SecurityContextCapabilities) String() string {
	return dara.Prettify(s)
}

func (s SecurityContextCapabilities) GoString() string {
	return s.String()
}

func (s *SecurityContextCapabilities) GetAdd() []*string {
	return s.Add
}

func (s *SecurityContextCapabilities) GetDrop() []*string {
	return s.Drop
}

func (s *SecurityContextCapabilities) SetAdd(v []*string) *SecurityContextCapabilities {
	s.Add = v
	return s
}

func (s *SecurityContextCapabilities) SetDrop(v []*string) *SecurityContextCapabilities {
	s.Drop = v
	return s
}

func (s *SecurityContextCapabilities) Validate() error {
	return dara.Validate(s)
}
