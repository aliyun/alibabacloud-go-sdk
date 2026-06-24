// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeDataLimitSetRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeDataLimitSetRequest
	GetLang() *string
	SetParentId(v string) *DescribeDataLimitSetRequest
	GetParentId() *string
	SetRegionType(v string) *DescribeDataLimitSetRequest
	GetRegionType() *string
	SetResourceType(v int32) *DescribeDataLimitSetRequest
	GetResourceType() *int32
}

type DescribeDataLimitSetRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh_cn**: Chinese (Simplified). This is the default value.
	//
	// - **en_us**: English (US).
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the parent asset.
	//
	// The [DescribeDataLimitDetail](~~DescribeDataLimitDetail~~) or [DescribeDataLimits](~~DescribeDataLimits~~) operation returns this ID in the **ParentId*	- parameter.
	//
	// example:
	//
	// db
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The region type.
	//
	// - **native_audit**: A region that supports traffic collection.
	//
	// example:
	//
	// native_audit
	RegionType *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
	// The type of data asset. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: ADS.
	//
	// - **4**: OTS.
	//
	// - **5**: RDS.
	//
	// - **6**: SELF_DB.
	//
	// > If you set this parameter to a value other than 2, the returned OssBucketList object is empty.
	//
	// example:
	//
	// 2
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDataLimitSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataLimitSetRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataLimitSetRequest) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDataLimitSetRequest) GetRegionType() *string {
	return s.RegionType
}

func (s *DescribeDataLimitSetRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *DescribeDataLimitSetRequest) SetFeatureType(v int32) *DescribeDataLimitSetRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetLang(v string) *DescribeDataLimitSetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetParentId(v string) *DescribeDataLimitSetRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetRegionType(v string) *DescribeDataLimitSetRequest {
	s.RegionType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetResourceType(v int32) *DescribeDataLimitSetRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) Validate() error {
	return dara.Validate(s)
}
