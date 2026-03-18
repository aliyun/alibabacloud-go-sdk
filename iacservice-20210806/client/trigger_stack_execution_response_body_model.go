// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerStackExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerStackExecutionResponseBody
	GetRequestId() *string
	SetTriggerId(v string) *TriggerStackExecutionResponseBody
	GetTriggerId() *string
}

type TriggerStackExecutionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F2D40488-3F74-568B-87EC-1C04D098DF8B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// event-xxxx
	TriggerId *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s TriggerStackExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerStackExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerStackExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerStackExecutionResponseBody) GetTriggerId() *string {
	return s.TriggerId
}

func (s *TriggerStackExecutionResponseBody) SetRequestId(v string) *TriggerStackExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerStackExecutionResponseBody) SetTriggerId(v string) *TriggerStackExecutionResponseBody {
	s.TriggerId = &v
	return s
}

func (s *TriggerStackExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
