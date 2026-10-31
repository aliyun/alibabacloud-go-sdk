// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamJobManagerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStreamJobManagerLogsResponseBody
	GetCode() *string
	SetData(v []*ListStreamJobManagerLogsResponseBodyData) *ListStreamJobManagerLogsResponseBody
	GetData() []*ListStreamJobManagerLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListStreamJobManagerLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStreamJobManagerLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStreamJobManagerLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStreamJobManagerLogsResponseBody
	GetSuccess() *bool
}

type ListStreamJobManagerLogsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The log file entries for JobManager or TaskManager. The logName parameter can be used to retrieve the log content.
	Data []*ListStreamJobManagerLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
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

func (s ListStreamJobManagerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamJobManagerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamJobManagerLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStreamJobManagerLogsResponseBody) GetData() []*ListStreamJobManagerLogsResponseBodyData {
	return s.Data
}

func (s *ListStreamJobManagerLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStreamJobManagerLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStreamJobManagerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamJobManagerLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStreamJobManagerLogsResponseBody) SetCode(v string) *ListStreamJobManagerLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) SetData(v []*ListStreamJobManagerLogsResponseBodyData) *ListStreamJobManagerLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) SetHttpStatusCode(v int32) *ListStreamJobManagerLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) SetMessage(v string) *ListStreamJobManagerLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) SetRequestId(v string) *ListStreamJobManagerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) SetSuccess(v bool) *ListStreamJobManagerLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBody) Validate() error {
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

type ListStreamJobManagerLogsResponseBodyData struct {
	// The original name of the log file, which corresponds to the LogName input parameter of GetStreamLogContent. Pass this value exactly as returned.
	//
	// example:
	//
	// flink--standalonesession-0-dataphin-session-7db678957dc449aab0086b931c9f8218-jobmanager.log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The last modified time of the log file, as a millisecond timestamp since the Unix epoch.
	//
	// example:
	//
	// 1792033500000
	Mtime *int64 `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The size of the log file, in bytes.
	//
	// example:
	//
	// 1048576
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListStreamJobManagerLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStreamJobManagerLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStreamJobManagerLogsResponseBodyData) GetLogName() *string {
	return s.LogName
}

func (s *ListStreamJobManagerLogsResponseBodyData) GetMtime() *int64 {
	return s.Mtime
}

func (s *ListStreamJobManagerLogsResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListStreamJobManagerLogsResponseBodyData) SetLogName(v string) *ListStreamJobManagerLogsResponseBodyData {
	s.LogName = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBodyData) SetMtime(v int64) *ListStreamJobManagerLogsResponseBodyData {
	s.Mtime = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBodyData) SetSize(v int64) *ListStreamJobManagerLogsResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListStreamJobManagerLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
