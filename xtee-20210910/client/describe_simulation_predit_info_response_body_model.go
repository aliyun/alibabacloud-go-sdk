// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationPreditInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSimulationPreditInfoResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSimulationPreditInfoResponseBody
	GetResultObject() *bool
}

type DescribeSimulationPreditInfoResponseBody struct {
	// Request ID
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

func (s DescribeSimulationPreditInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationPreditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimulationPreditInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimulationPreditInfoResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSimulationPreditInfoResponseBody) SetRequestId(v string) *DescribeSimulationPreditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimulationPreditInfoResponseBody) SetResultObject(v bool) *DescribeSimulationPreditInfoResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSimulationPreditInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
