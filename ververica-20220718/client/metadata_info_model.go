// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataInfo interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetadataInfo
	GetKey() *string
	SetVirtual(v bool) *MetadataInfo
	GetVirtual() *bool
}

type MetadataInfo struct {
	Key     *string `json:"key,omitempty" xml:"key,omitempty"`
	Virtual *bool   `json:"virtual,omitempty" xml:"virtual,omitempty"`
}

func (s MetadataInfo) String() string {
	return dara.Prettify(s)
}

func (s MetadataInfo) GoString() string {
	return s.String()
}

func (s *MetadataInfo) GetKey() *string {
	return s.Key
}

func (s *MetadataInfo) GetVirtual() *bool {
	return s.Virtual
}

func (s *MetadataInfo) SetKey(v string) *MetadataInfo {
	s.Key = &v
	return s
}

func (s *MetadataInfo) SetVirtual(v bool) *MetadataInfo {
	s.Virtual = &v
	return s
}

func (s *MetadataInfo) Validate() error {
	return dara.Validate(s)
}
