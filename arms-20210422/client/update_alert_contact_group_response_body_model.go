// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *UpdateAlertContactGroupResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateAlertContactGroupResponseBody
	GetRequestId() *string
}

type UpdateAlertContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateAlertContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertContactGroupResponseBody) SetIsSuccess(v bool) *UpdateAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactGroupResponseBody) SetRequestId(v string) *UpdateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
