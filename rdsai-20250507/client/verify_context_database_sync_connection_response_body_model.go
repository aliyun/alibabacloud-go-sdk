// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContextDatabaseSyncConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyContextDatabaseSyncConnectionResponseBody
	GetCode() *string
	SetData(v *VerifyContextDatabaseSyncConnectionResponseBodyData) *VerifyContextDatabaseSyncConnectionResponseBody
	GetData() *VerifyContextDatabaseSyncConnectionResponseBodyData
	SetHttpStatusCode(v int64) *VerifyContextDatabaseSyncConnectionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *VerifyContextDatabaseSyncConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyContextDatabaseSyncConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyContextDatabaseSyncConnectionResponseBody
	GetSuccess() *bool
}

type VerifyContextDatabaseSyncConnectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *VerifyContextDatabaseSyncConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Sample value
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

func (s VerifyContextDatabaseSyncConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyContextDatabaseSyncConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetData() *VerifyContextDatabaseSyncConnectionResponseBodyData {
	return s.Data
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetCode(v string) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetData(v *VerifyContextDatabaseSyncConnectionResponseBodyData) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.Data = v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetHttpStatusCode(v int64) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetMessage(v string) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetRequestId(v string) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) SetSuccess(v bool) *VerifyContextDatabaseSyncConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyContextDatabaseSyncConnectionResponseBodyData struct {
	// The authorization status.
	//
	// example:
	//
	// ready
	AuthState *string `json:"AuthState,omitempty" xml:"AuthState,omitempty"`
	// The connection ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s VerifyContextDatabaseSyncConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyContextDatabaseSyncConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyContextDatabaseSyncConnectionResponseBodyData) GetAuthState() *string {
	return s.AuthState
}

func (s *VerifyContextDatabaseSyncConnectionResponseBodyData) GetId() *string {
	return s.Id
}

func (s *VerifyContextDatabaseSyncConnectionResponseBodyData) SetAuthState(v string) *VerifyContextDatabaseSyncConnectionResponseBodyData {
	s.AuthState = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBodyData) SetId(v string) *VerifyContextDatabaseSyncConnectionResponseBodyData {
	s.Id = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
