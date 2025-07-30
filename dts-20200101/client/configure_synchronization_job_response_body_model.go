// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSynchronizationJobResponseBody
	GetSuccess() *string
}

type ConfigureSynchronizationJobResponseBody struct {
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
	// 2690E467-7773-43BC-A009-370EE2E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
