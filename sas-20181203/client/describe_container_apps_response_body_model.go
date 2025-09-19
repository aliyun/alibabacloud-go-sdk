// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeContainerAppsResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *DescribeContainerAppsResponseBody
	GetTagValues() []*string
}

type DescribeContainerAppsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeContainerAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerAppsResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeContainerAppsResponseBody) SetRequestId(v string) *DescribeContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerAppsResponseBody) SetTagValues(v []*string) *DescribeContainerAppsResponseBody {
	s.TagValues = v
	return s
}

func (s *DescribeContainerAppsResponseBody) Validate() error {
	return dara.Validate(s)
}
