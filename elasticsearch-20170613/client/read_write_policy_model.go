// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWritePolicy interface {
	dara.Model
	String() string
	GoString() string
	SetAutoGeneratePk(v bool) *ReadWritePolicy
	GetAutoGeneratePk() *bool
	SetWriteHa(v bool) *ReadWritePolicy
	GetWriteHa() *bool
	SetWritePolicy(v string) *ReadWritePolicy
	GetWritePolicy() *string
}

type ReadWritePolicy struct {
	AutoGeneratePk *bool   `json:"autoGeneratePk,omitempty" xml:"autoGeneratePk,omitempty"`
	WriteHa        *bool   `json:"writeHa,omitempty" xml:"writeHa,omitempty"`
	WritePolicy    *string `json:"writePolicy,omitempty" xml:"writePolicy,omitempty"`
}

func (s ReadWritePolicy) String() string {
	return dara.Prettify(s)
}

func (s ReadWritePolicy) GoString() string {
	return s.String()
}

func (s *ReadWritePolicy) GetAutoGeneratePk() *bool {
	return s.AutoGeneratePk
}

func (s *ReadWritePolicy) GetWriteHa() *bool {
	return s.WriteHa
}

func (s *ReadWritePolicy) GetWritePolicy() *string {
	return s.WritePolicy
}

func (s *ReadWritePolicy) SetAutoGeneratePk(v bool) *ReadWritePolicy {
	s.AutoGeneratePk = &v
	return s
}

func (s *ReadWritePolicy) SetWriteHa(v bool) *ReadWritePolicy {
	s.WriteHa = &v
	return s
}

func (s *ReadWritePolicy) SetWritePolicy(v string) *ReadWritePolicy {
	s.WritePolicy = &v
	return s
}

func (s *ReadWritePolicy) Validate() error {
	return dara.Validate(s)
}
