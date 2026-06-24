// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeEventTypesRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeEventTypesRequest
	GetLang() *string
	SetParentTypeId(v int64) *DescribeEventTypesRequest
	GetParentTypeId() *int64
	SetResourceId(v int32) *DescribeEventTypesRequest
	GetResourceId() *int32
	SetStatus(v int32) *DescribeEventTypesRequest
	GetStatus() *int32
}

type DescribeEventTypesRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the parent anomalous activity type to which the child anomalous activity type belongs. Valid values:
	//
	// - **01**: anomalous permission access.
	//
	// - **02**: anomalous data flow.
	//
	// - **03**: anomalous data operation.
	//
	// example:
	//
	// 01
	ParentTypeId *int64 `json:"ParentTypeId,omitempty" xml:"ParentTypeId,omitempty"`
	// The resource type of the product. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: Object Storage Service (OSS).
	//
	// - **3**: AnalyticDB for MySQL.
	//
	// - **4**: Tablestore.
	//
	// - **5**. ApsaraDB RDS.
	//
	// example:
	//
	// 5
	ResourceId *int32 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status. Valid values:
	//
	// - **1**: active.
	//
	// - **2**: inactive.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeEventTypesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventTypesRequest) GetParentTypeId() *int64 {
	return s.ParentTypeId
}

func (s *DescribeEventTypesRequest) GetResourceId() *int32 {
	return s.ResourceId
}

func (s *DescribeEventTypesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventTypesRequest) SetFeatureType(v int32) *DescribeEventTypesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeEventTypesRequest) SetLang(v string) *DescribeEventTypesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTypesRequest) SetParentTypeId(v int64) *DescribeEventTypesRequest {
	s.ParentTypeId = &v
	return s
}

func (s *DescribeEventTypesRequest) SetResourceId(v int32) *DescribeEventTypesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventTypesRequest) SetStatus(v int32) *DescribeEventTypesRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventTypesRequest) Validate() error {
	return dara.Validate(s)
}
