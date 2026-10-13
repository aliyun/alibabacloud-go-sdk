// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobExclusionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncJobExclusionsResponseBodyData) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetData() *ListContextDatabaseSyncJobExclusionsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncJobExclusionsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncJobExclusionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncJobExclusionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Example value
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

func (s ListContextDatabaseSyncJobExclusionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobExclusionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetData() *ListContextDatabaseSyncJobExclusionsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetCode(v string) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetData(v *ListContextDatabaseSyncJobExclusionsResponseBodyData) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetMessage(v string) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncJobExclusionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncJobExclusionsResponseBodyData struct {
	// The list of document IDs.
	//
	// example:
	//
	// ["doc-as1llqmkol****"]
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
}

func (s ListContextDatabaseSyncJobExclusionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobExclusionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBodyData) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBodyData) SetDocumentIds(v []*string) *ListContextDatabaseSyncJobExclusionsResponseBodyData {
	s.DocumentIds = v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
