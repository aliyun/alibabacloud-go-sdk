// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContextDatabaseSyncJobResponseBody
	GetCode() *string
	SetData(v *UpdateContextDatabaseSyncJobResponseBodyData) *UpdateContextDatabaseSyncJobResponseBody
	GetData() *UpdateContextDatabaseSyncJobResponseBodyData
	SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncJobResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateContextDatabaseSyncJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContextDatabaseSyncJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContextDatabaseSyncJobResponseBody
	GetSuccess() *bool
}

type UpdateContextDatabaseSyncJobResponseBody struct {
	// example:
	//
	// 示例值
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateContextDatabaseSyncJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateContextDatabaseSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetData() *UpdateContextDatabaseSyncJobResponseBodyData {
	return s.Data
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseSyncJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetCode(v string) *UpdateContextDatabaseSyncJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetData(v *UpdateContextDatabaseSyncJobResponseBodyData) *UpdateContextDatabaseSyncJobResponseBody {
	s.Data = v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetMessage(v string) *UpdateContextDatabaseSyncJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetRequestId(v string) *UpdateContextDatabaseSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) SetSuccess(v bool) *UpdateContextDatabaseSyncJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseSyncJobResponseBodyData struct {
	// example:
	//
	// 3
	ConfigVersion *int32 `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateContextDatabaseSyncJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobResponseBodyData) GetConfigVersion() *int32 {
	return s.ConfigVersion
}

func (s *UpdateContextDatabaseSyncJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateContextDatabaseSyncJobResponseBodyData) SetConfigVersion(v int32) *UpdateContextDatabaseSyncJobResponseBodyData {
	s.ConfigVersion = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBodyData) SetId(v string) *UpdateContextDatabaseSyncJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
