// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelDetail interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *LabelDetail
	GetKey() *string
	SetValues(v []*string) *LabelDetail
	GetValues() []*string
}

type LabelDetail struct {
	// The label key
	//
	// example:
	//
	// version
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The list of available values for this label
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s LabelDetail) String() string {
	return dara.Prettify(s)
}

func (s LabelDetail) GoString() string {
	return s.String()
}

func (s *LabelDetail) GetKey() *string {
	return s.Key
}

func (s *LabelDetail) GetValues() []*string {
	return s.Values
}

func (s *LabelDetail) SetKey(v string) *LabelDetail {
	s.Key = &v
	return s
}

func (s *LabelDetail) SetValues(v []*string) *LabelDetail {
	s.Values = v
	return s
}

func (s *LabelDetail) Validate() error {
	return dara.Validate(s)
}
