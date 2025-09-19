// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelOnceTaskResponseBody
	GetRequestId() *string
}

type CancelOnceTaskResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOnceTaskResponseBody) SetRequestId(v string) *CancelOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOnceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
