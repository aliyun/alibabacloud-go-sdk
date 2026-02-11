// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *StopAlertResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *StopAlertResponseBody
	GetRequestId() *string
}

type StopAlertResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StopAlertResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *StopAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAlertResponseBody) SetIsSuccess(v bool) *StopAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StopAlertResponseBody) SetRequestId(v string) *StopAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
