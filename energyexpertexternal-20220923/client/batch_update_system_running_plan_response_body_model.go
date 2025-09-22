// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateSystemRunningPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *BatchUpdateSystemRunningPlanResponseBody
	GetData() *bool
	SetRequestId(v string) *BatchUpdateSystemRunningPlanResponseBody
	GetRequestId() *string
}

type BatchUpdateSystemRunningPlanResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchUpdateSystemRunningPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *BatchUpdateSystemRunningPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateSystemRunningPlanResponseBody) SetData(v bool) *BatchUpdateSystemRunningPlanResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponseBody) SetRequestId(v string) *BatchUpdateSystemRunningPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
