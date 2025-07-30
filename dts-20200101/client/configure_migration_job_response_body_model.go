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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The start offset of incremental data migration. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 40E35BD9-002E-4D63-9BE5-FBA48833****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SID of the Oracle database.
	//
	// >  You must specify this parameter only if the **DestinationEndpoint.EngineName*	- parameter is set to **Oracle*	- and the **Oracle*	- database is deployed in a non-RAC architecture.
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
