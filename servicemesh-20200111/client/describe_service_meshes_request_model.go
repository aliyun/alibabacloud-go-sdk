// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*DescribeServiceMeshesRequestTag) *DescribeServiceMeshesRequest
	GetTag() []*DescribeServiceMeshesRequestTag
}

type DescribeServiceMeshesRequest struct {
	// The tags.
	Tag []*DescribeServiceMeshesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesRequest) GetTag() []*DescribeServiceMeshesRequestTag {
	return s.Tag
}

func (s *DescribeServiceMeshesRequest) SetTag(v []*DescribeServiceMeshesRequestTag) *DescribeServiceMeshesRequest {
	s.Tag = v
	return s
}

func (s *DescribeServiceMeshesRequest) Validate() error {
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

type DescribeServiceMeshesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// yahaha
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeServiceMeshesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeServiceMeshesRequestTag) SetKey(v string) *DescribeServiceMeshesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeServiceMeshesRequestTag) SetValue(v string) *DescribeServiceMeshesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeServiceMeshesRequestTag) Validate() error {
	return dara.Validate(s)
}
