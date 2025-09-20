// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRow interface {
	dara.Model
	String() string
	GoString() string
	SetCustomLabels(v []*KeyValuePair) *Row
	GetCustomLabels() []*KeyValuePair
	SetURI(v string) *Row
	GetURI() *string
}

type Row struct {
	CustomLabels []*KeyValuePair `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty" type:"Repeated"`
	URI          *string         `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s Row) String() string {
	return dara.Prettify(s)
}

func (s Row) GoString() string {
	return s.String()
}

func (s *Row) GetCustomLabels() []*KeyValuePair {
	return s.CustomLabels
}

func (s *Row) GetURI() *string {
	return s.URI
}

func (s *Row) SetCustomLabels(v []*KeyValuePair) *Row {
	s.CustomLabels = v
	return s
}

func (s *Row) SetURI(v string) *Row {
	s.URI = &v
	return s
}

func (s *Row) Validate() error {
	return dara.Validate(s)
}
