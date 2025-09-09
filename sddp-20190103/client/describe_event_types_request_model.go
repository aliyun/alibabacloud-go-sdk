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
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of anomalous event for which you want to query the anomalous events of subtypes. Valid values:
	//
	// 	- **01**: anomalous permission usage
	//
	// 	- **02**: anomalous data flow
	//
	// 	- **03**: anomalous data operation
	//
	// example:
	//
	// 01
	ParentTypeId *int64 `json:"ParentTypeId,omitempty" xml:"ParentTypeId,omitempty"`
	// The type of the resource. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	ResourceId *int32 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the anomalous event. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **2**: disabled
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
