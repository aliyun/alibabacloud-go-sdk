// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfirmNotificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmNotificationResponseBody
	GetSuccess() *bool
}

type ConfirmNotificationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A077D99E-0C4D-421E-A5D4-F533F6657817
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmNotificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmNotificationResponseBody) SetRequestId(v string) *ConfirmNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmNotificationResponseBody) SetSuccess(v bool) *ConfirmNotificationResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
