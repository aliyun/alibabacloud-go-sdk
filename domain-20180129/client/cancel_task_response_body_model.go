// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelTaskResponseBody
	GetRequestId() *string
}

type CancelTaskResponseBody struct {
	// example:
	//
	// 010E55C9-C64C-4C85-9BB2-7C225ADA6C86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
