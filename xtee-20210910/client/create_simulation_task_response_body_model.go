// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSimulationTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateSimulationTaskResponseBody
	GetResultObject() *bool
}

type CreateSimulationTaskResponseBody struct {
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

func (s CreateSimulationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimulationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimulationTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateSimulationTaskResponseBody) SetRequestId(v string) *CreateSimulationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimulationTaskResponseBody) SetResultObject(v bool) *CreateSimulationTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateSimulationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
