// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryTraceM3u8JobListResponseBodyData) *QueryTraceM3u8JobListResponseBody
	GetData() []*QueryTraceM3u8JobListResponseBodyData
	SetMessage(v string) *QueryTraceM3u8JobListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTraceM3u8JobListResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryTraceM3u8JobListResponseBody
	GetStatusCode() *int64
}

type QueryTraceM3u8JobListResponseBody struct {
	Data []*QueryTraceM3u8JobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceM3u8JobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobListResponseBody) GetData() []*QueryTraceM3u8JobListResponseBodyData {
	return s.Data
}

func (s *QueryTraceM3u8JobListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTraceM3u8JobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTraceM3u8JobListResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryTraceM3u8JobListResponseBody) SetData(v []*QueryTraceM3u8JobListResponseBodyData) *QueryTraceM3u8JobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceM3u8JobListResponseBody) SetMessage(v string) *QueryTraceM3u8JobListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBody) SetRequestId(v string) *QueryTraceM3u8JobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBody) SetStatusCode(v int64) *QueryTraceM3u8JobListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTraceM3u8JobListResponseBodyData struct {
	// example:
	//
	// 1627357322
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1627357322
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId  *string                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output *QueryTraceM3u8JobListResponseBodyDataOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// test
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// example:
	//
	// ****437bd2b105d07b12a9a82****
	TraceMediaId *string `json:"TraceMediaId,omitempty" xml:"TraceMediaId,omitempty"`
	// example:
	//
	// 112
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// example:
	//
	// 1346693276****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceM3u8JobListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetOutput() *QueryTraceM3u8JobListResponseBodyDataOutput {
	return s.Output
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetTrace() *string {
	return s.Trace
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryTraceM3u8JobListResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetGmtCreate(v int64) *QueryTraceM3u8JobListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetGmtModified(v int64) *QueryTraceM3u8JobListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetJobId(v string) *QueryTraceM3u8JobListResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetOutput(v *QueryTraceM3u8JobListResponseBodyDataOutput) *QueryTraceM3u8JobListResponseBodyData {
	s.Output = v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetStatus(v string) *QueryTraceM3u8JobListResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetTrace(v string) *QueryTraceM3u8JobListResponseBodyData {
	s.Trace = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetTraceMediaId(v string) *QueryTraceM3u8JobListResponseBodyData {
	s.TraceMediaId = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetUserData(v string) *QueryTraceM3u8JobListResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) SetUserId(v int64) *QueryTraceM3u8JobListResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTraceM3u8JobListResponseBodyDataOutput struct {
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryTraceM3u8JobListResponseBodyDataOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobListResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobListResponseBodyDataOutput) GetMedia() *string {
	return s.Media
}

func (s *QueryTraceM3u8JobListResponseBodyDataOutput) GetType() *string {
	return s.Type
}

func (s *QueryTraceM3u8JobListResponseBodyDataOutput) SetMedia(v string) *QueryTraceM3u8JobListResponseBodyDataOutput {
	s.Media = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyDataOutput) SetType(v string) *QueryTraceM3u8JobListResponseBodyDataOutput {
	s.Type = &v
	return s
}

func (s *QueryTraceM3u8JobListResponseBodyDataOutput) Validate() error {
	return dara.Validate(s)
}
