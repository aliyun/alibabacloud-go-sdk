// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFailedSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetCode() *string
	SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetDirectoryId() *string
	SetEnqueuedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetEnqueuedCount() *int64
	SetEnqueuedIds(v []*string) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetEnqueuedIds() []*string
	SetFailedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetFailedCount() *int64
	SetFailedSources(v []*RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetFailedSources() []*RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources
	SetMessage(v string) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetRequestId() *string
	SetSkippedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody
	GetSkippedCount() *int64
}

type RetryKnowledgeBaseFailedSourcesResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The enterprise knowledge base directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The number of resources successfully enqueued for retry.
	//
	// example:
	//
	// 1
	EnqueuedCount *int64 `json:"enqueuedCount,omitempty" xml:"enqueuedCount,omitempty"`
	// enqueuedIds。
	//
	// example:
	//
	// string_value
	EnqueuedIds []*string `json:"enqueuedIds,omitempty" xml:"enqueuedIds,omitempty" type:"Repeated"`
	// The total number of failed resources under the directory.
	//
	// example:
	//
	// 1
	FailedCount *int64 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// The list of failed data sources.
	FailedSources []*RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources `json:"failedSources,omitempty" xml:"failedSources,omitempty" type:"Repeated"`
	// The error description. This value is empty on success.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of resources skipped because they are not in FAILED status.
	//
	// example:
	//
	// 1
	SkippedCount *int64 `json:"skippedCount,omitempty" xml:"skippedCount,omitempty"`
}

func (s RetryKnowledgeBaseFailedSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFailedSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetEnqueuedCount() *int64 {
	return s.EnqueuedCount
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetEnqueuedIds() []*string {
	return s.EnqueuedIds
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetFailedSources() []*RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources {
	return s.FailedSources
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) GetSkippedCount() *int64 {
	return s.SkippedCount
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetCode(v string) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetEnqueuedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.EnqueuedCount = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetEnqueuedIds(v []*string) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.EnqueuedIds = v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetFailedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.FailedCount = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetFailedSources(v []*RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.FailedSources = v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetMessage(v string) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetRequestId(v string) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) SetSkippedCount(v int64) *RetryKnowledgeBaseFailedSourcesResponseBody {
	s.SkippedCount = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBody) Validate() error {
	if s.FailedSources != nil {
		for _, item := range s.FailedSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources struct {
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source type.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) GetName() *string {
	return s.Name
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) GetSourceId() *string {
	return s.SourceId
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) GetSourceType() *string {
	return s.SourceType
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) SetName(v string) *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources {
	s.Name = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) SetSourceId(v string) *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources {
	s.SourceId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) SetSourceType(v string) *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources {
	s.SourceType = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponseBodyFailedSources) Validate() error {
	return dara.Validate(s)
}
