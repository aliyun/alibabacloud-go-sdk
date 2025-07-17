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
	// Indicates whether the alert contact was updated. Valid values:
	//
	// 	- true: The alert contact was updated.
	//
	// 	- false: The alert contact failed to be updated.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A474FF8-7861-4D00-81B5-5BC3DA4E****
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
