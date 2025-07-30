// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobReplicatorCompareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody
	GetSuccess() *string
}

type ConfigureSynchronizationJobReplicatorCompareResponseBody struct {
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
	// 86A8FF0F-FA92-449D-B559-05CFF9F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSynchronizationJobReplicatorCompareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) Validate() error {
	return dara.Validate(s)
}
