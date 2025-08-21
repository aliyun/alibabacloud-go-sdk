// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckOperatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAckOperatorResponseBody
	GetRequestId() *string
	SetResult(v *DescribeAckOperatorResponseBodyResult) *DescribeAckOperatorResponseBody
	GetResult() *DescribeAckOperatorResponseBodyResult
}

type DescribeAckOperatorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6615EE8D-FD9D-4FD3-997E-6FEA5B8D82ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DescribeAckOperatorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAckOperatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckOperatorResponseBody) GetResult() *DescribeAckOperatorResponseBodyResult {
	return s.Result
}

func (s *DescribeAckOperatorResponseBody) SetRequestId(v string) *DescribeAckOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckOperatorResponseBody) SetResult(v *DescribeAckOperatorResponseBodyResult) *DescribeAckOperatorResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAckOperatorResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAckOperatorResponseBodyResult struct {
	// The installation status of ES-operator. Valid values:
	//
	// 	- deployed: ES-operator is installed.
	//
	// 	- not-deploy: ES-operator is not installed.
	//
	// 	- failed: ES-operator fails to be installed.
	//
	// 	- unknown: The installation status of ES-operator is unknown.
	//
	// example:
	//
	// deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The version of ES-operator.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAckOperatorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckOperatorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeAckOperatorResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeAckOperatorResponseBodyResult) SetStatus(v string) *DescribeAckOperatorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAckOperatorResponseBodyResult) SetVersion(v string) *DescribeAckOperatorResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeAckOperatorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
