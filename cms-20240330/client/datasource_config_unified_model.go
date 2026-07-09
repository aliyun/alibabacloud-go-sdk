// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasourceConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DatasourceConfigUnified
	GetInstanceId() *string
	SetLegacyRaw(v string) *DatasourceConfigUnified
	GetLegacyRaw() *string
	SetLegacyType(v string) *DatasourceConfigUnified
	GetLegacyType() *string
	SetProductCategory(v string) *DatasourceConfigUnified
	GetProductCategory() *string
	SetRegionId(v string) *DatasourceConfigUnified
	GetRegionId() *string
	SetType(v string) *DatasourceConfigUnified
	GetType() *string
}

type DatasourceConfigUnified struct {
	// The unique instance ID.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The raw string value of the legacy configuration, used for backward compatibility.
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// The legacy type of the data source, used for backward compatibility.
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The product category of the data source.
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The region ID.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The type of the data source.
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DatasourceConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s DatasourceConfigUnified) GoString() string {
	return s.String()
}

func (s *DatasourceConfigUnified) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DatasourceConfigUnified) GetLegacyRaw() *string {
	return s.LegacyRaw
}

func (s *DatasourceConfigUnified) GetLegacyType() *string {
	return s.LegacyType
}

func (s *DatasourceConfigUnified) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *DatasourceConfigUnified) GetRegionId() *string {
	return s.RegionId
}

func (s *DatasourceConfigUnified) GetType() *string {
	return s.Type
}

func (s *DatasourceConfigUnified) SetInstanceId(v string) *DatasourceConfigUnified {
	s.InstanceId = &v
	return s
}

func (s *DatasourceConfigUnified) SetLegacyRaw(v string) *DatasourceConfigUnified {
	s.LegacyRaw = &v
	return s
}

func (s *DatasourceConfigUnified) SetLegacyType(v string) *DatasourceConfigUnified {
	s.LegacyType = &v
	return s
}

func (s *DatasourceConfigUnified) SetProductCategory(v string) *DatasourceConfigUnified {
	s.ProductCategory = &v
	return s
}

func (s *DatasourceConfigUnified) SetRegionId(v string) *DatasourceConfigUnified {
	s.RegionId = &v
	return s
}

func (s *DatasourceConfigUnified) SetType(v string) *DatasourceConfigUnified {
	s.Type = &v
	return s
}

func (s *DatasourceConfigUnified) Validate() error {
	return dara.Validate(s)
}
