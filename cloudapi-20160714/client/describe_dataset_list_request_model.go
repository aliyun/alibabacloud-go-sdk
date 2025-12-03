// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetIds(v string) *DescribeDatasetListRequest
	GetDatasetIds() *string
	SetDatasetName(v string) *DescribeDatasetListRequest
	GetDatasetName() *string
	SetPageNumber(v int32) *DescribeDatasetListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatasetListRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeDatasetListRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeDatasetListRequestTag) *DescribeDatasetListRequest
	GetTag() []*DescribeDatasetListRequestTag
}

type DescribeDatasetListRequest struct {
	// The IDs of the datasets.
	//
	// example:
	//
	// 4add6a61804e47858266883e********
	DatasetIds *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// IPwhitelist
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specify the object labels to which the rule applies, and multiple labels can be set
	Tag []*DescribeDatasetListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDatasetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListRequest) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *DescribeDatasetListRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DescribeDatasetListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatasetListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatasetListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDatasetListRequest) GetTag() []*DescribeDatasetListRequestTag {
	return s.Tag
}

func (s *DescribeDatasetListRequest) SetDatasetIds(v string) *DescribeDatasetListRequest {
	s.DatasetIds = &v
	return s
}

func (s *DescribeDatasetListRequest) SetDatasetName(v string) *DescribeDatasetListRequest {
	s.DatasetName = &v
	return s
}

func (s *DescribeDatasetListRequest) SetPageNumber(v int32) *DescribeDatasetListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatasetListRequest) SetPageSize(v int32) *DescribeDatasetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatasetListRequest) SetSecurityToken(v string) *DescribeDatasetListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDatasetListRequest) SetTag(v []*DescribeDatasetListRequestTag) *DescribeDatasetListRequest {
	s.Tag = v
	return s
}

func (s *DescribeDatasetListRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatasetListRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDatasetListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDatasetListRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDatasetListRequestTag) SetKey(v string) *DescribeDatasetListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDatasetListRequestTag) SetValue(v string) *DescribeDatasetListRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDatasetListRequestTag) Validate() error {
	return dara.Validate(s)
}
