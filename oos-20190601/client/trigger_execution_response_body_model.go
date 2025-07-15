// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerExecutionResponseBody
	GetRequestId() *string
}

type TriggerExecutionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerExecutionResponseBody) SetRequestId(v string) *TriggerExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
