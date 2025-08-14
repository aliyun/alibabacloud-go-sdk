// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSimulationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSimulationTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *StopSimulationTaskResponseBody
	GetResultObject() *bool
}

type StopSimulationTaskResponseBody struct {
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

func (s StopSimulationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSimulationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopSimulationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSimulationTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *StopSimulationTaskResponseBody) SetRequestId(v string) *StopSimulationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSimulationTaskResponseBody) SetResultObject(v bool) *StopSimulationTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *StopSimulationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
