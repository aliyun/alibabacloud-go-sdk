// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *RetryKnowledgeBaseFilesResponseBody
	GetFailedCount() *int32
	SetItems(v []*RetryKnowledgeBaseFilesResponseBodyItems) *RetryKnowledgeBaseFilesResponseBody
	GetItems() []*RetryKnowledgeBaseFilesResponseBodyItems
	SetRequestId(v string) *RetryKnowledgeBaseFilesResponseBody
	GetRequestId() *string
	SetSucceededCount(v int32) *RetryKnowledgeBaseFilesResponseBody
	GetSucceededCount() *int32
	SetTotalCount(v int32) *RetryKnowledgeBaseFilesResponseBody
	GetTotalCount() *int32
}

type RetryKnowledgeBaseFilesResponseBody struct {
	// The number of failed retries.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The list of retry results.
	Items []*RetryKnowledgeBaseFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of successful retries.
	//
	// example:
	//
	// 1
	SucceededCount *int32 `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RetryKnowledgeBaseFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFilesResponseBody) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFilesResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *RetryKnowledgeBaseFilesResponseBody) GetItems() []*RetryKnowledgeBaseFilesResponseBodyItems {
	return s.Items
}

func (s *RetryKnowledgeBaseFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryKnowledgeBaseFilesResponseBody) GetSucceededCount() *int32 {
	return s.SucceededCount
}

func (s *RetryKnowledgeBaseFilesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RetryKnowledgeBaseFilesResponseBody) SetFailedCount(v int32) *RetryKnowledgeBaseFilesResponseBody {
	s.FailedCount = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBody) SetItems(v []*RetryKnowledgeBaseFilesResponseBodyItems) *RetryKnowledgeBaseFilesResponseBody {
	s.Items = v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBody) SetRequestId(v string) *RetryKnowledgeBaseFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBody) SetSucceededCount(v int32) *RetryKnowledgeBaseFilesResponseBody {
	s.SucceededCount = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBody) SetTotalCount(v int32) *RetryKnowledgeBaseFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetryKnowledgeBaseFilesResponseBodyItems struct {
	// The error code.
	//
	// example:
	//
	// AclProjectionNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// No active ACL projection exists for the IM document
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 91b97b71-xxxx-xxxx-xxxx-33c6a6341cdc
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The request result. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryKnowledgeBaseFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) GetFileId() *string {
	return s.FileId
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) GetSuccess() *bool {
	return s.Success
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) SetErrorCode(v string) *RetryKnowledgeBaseFilesResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) SetErrorMessage(v string) *RetryKnowledgeBaseFilesResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) SetFileId(v string) *RetryKnowledgeBaseFilesResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) SetSuccess(v bool) *RetryKnowledgeBaseFilesResponseBodyItems {
	s.Success = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
