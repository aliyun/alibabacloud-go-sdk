// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKey interface {
	dara.Model
	String() string
	GoString() string
	SetPrefix(v string) *Key
	GetPrefix() *string
	SetSuffix(v string) *Key
	GetSuffix() *string
}

type Key struct {
	// example:
	//
	// serverless_
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// example:
	//
	// .zip
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
}

func (s Key) String() string {
	return dara.Prettify(s)
}

func (s Key) GoString() string {
	return s.String()
}

func (s *Key) GetPrefix() *string {
	return s.Prefix
}

func (s *Key) GetSuffix() *string {
	return s.Suffix
}

func (s *Key) SetPrefix(v string) *Key {
	s.Prefix = &v
	return s
}

func (s *Key) SetSuffix(v string) *Key {
	s.Suffix = &v
	return s
}

func (s *Key) Validate() error {
	return dara.Validate(s)
}
