// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContextDatabaseSyncJobExclusionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetCode() *string
	SetData(v *RemoveContextDatabaseSyncJobExclusionResponseBodyData) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetData() *RemoveContextDatabaseSyncJobExclusionResponseBodyData
	SetHttpStatusCode(v int64) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveContextDatabaseSyncJobExclusionResponseBody
	GetSuccess() *bool
}

type RemoveContextDatabaseSyncJobExclusionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *RemoveContextDatabaseSyncJobExclusionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveContextDatabaseSyncJobExclusionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveContextDatabaseSyncJobExclusionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetData() *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	return s.Data
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetCode(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetData(v *RemoveContextDatabaseSyncJobExclusionResponseBodyData) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.Data = v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetHttpStatusCode(v int64) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetMessage(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetRequestId(v string) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) SetSuccess(v bool) *RemoveContextDatabaseSyncJobExclusionResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveContextDatabaseSyncJobExclusionResponseBodyData struct {
	// Indicates whether the data has changed.
	//
	// example:
	//
	// true
	Changed *bool `json:"Changed,omitempty" xml:"Changed,omitempty"`
	// The configuration version number.
	//
	// example:
	//
	// 3
	ConfigVersion *int32 `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// The document ID.
	//
	// example:
	//
	// doc-as1llqmkol****
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// Indicates whether the item is excluded.
	//
	// example:
	//
	// true
	Excluded *bool `json:"Excluded,omitempty" xml:"Excluded,omitempty"`
	// The job ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RemoveContextDatabaseSyncJobExclusionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveContextDatabaseSyncJobExclusionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) GetChanged() *bool {
	return s.Changed
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) GetConfigVersion() *int32 {
	return s.ConfigVersion
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) GetDocumentId() *string {
	return s.DocumentId
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) GetExcluded() *bool {
	return s.Excluded
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) SetChanged(v bool) *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	s.Changed = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) SetConfigVersion(v int32) *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	s.ConfigVersion = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) SetDocumentId(v string) *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	s.DocumentId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) SetExcluded(v bool) *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	s.Excluded = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) SetJobId(v string) *RemoveContextDatabaseSyncJobExclusionResponseBodyData {
	s.JobId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
