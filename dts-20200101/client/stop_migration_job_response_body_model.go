// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StopMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StopMigrationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *StopMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StopMigrationJobResponseBody
	GetSuccess() *string
}

type StopMigrationJobResponseBody struct {
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query all data migration instances.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// C306C198-7807-409D-930A-D6CE6C32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StopMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StopMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StopMigrationJobResponseBody) SetErrCode(v string) *StopMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetErrMessage(v string) *StopMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetRequestId(v string) *StopMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetSuccess(v string) *StopMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
