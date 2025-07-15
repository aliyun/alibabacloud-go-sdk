// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *DescribeDatasetItemInfoRequest
	GetDatasetId() *string
	SetDatasetItemId(v string) *DescribeDatasetItemInfoRequest
	GetDatasetItemId() *string
	SetSecurityToken(v string) *DescribeDatasetItemInfoRequest
	GetSecurityToken() *string
	SetValue(v string) *DescribeDatasetItemInfoRequest
	GetValue() *string
}

type DescribeDatasetItemInfoRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62b91a790a693238********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the data entry.
	//
	// example:
	//
	// 5045****
	DatasetItemId *string `json:"DatasetItemId,omitempty" xml:"DatasetItemId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The value of the data entry.
	//
	// example:
	//
	// 106.43.XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDatasetItemInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemInfoRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetItemInfoRequest) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *DescribeDatasetItemInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDatasetItemInfoRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeDatasetItemInfoRequest) SetDatasetId(v string) *DescribeDatasetItemInfoRequest {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetItemInfoRequest) SetDatasetItemId(v string) *DescribeDatasetItemInfoRequest {
	s.DatasetItemId = &v
	return s
}

func (s *DescribeDatasetItemInfoRequest) SetSecurityToken(v string) *DescribeDatasetItemInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDatasetItemInfoRequest) SetValue(v string) *DescribeDatasetItemInfoRequest {
	s.Value = &v
	return s
}

func (s *DescribeDatasetItemInfoRequest) Validate() error {
	return dara.Validate(s)
}
