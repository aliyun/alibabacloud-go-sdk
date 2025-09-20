// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegionType interface {
	dara.Model
	String() string
	GoString() string
	SetLocalName(v string) *RegionType
	GetLocalName() *string
	SetRegionId(v string) *RegionType
	GetRegionId() *string
}

type RegionType struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegionType) String() string {
	return dara.Prettify(s)
}

func (s RegionType) GoString() string {
	return s.String()
}

func (s *RegionType) GetLocalName() *string {
	return s.LocalName
}

func (s *RegionType) GetRegionId() *string {
	return s.RegionId
}

func (s *RegionType) SetLocalName(v string) *RegionType {
	s.LocalName = &v
	return s
}

func (s *RegionType) SetRegionId(v string) *RegionType {
	s.RegionId = &v
	return s
}

func (s *RegionType) Validate() error {
	return dara.Validate(s)
}
