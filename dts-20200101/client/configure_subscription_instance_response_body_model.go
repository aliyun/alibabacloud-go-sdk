// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureSubscriptionInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSubscriptionInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureSubscriptionInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSubscriptionInstanceResponseBody
	GetSuccess() *string
}

type ConfigureSubscriptionInstanceResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0CC15092-8957-4532-B559-B4FB80AC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSubscriptionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSubscriptionInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSubscriptionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSubscriptionInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetRequestId(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
