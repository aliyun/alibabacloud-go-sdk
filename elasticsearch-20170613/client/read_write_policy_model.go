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
	// Specifies whether to automatically generate a document hash primary key when no primary key exists. Valid values:
	//
	// - true (default): automatically generates a primary key.
	//
	// - false: does not automatically generate a primary key.
	//
	// 	Notice:  autoGeneratePk cannot be modified independently. The autoGeneratePk setting takes effect only when writeHa is changed from false to true.
	//
	// .
	//
	// example:
	//
	// true
	AutoGeneratePk *bool `json:"autoGeneratePk,omitempty" xml:"autoGeneratePk,omitempty"`
	// Specifies whether to enable the write high-availability feature. Valid values:
	//
	// - true: enabled.
	//
	// - false: not enabled.
	//
	// example:
	//
	// true
	WriteHa *bool `json:"writeHa,omitempty" xml:"writeHa,omitempty"`
	// Temporarily switches between synchronous and asynchronous high availability. Valid values:
	//
	// - sync: temporarily switches from asynchronous write high availability to synchronous write.
	//
	// - async: restores asynchronous write high availability after synchronous write is temporarily enabled.
	//
	// > This field takes effect only when high availability is enabled, that is, writeHa is set to true. You do not need to pass in the writeHa field when setting this field.
	//
	// example:
	//
	// sync
	WritePolicy *string `json:"writePolicy,omitempty" xml:"writePolicy,omitempty"`
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
