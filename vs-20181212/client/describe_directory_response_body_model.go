// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribeDirectoryResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeDirectoryResponseBody
	GetDescription() *string
	SetGroupId(v string) *DescribeDirectoryResponseBody
	GetGroupId() *string
	SetId(v string) *DescribeDirectoryResponseBody
	GetId() *string
	SetName(v string) *DescribeDirectoryResponseBody
	GetName() *string
	SetParentId(v string) *DescribeDirectoryResponseBody
	GetParentId() *string
	SetRequestId(v string) *DescribeDirectoryResponseBody
	GetRequestId() *string
}

type DescribeDirectoryResponseBody struct {
	// example:
	//
	// 2021-09-10T10:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 399*****488-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDirectoryResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDirectoryResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeDirectoryResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDirectoryResponseBody) SetCreatedTime(v string) *DescribeDirectoryResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetDescription(v string) *DescribeDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetGroupId(v string) *DescribeDirectoryResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetId(v string) *DescribeDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetName(v string) *DescribeDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetParentId(v string) *DescribeDirectoryResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeDirectoryResponseBody) SetRequestId(v string) *DescribeDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
