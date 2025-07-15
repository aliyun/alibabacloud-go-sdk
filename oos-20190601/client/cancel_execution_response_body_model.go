// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelExecutionResponseBody
	GetRequestId() *string
}

type CancelExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelExecutionResponseBody) SetRequestId(v string) *CancelExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
