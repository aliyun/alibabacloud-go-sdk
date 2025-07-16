// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentifier interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *Identifier
	GetDatabase() *string
	SetObject(v string) *Identifier
	GetObject() *string
}

type Identifier struct {
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	Object   *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s Identifier) String() string {
	return dara.Prettify(s)
}

func (s Identifier) GoString() string {
	return s.String()
}

func (s *Identifier) GetDatabase() *string {
	return s.Database
}

func (s *Identifier) GetObject() *string {
	return s.Object
}

func (s *Identifier) SetDatabase(v string) *Identifier {
	s.Database = &v
	return s
}

func (s *Identifier) SetObject(v string) *Identifier {
	s.Object = &v
	return s
}

func (s *Identifier) Validate() error {
	return dara.Validate(s)
}
