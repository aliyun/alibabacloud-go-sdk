// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcAssetCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v int32) *DescribeIdcAssetCriteriaRequest
	GetStatus() *int32
	SetValue(v string) *DescribeIdcAssetCriteriaRequest
	GetValue() *string
}

type DescribeIdcAssetCriteriaRequest struct {
	// Indicates whether the corresponding IP address is valid. Valid values:
	//
	// - **0**: valid
	//
	// - **1**: ignored
	//
	// - **2**: invalid
	//
	// - **3**: expired
	//
	// - **4**: probe does not exist.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The fuzzy match value entered when querying assets.
	//
	// example:
	//
	// testwww
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIdcAssetCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcAssetCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdcAssetCriteriaRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeIdcAssetCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeIdcAssetCriteriaRequest) SetStatus(v int32) *DescribeIdcAssetCriteriaRequest {
	s.Status = &v
	return s
}

func (s *DescribeIdcAssetCriteriaRequest) SetValue(v string) *DescribeIdcAssetCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeIdcAssetCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
