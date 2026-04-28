// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBoxPrivileges interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureAttrId(v string) *DataBoxPrivileges
	GetFeatureAttrId() *string
	SetFeatureId(v string) *DataBoxPrivileges
	GetFeatureId() *string
	SetQuota(v int64) *DataBoxPrivileges
	GetQuota() *int64
}

type DataBoxPrivileges struct {
	FeatureAttrId *string `json:"feature_attr_id,omitempty" xml:"feature_attr_id,omitempty"`
	FeatureId     *string `json:"feature_id,omitempty" xml:"feature_id,omitempty"`
	Quota         *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s DataBoxPrivileges) String() string {
	return dara.Prettify(s)
}

func (s DataBoxPrivileges) GoString() string {
	return s.String()
}

func (s *DataBoxPrivileges) GetFeatureAttrId() *string {
	return s.FeatureAttrId
}

func (s *DataBoxPrivileges) GetFeatureId() *string {
	return s.FeatureId
}

func (s *DataBoxPrivileges) GetQuota() *int64 {
	return s.Quota
}

func (s *DataBoxPrivileges) SetFeatureAttrId(v string) *DataBoxPrivileges {
	s.FeatureAttrId = &v
	return s
}

func (s *DataBoxPrivileges) SetFeatureId(v string) *DataBoxPrivileges {
	s.FeatureId = &v
	return s
}

func (s *DataBoxPrivileges) SetQuota(v int64) *DataBoxPrivileges {
	s.Quota = &v
	return s
}

func (s *DataBoxPrivileges) Validate() error {
	return dara.Validate(s)
}
