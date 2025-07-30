// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureSubscriptionInstanceAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSubscriptionInstanceAlertResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureSubscriptionInstanceAlertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSubscriptionInstanceAlertResponseBody
	GetSuccess() *string
}

type ConfigureSubscriptionInstanceAlertResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// InternalError  The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a2a047516051973705541561d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSubscriptionInstanceAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetRequestId(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
