// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetLabel interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DatasetLabel
	GetKey() *string
	SetValue(v string) *DatasetLabel
	GetValue() *string
}

type DatasetLabel struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DatasetLabel) String() string {
	return dara.Prettify(s)
}

func (s DatasetLabel) GoString() string {
	return s.String()
}

func (s *DatasetLabel) GetKey() *string {
	return s.Key
}

func (s *DatasetLabel) GetValue() *string {
	return s.Value
}

func (s *DatasetLabel) SetKey(v string) *DatasetLabel {
	s.Key = &v
	return s
}

func (s *DatasetLabel) SetValue(v string) *DatasetLabel {
	s.Value = &v
	return s
}

func (s *DatasetLabel) Validate() error {
	return dara.Validate(s)
}
