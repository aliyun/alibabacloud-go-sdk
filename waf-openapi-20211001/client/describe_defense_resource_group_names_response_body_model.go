// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupNames(v []*string) *DescribeDefenseResourceGroupNamesResponseBody
	GetGroupNames() []*string
	SetRequestId(v string) *DescribeDefenseResourceGroupNamesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDefenseResourceGroupNamesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseResourceGroupNamesResponseBody struct {
	// The names of the protected object groups.
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 59DA4258-2F32-5095-B283-57AC****70B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) GetGroupNames() []*string {
	return s.GroupNames
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetGroupNames(v []*string) *DescribeDefenseResourceGroupNamesResponseBody {
	s.GroupNames = v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceGroupNamesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) Validate() error {
	return dara.Validate(s)
}
