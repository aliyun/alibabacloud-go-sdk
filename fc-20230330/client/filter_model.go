// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilter interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v *Key) *Filter
	GetKey() *Key
}

type Filter struct {
	Key *Key `json:"key,omitempty" xml:"key,omitempty"`
}

func (s Filter) String() string {
	return dara.Prettify(s)
}

func (s Filter) GoString() string {
	return s.String()
}

func (s *Filter) GetKey() *Key {
	return s.Key
}

func (s *Filter) SetKey(v *Key) *Filter {
	s.Key = v
	return s
}

func (s *Filter) Validate() error {
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	return nil
}
