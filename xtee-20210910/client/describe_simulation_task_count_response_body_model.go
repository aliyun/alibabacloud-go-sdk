// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSimulationTaskCountResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSimulationTaskCountResponseBody
	GetResultObject() *bool
}

type DescribeSimulationTaskCountResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeSimulationTaskCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimulationTaskCountResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSimulationTaskCountResponseBody) SetRequestId(v string) *DescribeSimulationTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimulationTaskCountResponseBody) SetResultObject(v bool) *DescribeSimulationTaskCountResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSimulationTaskCountResponseBody) Validate() error {
	return dara.Validate(s)
}
