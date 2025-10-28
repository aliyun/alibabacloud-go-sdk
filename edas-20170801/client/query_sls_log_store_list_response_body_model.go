// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlsLogStoreListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QuerySlsLogStoreListResponseBody
	GetCode() *int32
	SetMessage(v string) *QuerySlsLogStoreListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySlsLogStoreListResponseBody
	GetRequestId() *string
	SetResult(v []*QuerySlsLogStoreListResponseBodyResult) *QuerySlsLogStoreListResponseBody
	GetResult() []*QuerySlsLogStoreListResponseBodyResult
	SetTotalSize(v int32) *QuerySlsLogStoreListResponseBody
	GetTotalSize() *int32
}

type QuerySlsLogStoreListResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// d6834ee9-5045-*************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of Log Service for the application.
	Result []*QuerySlsLogStoreListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The number of log sources configured for the application.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QuerySlsLogStoreListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySlsLogStoreListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QuerySlsLogStoreListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySlsLogStoreListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySlsLogStoreListResponseBody) GetResult() []*QuerySlsLogStoreListResponseBodyResult {
	return s.Result
}

func (s *QuerySlsLogStoreListResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *QuerySlsLogStoreListResponseBody) SetCode(v int32) *QuerySlsLogStoreListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetMessage(v string) *QuerySlsLogStoreListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetRequestId(v string) *QuerySlsLogStoreListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetResult(v []*QuerySlsLogStoreListResponseBodyResult) *QuerySlsLogStoreListResponseBody {
	s.Result = v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetTotalSize(v int32) *QuerySlsLogStoreListResponseBody {
	s.TotalSize = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySlsLogStoreListResponseBodyResult struct {
	// The type of the logging service.
	//
	// example:
	//
	// Log Service
	ConsumerSide *string `json:"ConsumerSide,omitempty" xml:"ConsumerSide,omitempty"`
	// The time when the logging service was created.
	//
	// example:
	//
	// 2020-05-18 22:08:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the logging service.
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// thisisatestlogstore
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// k8s-log-c846f28edbd1d4c6aa9d78c0e********
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The source of logs. Valid values:
	//
	// 	- Standard output: stdout.log
	//
	// 	- File log: the directory that stores logs
	//
	// example:
	//
	// /var/log/*
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s QuerySlsLogStoreListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QuerySlsLogStoreListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetConsumerSide() *string {
	return s.ConsumerSide
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetLink() *string {
	return s.Link
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetLogstore() *string {
	return s.Logstore
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetProject() *string {
	return s.Project
}

func (s *QuerySlsLogStoreListResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetConsumerSide(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.ConsumerSide = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetCreateTime(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetLink(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Link = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetLogstore(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Logstore = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetProject(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Project = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetSource(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
