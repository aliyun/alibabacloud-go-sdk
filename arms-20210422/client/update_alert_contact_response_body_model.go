// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *UpdateAlertContactResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateAlertContactResponseBody
	GetRequestId() *string
}

type UpdateAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateAlertContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertContactResponseBody) SetIsSuccess(v bool) *UpdateAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactResponseBody) SetRequestId(v string) *UpdateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertContactResponseBody) Validate() error {
	return dara.Validate(s)
}
