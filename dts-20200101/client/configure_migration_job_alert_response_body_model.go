// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureMigrationJobAlertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureMigrationJobAlertResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureMigrationJobAlertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureMigrationJobAlertResponseBody
	GetSuccess() *string
}

type ConfigureMigrationJobAlertResponseBody struct {
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
	// 0a2a047516051973705541561d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureMigrationJobAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureMigrationJobAlertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureMigrationJobAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureMigrationJobAlertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureMigrationJobAlertResponseBody) SetErrCode(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetErrMessage(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetRequestId(v string) *ConfigureMigrationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetSuccess(v string) *ConfigureMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
