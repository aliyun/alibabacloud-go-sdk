// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDirectoryFailedSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryDirectoryFailedSourcesResponseBody
	GetCode() *string
	SetDirectoryId(v string) *RetryDirectoryFailedSourcesResponseBody
	GetDirectoryId() *string
	SetEnqueuedCount(v int64) *RetryDirectoryFailedSourcesResponseBody
	GetEnqueuedCount() *int64
	SetEnqueuedIds(v []*string) *RetryDirectoryFailedSourcesResponseBody
	GetEnqueuedIds() []*string
	SetFailedCount(v int64) *RetryDirectoryFailedSourcesResponseBody
	GetFailedCount() *int64
	SetFailedSources(v []*RetryDirectoryFailedSourcesResponseBodyFailedSources) *RetryDirectoryFailedSourcesResponseBody
	GetFailedSources() []*RetryDirectoryFailedSourcesResponseBodyFailedSources
	SetMessage(v string) *RetryDirectoryFailedSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryDirectoryFailedSourcesResponseBody
	GetRequestId() *string
	SetSkippedCount(v int64) *RetryDirectoryFailedSourcesResponseBody
	GetSkippedCount() *int64
}

type RetryDirectoryFailedSourcesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the enterprise knowledge base folder. Failed resources in subfolders are recursively included.
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
	// The list of source IDs that have been re-enqueued.
	//
	// example:
	//
	// string_value
	EnqueuedIds []*string `json:"enqueuedIds,omitempty" xml:"enqueuedIds,omitempty" type:"Repeated"`
	// The total number of failed resources in the folder.
	//
	// example:
	//
	// 1
	FailedCount *int64 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// The list of failed data source IDs.
	FailedSources []*RetryDirectoryFailedSourcesResponseBodyFailedSources `json:"failedSources,omitempty" xml:"failedSources,omitempty" type:"Repeated"`
	// The description of the status code.
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

func (s RetryDirectoryFailedSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryDirectoryFailedSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetEnqueuedCount() *int64 {
	return s.EnqueuedCount
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetEnqueuedIds() []*string {
	return s.EnqueuedIds
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetFailedSources() []*RetryDirectoryFailedSourcesResponseBodyFailedSources {
	return s.FailedSources
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryDirectoryFailedSourcesResponseBody) GetSkippedCount() *int64 {
	return s.SkippedCount
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetCode(v string) *RetryDirectoryFailedSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetDirectoryId(v string) *RetryDirectoryFailedSourcesResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetEnqueuedCount(v int64) *RetryDirectoryFailedSourcesResponseBody {
	s.EnqueuedCount = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetEnqueuedIds(v []*string) *RetryDirectoryFailedSourcesResponseBody {
	s.EnqueuedIds = v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetFailedCount(v int64) *RetryDirectoryFailedSourcesResponseBody {
	s.FailedCount = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetFailedSources(v []*RetryDirectoryFailedSourcesResponseBodyFailedSources) *RetryDirectoryFailedSourcesResponseBody {
	s.FailedSources = v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetMessage(v string) *RetryDirectoryFailedSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetRequestId(v string) *RetryDirectoryFailedSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) SetSkippedCount(v int64) *RetryDirectoryFailedSourcesResponseBody {
	s.SkippedCount = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBody) Validate() error {
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

type RetryDirectoryFailedSourcesResponseBodyFailedSources struct {
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

func (s RetryDirectoryFailedSourcesResponseBodyFailedSources) String() string {
	return dara.Prettify(s)
}

func (s RetryDirectoryFailedSourcesResponseBodyFailedSources) GoString() string {
	return s.String()
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) GetName() *string {
	return s.Name
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) GetSourceId() *string {
	return s.SourceId
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) GetSourceType() *string {
	return s.SourceType
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) SetName(v string) *RetryDirectoryFailedSourcesResponseBodyFailedSources {
	s.Name = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) SetSourceId(v string) *RetryDirectoryFailedSourcesResponseBodyFailedSources {
	s.SourceId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) SetSourceType(v string) *RetryDirectoryFailedSourcesResponseBodyFailedSources {
	s.SourceType = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponseBodyFailedSources) Validate() error {
	return dara.Validate(s)
}
