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
	// The status of the IP address. Valid values:
	//
	// 	- **0**: The IP address is valid.
	//
	// 	- **1**: The IP address is ignored.
	//
	// 	- *2*: The IP address is invalid.
	//
	// 	- *3*: The IP address is expired.
	//
	// 	- *4*: The probe that is used to scan the IP address does not exist.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The keyword that is used to match assets in fuzzy mode.
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
