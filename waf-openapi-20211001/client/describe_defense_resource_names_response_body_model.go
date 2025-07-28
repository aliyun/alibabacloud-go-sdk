// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseResourceNamesResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribeDefenseResourceNamesResponseBody
	GetResources() []*string
	SetTotalCount(v int64) *DescribeDefenseResourceNamesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseResourceNamesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C1823E96-EF4B-5BD2-9E02-1D18****3ED8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The names of the protected objects.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 75
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceNamesResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribeDefenseResourceNamesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseResourceNamesResponseBody) SetRequestId(v string) *DescribeDefenseResourceNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceNamesResponseBody) SetResources(v []*string) *DescribeDefenseResourceNamesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseResourceNamesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceNamesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseResourceNamesResponseBody) Validate() error {
	return dara.Validate(s)
}
