// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSourceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupSourceEndpointResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyBackupSourceEndpointResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupSourceEndpointResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyBackupSourceEndpointResponseBody
	GetHttpStatusCode() *int32
	SetNeedPrecheck(v bool) *ModifyBackupSourceEndpointResponseBody
	GetNeedPrecheck() *bool
	SetRequestId(v string) *ModifyBackupSourceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupSourceEndpointResponseBody
	GetSuccess() *bool
}

type ModifyBackupSourceEndpointResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether a precheck is triggered. If the value of this parameter is true, you must start the backup schedule by calling the StartBackupPlan operation.
	//
	// example:
	//
	// false
	NeedPrecheck *bool `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBackupSourceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSourceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupSourceEndpointResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupSourceEndpointResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupSourceEndpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBackupSourceEndpointResponseBody) GetNeedPrecheck() *bool {
	return s.NeedPrecheck
}

func (s *ModifyBackupSourceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupSourceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupSourceEndpointResponseBody) SetBackupPlanId(v string) *ModifyBackupSourceEndpointResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetErrCode(v string) *ModifyBackupSourceEndpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetErrMessage(v string) *ModifyBackupSourceEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetHttpStatusCode(v int32) *ModifyBackupSourceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetNeedPrecheck(v bool) *ModifyBackupSourceEndpointResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetRequestId(v string) *ModifyBackupSourceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetSuccess(v bool) *ModifyBackupSourceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
