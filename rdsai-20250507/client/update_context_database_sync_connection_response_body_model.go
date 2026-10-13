// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContextDatabaseSyncConnectionResponseBody
	GetCode() *string
	SetData(v *UpdateContextDatabaseSyncConnectionResponseBodyData) *UpdateContextDatabaseSyncConnectionResponseBody
	GetData() *UpdateContextDatabaseSyncConnectionResponseBodyData
	SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncConnectionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateContextDatabaseSyncConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContextDatabaseSyncConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContextDatabaseSyncConnectionResponseBody
	GetSuccess() *bool
}

type UpdateContextDatabaseSyncConnectionResponseBody struct {
	// example:
	//
	// 示例值
	Code *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateContextDatabaseSyncConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateContextDatabaseSyncConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetData() *UpdateContextDatabaseSyncConnectionResponseBodyData {
	return s.Data
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetCode(v string) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetData(v *UpdateContextDatabaseSyncConnectionResponseBodyData) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetMessage(v string) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetRequestId(v string) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) SetSuccess(v bool) *UpdateContextDatabaseSyncConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseSyncConnectionResponseBodyData struct {
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateContextDatabaseSyncConnectionResponseBodyData) SetId(v string) *UpdateContextDatabaseSyncConnectionResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
