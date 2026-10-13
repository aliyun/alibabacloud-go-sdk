// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateContextDatabaseSyncSourceResponseBody
	GetCode() *string
	SetData(v *CreateContextDatabaseSyncSourceResponseBodyData) *CreateContextDatabaseSyncSourceResponseBody
	GetData() *CreateContextDatabaseSyncSourceResponseBodyData
	SetHttpStatusCode(v int64) *CreateContextDatabaseSyncSourceResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateContextDatabaseSyncSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateContextDatabaseSyncSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateContextDatabaseSyncSourceResponseBody
	GetSuccess() *bool
}

type CreateContextDatabaseSyncSourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// SampleValue
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateContextDatabaseSyncSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// SampleValue
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateContextDatabaseSyncSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetData() *CreateContextDatabaseSyncSourceResponseBodyData {
	return s.Data
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDatabaseSyncSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetCode(v string) *CreateContextDatabaseSyncSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetData(v *CreateContextDatabaseSyncSourceResponseBodyData) *CreateContextDatabaseSyncSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetHttpStatusCode(v int64) *CreateContextDatabaseSyncSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetMessage(v string) *CreateContextDatabaseSyncSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetRequestId(v string) *CreateContextDatabaseSyncSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) SetSuccess(v bool) *CreateContextDatabaseSyncSourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextDatabaseSyncSourceResponseBodyData struct {
	// The data source ID, a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateContextDatabaseSyncSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncSourceResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateContextDatabaseSyncSourceResponseBodyData) SetId(v string) *CreateContextDatabaseSyncSourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
