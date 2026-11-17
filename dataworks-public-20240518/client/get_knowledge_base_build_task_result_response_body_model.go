// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseBuildTaskResultResponseBody
	GetCode() *string
	SetData(v *GetKnowledgeBaseBuildTaskResultResponseBodyData) *GetKnowledgeBaseBuildTaskResultResponseBody
	GetData() *GetKnowledgeBaseBuildTaskResultResponseBodyData
	SetMessage(v string) *GetKnowledgeBaseBuildTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseBuildTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeBaseBuildTaskResultResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBaseBuildTaskResultResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *GetKnowledgeBaseBuildTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) GetData() *GetKnowledgeBaseBuildTaskResultResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) SetCode(v string) *GetKnowledgeBaseBuildTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) SetData(v *GetKnowledgeBaseBuildTaskResultResponseBodyData) *GetKnowledgeBaseBuildTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) SetMessage(v string) *GetKnowledgeBaseBuildTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) SetRequestId(v string) *GetKnowledgeBaseBuildTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) SetSuccess(v bool) *GetKnowledgeBaseBuildTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBaseBuildTaskResultResponseBodyData struct {
	// The temporary Object Storage Service (OSS) download URL issued to the current caller when this operation is called. The URL is not stored.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The validity period of the download URL, in seconds. Request a new URL after it expires.
	//
	// example:
	//
	// 300
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// The ID of the knowledge base to which the result belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the task instance that produced the result.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) SetDownloadUrl(v string) *GetKnowledgeBaseBuildTaskResultResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) SetExpiresIn(v int32) *GetKnowledgeBaseBuildTaskResultResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskResultResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskResultResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
