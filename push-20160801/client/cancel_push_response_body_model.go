// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelPushResponseBody
	GetRequestId() *string
}

type CancelPushResponseBody struct {
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPushResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPushResponseBody) SetRequestId(v string) *CancelPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPushResponseBody) Validate() error {
	return dara.Validate(s)
}
