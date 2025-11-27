// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfirmNotifyResponseBody
	GetRequestId() *string
}

type ConfirmNotifyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6919FA39-04CF-571F-A6B0-DCC7AECB4170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmNotifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmNotifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmNotifyResponseBody) SetRequestId(v string) *ConfirmNotifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmNotifyResponseBody) Validate() error {
	return dara.Validate(s)
}
