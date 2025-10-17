// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClustersRequest
	GetClusterId() *string
	SetClusterName(v string) *ListClustersRequest
	GetClusterName() *string
	SetPageNum(v int32) *ListClustersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListClustersRequest
	GetPageSize() *int32
	SetTag(v []*ListClustersRequestTag) *ListClustersRequest
	GetTag() []*ListClustersRequestTag
}

type ListClustersRequest struct {
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// cluster-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 5
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag      []*ListClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersRequest) GetTag() []*ListClustersRequestTag {
	return s.Tag
}

func (s *ListClustersRequest) SetClusterId(v string) *ListClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClustersRequest) SetClusterName(v string) *ListClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *ListClustersRequest) SetPageNum(v int32) *ListClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetTag(v []*ListClustersRequestTag) *ListClustersRequest {
	s.Tag = v
	return s
}

func (s *ListClustersRequest) Validate() error {
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

type ListClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequestTag) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListClustersRequestTag) SetKey(v string) *ListClustersRequestTag {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTag) SetValue(v string) *ListClustersRequestTag {
	s.Value = &v
	return s
}

func (s *ListClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
