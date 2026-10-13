// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetCode() *string
	SetData(v *UpdateContextDatabaseSyncJobScheduleResponseBodyData) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetData() *UpdateContextDatabaseSyncJobScheduleResponseBodyData
	SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContextDatabaseSyncJobScheduleResponseBody
	GetSuccess() *bool
}

type UpdateContextDatabaseSyncJobScheduleResponseBody struct {
	// example:
	//
	// 示例值
	Code *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateContextDatabaseSyncJobScheduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateContextDatabaseSyncJobScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetData() *UpdateContextDatabaseSyncJobScheduleResponseBodyData {
	return s.Data
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetCode(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetData(v *UpdateContextDatabaseSyncJobScheduleResponseBodyData) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetMessage(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetRequestId(v string) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) SetSuccess(v bool) *UpdateContextDatabaseSyncJobScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseSyncJobScheduleResponseBodyData struct {
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateContextDatabaseSyncJobScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBodyData) SetJobId(v string) *UpdateContextDatabaseSyncJobScheduleResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
