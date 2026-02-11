// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteAlertContactResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteAlertContactResponseBody
	GetRequestId() *string
}

type DeleteAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteAlertContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertContactResponseBody) SetIsSuccess(v bool) *DeleteAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactResponseBody) SetRequestId(v string) *DeleteAlertContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertContactResponseBody) Validate() error {
	return dara.Validate(s)
}
