// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*string) *DescribeDefenseTemplateValidGroupsResponseBody
	GetGroups() []*string
	SetRequestId(v string) *DescribeDefenseTemplateValidGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDefenseTemplateValidGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseTemplateValidGroupsResponseBody struct {
	// The names of the protected object groups.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6EA4B39A-9C0C-5E57-993E-30B6****3AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 27
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) GetGroups() []*string {
	return s.Groups
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetGroups(v []*string) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetRequestId(v string) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetTotalCount(v int64) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
