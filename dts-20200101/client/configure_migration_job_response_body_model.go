// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureMigrationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureMigrationJobResponseBody
	GetSuccess() *string
}

type ConfigureMigrationJobResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40E35BD9-002E-4D63-9BE5-FBA48833****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureMigrationJobResponseBody) SetErrCode(v string) *ConfigureMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetErrMessage(v string) *ConfigureMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetRequestId(v string) *ConfigureMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetSuccess(v string) *ConfigureMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
