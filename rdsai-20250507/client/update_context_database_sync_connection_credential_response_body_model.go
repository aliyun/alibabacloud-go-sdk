// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetCode() *string
	SetData(v *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetData() *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData
	SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContextDatabaseSyncConnectionCredentialResponseBody
	GetSuccess() *bool
}

type UpdateContextDatabaseSyncConnectionCredentialResponseBody struct {
	// example:
	//
	// 示例值
	Code *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateContextDatabaseSyncConnectionCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetData() *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData {
	return s.Data
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetCode(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetData(v *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.Data = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetMessage(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetRequestId(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) SetSuccess(v bool) *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseSyncConnectionCredentialResponseBodyData struct {
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) SetId(v string) *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
