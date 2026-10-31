// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStreamTaskManagerLogsResponseBody
	GetCode() *string
	SetData(v []*ListStreamTaskManagerLogsResponseBodyData) *ListStreamTaskManagerLogsResponseBody
	GetData() []*ListStreamTaskManagerLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListStreamTaskManagerLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStreamTaskManagerLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStreamTaskManagerLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStreamTaskManagerLogsResponseBody
	GetSuccess() *bool
}

type ListStreamTaskManagerLogsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of log files for the specified TaskManager.
	Data []*ListStreamTaskManagerLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListStreamTaskManagerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagerLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStreamTaskManagerLogsResponseBody) GetData() []*ListStreamTaskManagerLogsResponseBodyData {
	return s.Data
}

func (s *ListStreamTaskManagerLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStreamTaskManagerLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStreamTaskManagerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamTaskManagerLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStreamTaskManagerLogsResponseBody) SetCode(v string) *ListStreamTaskManagerLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) SetData(v []*ListStreamTaskManagerLogsResponseBodyData) *ListStreamTaskManagerLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) SetHttpStatusCode(v int32) *ListStreamTaskManagerLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) SetMessage(v string) *ListStreamTaskManagerLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) SetRequestId(v string) *ListStreamTaskManagerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) SetSuccess(v bool) *ListStreamTaskManagerLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStreamTaskManagerLogsResponseBodyData struct {
	// The log file name. This name can be used as the logName input parameter for the logContent operation.
	//
	// example:
	//
	// flink--standalonesession-0-dataphin-session-7db678957dc449aab0086b931c9f8218-jobmanager.log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The last modified time in milliseconds.
	//
	// example:
	//
	// 1792033500000
	Mtime *int64 `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 1048576
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListStreamTaskManagerLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagerLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagerLogsResponseBodyData) GetLogName() *string {
	return s.LogName
}

func (s *ListStreamTaskManagerLogsResponseBodyData) GetMtime() *int64 {
	return s.Mtime
}

func (s *ListStreamTaskManagerLogsResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListStreamTaskManagerLogsResponseBodyData) SetLogName(v string) *ListStreamTaskManagerLogsResponseBodyData {
	s.LogName = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBodyData) SetMtime(v int64) *ListStreamTaskManagerLogsResponseBodyData {
	s.Mtime = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBodyData) SetSize(v int64) *ListStreamTaskManagerLogsResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListStreamTaskManagerLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
