// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCdnTypes(v *DescribeCdnTypesResponseBodyCdnTypes) *DescribeCdnTypesResponseBody
	GetCdnTypes() *DescribeCdnTypesResponseBodyCdnTypes
	SetRequestId(v string) *DescribeCdnTypesResponseBody
	GetRequestId() *string
}

type DescribeCdnTypesResponseBody struct {
	// The types of the domain names.
	CdnTypes *DescribeCdnTypesResponseBodyCdnTypes `json:"CdnTypes,omitempty" xml:"CdnTypes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BDA62CE4-3477-439A-B52E-D2D7C829D7C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnTypesResponseBody) GetCdnTypes() *DescribeCdnTypesResponseBodyCdnTypes {
	return s.CdnTypes
}

func (s *DescribeCdnTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnTypesResponseBody) SetCdnTypes(v *DescribeCdnTypesResponseBodyCdnTypes) *DescribeCdnTypesResponseBody {
	s.CdnTypes = v
	return s
}

func (s *DescribeCdnTypesResponseBody) SetRequestId(v string) *DescribeCdnTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnTypesResponseBodyCdnTypes struct {
	CdnType []*DescribeCdnTypesResponseBodyCdnTypesCdnType `json:"CdnType,omitempty" xml:"CdnType,omitempty" type:"Repeated"`
}

func (s DescribeCdnTypesResponseBodyCdnTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTypesResponseBodyCdnTypes) GoString() string {
	return s.String()
}

func (s *DescribeCdnTypesResponseBodyCdnTypes) GetCdnType() []*DescribeCdnTypesResponseBodyCdnTypesCdnType {
	return s.CdnType
}

func (s *DescribeCdnTypesResponseBodyCdnTypes) SetCdnType(v []*DescribeCdnTypesResponseBodyCdnTypesCdnType) *DescribeCdnTypesResponseBodyCdnTypes {
	s.CdnType = v
	return s
}

func (s *DescribeCdnTypesResponseBodyCdnTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnTypesResponseBodyCdnTypesCdnType struct {
	// The description of the domain name type.
	//
	// example:
	//
	// Download Acceleration
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The type of the domain name.
	//
	// example:
	//
	// download
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCdnTypesResponseBodyCdnTypesCdnType) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTypesResponseBodyCdnTypesCdnType) GoString() string {
	return s.String()
}

func (s *DescribeCdnTypesResponseBodyCdnTypesCdnType) GetDesc() *string {
	return s.Desc
}

func (s *DescribeCdnTypesResponseBodyCdnTypesCdnType) GetType() *string {
	return s.Type
}

func (s *DescribeCdnTypesResponseBodyCdnTypesCdnType) SetDesc(v string) *DescribeCdnTypesResponseBodyCdnTypesCdnType {
	s.Desc = &v
	return s
}

func (s *DescribeCdnTypesResponseBodyCdnTypesCdnType) SetType(v string) *DescribeCdnTypesResponseBodyCdnTypesCdnType {
	s.Type = &v
	return s
}

func (s *DescribeCdnTypesResponseBodyCdnTypesCdnType) Validate() error {
	return dara.Validate(s)
}
