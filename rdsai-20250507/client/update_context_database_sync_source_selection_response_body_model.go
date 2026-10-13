// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncSourceSelectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetCode() *string
	SetData(v *UpdateContextDatabaseSyncSourceSelectionResponseBodyData) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetData() *UpdateContextDatabaseSyncSourceSelectionResponseBodyData
	SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContextDatabaseSyncSourceSelectionResponseBody
	GetSuccess() *bool
}

type UpdateContextDatabaseSyncSourceSelectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdateContextDatabaseSyncSourceSelectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateContextDatabaseSyncSourceSelectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceSelectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetData() *UpdateContextDatabaseSyncSourceSelectionResponseBodyData {
	return s.Data
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetCode(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetData(v *UpdateContextDatabaseSyncSourceSelectionResponseBodyData) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetHttpStatusCode(v int64) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetMessage(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetRequestId(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) SetSuccess(v bool) *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseSyncSourceSelectionResponseBodyData struct {
	// The data source ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s UpdateContextDatabaseSyncSourceSelectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceSelectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBodyData) SetSourceId(v string) *UpdateContextDatabaseSyncSourceSelectionResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
