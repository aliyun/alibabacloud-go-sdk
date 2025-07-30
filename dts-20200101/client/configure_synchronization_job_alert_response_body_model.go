// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureSynchronizationJobAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSynchronizationJobAlertResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureSynchronizationJobAlertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSynchronizationJobAlertResponseBody
	GetSuccess() *string
}

type ConfigureSynchronizationJobAlertResponseBody struct {
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
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8EEE7858-7D41-4EDF-9435-AEED2A34****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSynchronizationJobAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSynchronizationJobAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSynchronizationJobAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSynchronizationJobAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
