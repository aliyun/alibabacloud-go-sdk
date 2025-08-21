// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAnalysisMaxQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQps(v int64) *DescribeAttackAnalysisMaxQpsResponseBody
	GetQps() *int64
	SetRequestId(v string) *DescribeAttackAnalysisMaxQpsResponseBody
	GetRequestId() *string
}

type DescribeAttackAnalysisMaxQpsResponseBody struct {
	// The peak queries per second (QPS) of DDoS attacks. Units: QPS.
	//
	// example:
	//
	// 41652
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8DFB602D-1AAC-46C4-90F2-C84086E7A6E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAttackAnalysisMaxQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAnalysisMaxQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackAnalysisMaxQpsResponseBody) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeAttackAnalysisMaxQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAttackAnalysisMaxQpsResponseBody) SetQps(v int64) *DescribeAttackAnalysisMaxQpsResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsResponseBody) SetRequestId(v string) *DescribeAttackAnalysisMaxQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsResponseBody) Validate() error {
	return dara.Validate(s)
}
