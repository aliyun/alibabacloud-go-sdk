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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 目录 ID
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 成功入队重试的数量
	//
	// example:
	//
	// 1
	EnqueuedCount *int64 `json:"enqueuedCount,omitempty" xml:"enqueuedCount,omitempty"`
	// enqueuedIds
	//
	// example:
	//
	// string_value
	EnqueuedIds []*string `json:"enqueuedIds,omitempty" xml:"enqueuedIds,omitempty" type:"Repeated"`
	// 目录下失败资源总数
	//
	// example:
	//
	// 1
	FailedCount   *int64                                                  `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	FailedSources []*RetryDirectoryFailedSourcesResponseBodyFailedSources `json:"failedSources,omitempty" xml:"failedSources,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 跳过（非 FAILED 或无权限）的数量
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
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 数据源类型
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
