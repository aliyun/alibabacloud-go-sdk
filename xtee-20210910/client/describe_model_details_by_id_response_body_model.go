// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelDetailsByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeModelDetailsByIdResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeModelDetailsByIdResponseBodyResultObject) *DescribeModelDetailsByIdResponseBody
	GetResultObject() []*DescribeModelDetailsByIdResponseBodyResultObject
}

type DescribeModelDetailsByIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject []*DescribeModelDetailsByIdResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s DescribeModelDetailsByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelDetailsByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelDetailsByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelDetailsByIdResponseBody) GetResultObject() []*DescribeModelDetailsByIdResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeModelDetailsByIdResponseBody) SetRequestId(v string) *DescribeModelDetailsByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelDetailsByIdResponseBody) SetResultObject(v []*DescribeModelDetailsByIdResponseBodyResultObject) *DescribeModelDetailsByIdResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeModelDetailsByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeModelDetailsByIdResponseBodyResultObject struct {
	// Model prediction result.
	//
	// example:
	//
	// {\\"AUC\\":0.9895246624946594,\\"Count\\":2489,\\"DecisionMark\\":[0.0,0.0010000000474974513,0.05287817938420348,0.0]}
	ModelEffectEvaluation *string `json:"modelEffectEvaluation,omitempty" xml:"modelEffectEvaluation,omitempty"`
}

func (s DescribeModelDetailsByIdResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelDetailsByIdResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeModelDetailsByIdResponseBodyResultObject) GetModelEffectEvaluation() *string {
	return s.ModelEffectEvaluation
}

func (s *DescribeModelDetailsByIdResponseBodyResultObject) SetModelEffectEvaluation(v string) *DescribeModelDetailsByIdResponseBodyResultObject {
	s.ModelEffectEvaluation = &v
	return s
}

func (s *DescribeModelDetailsByIdResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
