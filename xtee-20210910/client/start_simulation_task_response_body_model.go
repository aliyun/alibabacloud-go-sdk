// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSimulationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartSimulationTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *StartSimulationTaskResponseBody
	GetResultObject() *bool
}

type StartSimulationTaskResponseBody struct {
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

func (s StartSimulationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSimulationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartSimulationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSimulationTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *StartSimulationTaskResponseBody) SetRequestId(v string) *StartSimulationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSimulationTaskResponseBody) SetResultObject(v bool) *StartSimulationTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *StartSimulationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
