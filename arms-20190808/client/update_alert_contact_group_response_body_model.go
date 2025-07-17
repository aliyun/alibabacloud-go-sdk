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
	// Indicates whether the call was successful.
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9319A57D-2D9E-472A-B69B-CF3CD16D****
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
