// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsSuperAccountInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstances(v *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) *DescribeRdsSuperAccountInstancesResponseBody
	GetDbInstances() *DescribeRdsSuperAccountInstancesResponseBodyDbInstances
	SetRequestId(v string) *DescribeRdsSuperAccountInstancesResponseBody
	GetRequestId() *string
}

type DescribeRdsSuperAccountInstancesResponseBody struct {
	// The privileged accounts.
	DbInstances *DescribeRdsSuperAccountInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5D64DE5944A1E541E0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsSuperAccountInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) GetDbInstances() *DescribeRdsSuperAccountInstancesResponseBodyDbInstances {
	return s.DbInstances
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetDbInstances(v *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) *DescribeRdsSuperAccountInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) SetRequestId(v string) *DescribeRdsSuperAccountInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsSuperAccountInstancesResponseBodyDbInstances struct {
	DbInstance []*string `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) GetDbInstance() []*string {
	return s.DbInstance
}

func (s *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) SetDbInstance(v []*string) *DescribeRdsSuperAccountInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponseBodyDbInstances) Validate() error {
	return dara.Validate(s)
}
