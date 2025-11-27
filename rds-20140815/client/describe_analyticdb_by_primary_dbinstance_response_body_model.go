// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalyticdbByPrimaryDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyticDBCount(v int32) *DescribeAnalyticdbByPrimaryDBInstanceResponseBody
	GetAnalyticDBCount() *int32
	SetRequestId(v string) *DescribeAnalyticdbByPrimaryDBInstanceResponseBody
	GetRequestId() *string
}

type DescribeAnalyticdbByPrimaryDBInstanceResponseBody struct {
	// The number of associated analytic instances.
	//
	// example:
	//
	// 0
	AnalyticDBCount *int32 `json:"AnalyticDBCount,omitempty" xml:"AnalyticDBCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 77862BFF-ED59-552A-A2E8-692FEAFC9527
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAnalyticdbByPrimaryDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalyticdbByPrimaryDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) GetAnalyticDBCount() *int32 {
	return s.AnalyticDBCount
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) SetAnalyticDBCount(v int32) *DescribeAnalyticdbByPrimaryDBInstanceResponseBody {
	s.AnalyticDBCount = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) SetRequestId(v string) *DescribeAnalyticdbByPrimaryDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
