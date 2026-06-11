// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexJsonKey interface {
	dara.Model
	String() string
	GoString() string
	SetChn(v bool) *IndexJsonKey
	GetChn() *bool
	SetType(v string) *IndexJsonKey
	GetType() *string
}

type IndexJsonKey struct {
	// Specifies whether the field contains Chinese characters.
	//
	// - `true`: The field contains Chinese characters.
	//
	// - `false`: The field does not contain Chinese characters.
	//
	// example:
	//
	// true
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The field type.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IndexJsonKey) String() string {
	return dara.Prettify(s)
}

func (s IndexJsonKey) GoString() string {
	return s.String()
}

func (s *IndexJsonKey) GetChn() *bool {
	return s.Chn
}

func (s *IndexJsonKey) GetType() *string {
	return s.Type
}

func (s *IndexJsonKey) SetChn(v bool) *IndexJsonKey {
	s.Chn = &v
	return s
}

func (s *IndexJsonKey) SetType(v string) *IndexJsonKey {
	s.Type = &v
	return s
}

func (s *IndexJsonKey) Validate() error {
	return dara.Validate(s)
}
