// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryTraceAbJobListResponseBodyData) *QueryTraceAbJobListResponseBody
	GetData() []*QueryTraceAbJobListResponseBodyData
	SetMessage(v string) *QueryTraceAbJobListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTraceAbJobListResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryTraceAbJobListResponseBody
	GetStatusCode() *int64
}

type QueryTraceAbJobListResponseBody struct {
	Data []*QueryTraceAbJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QueryTraceAbJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListResponseBody) GetData() []*QueryTraceAbJobListResponseBodyData {
	return s.Data
}

func (s *QueryTraceAbJobListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTraceAbJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTraceAbJobListResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryTraceAbJobListResponseBody) SetData(v []*QueryTraceAbJobListResponseBodyData) *QueryTraceAbJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceAbJobListResponseBody) SetMessage(v string) *QueryTraceAbJobListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceAbJobListResponseBody) SetRequestId(v string) *QueryTraceAbJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTraceAbJobListResponseBody) SetStatusCode(v int64) *QueryTraceAbJobListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceAbJobListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTraceAbJobListResponseBodyData struct {
	// example:
	//
	// 1627357322
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1627357322
	GmtModified *int64                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Input       *QueryTraceAbJobListResponseBodyDataInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// bfb786c639894f4d80648792021eff90
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2
	Level  *int64                                     `json:"Level,omitempty" xml:"Level,omitempty"`
	Output *QueryTraceAbJobListResponseBodyDataOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// {"Code":"success","Message":"ok"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ****437bd2b51105d07b12a9****
	TraceMediaId *string `json:"TraceMediaId,omitempty" xml:"TraceMediaId,omitempty"`
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// example:
	//
	// 13466932****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceAbJobListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTraceAbJobListResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTraceAbJobListResponseBodyData) GetInput() *QueryTraceAbJobListResponseBodyDataInput {
	return s.Input
}

func (s *QueryTraceAbJobListResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceAbJobListResponseBodyData) GetLevel() *int64 {
	return s.Level
}

func (s *QueryTraceAbJobListResponseBodyData) GetOutput() *QueryTraceAbJobListResponseBodyDataOutput {
	return s.Output
}

func (s *QueryTraceAbJobListResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *QueryTraceAbJobListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryTraceAbJobListResponseBodyData) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *QueryTraceAbJobListResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryTraceAbJobListResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryTraceAbJobListResponseBodyData) SetGmtCreate(v int64) *QueryTraceAbJobListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetGmtModified(v int64) *QueryTraceAbJobListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetInput(v *QueryTraceAbJobListResponseBodyDataInput) *QueryTraceAbJobListResponseBodyData {
	s.Input = v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetJobId(v string) *QueryTraceAbJobListResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetLevel(v int64) *QueryTraceAbJobListResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetOutput(v *QueryTraceAbJobListResponseBodyDataOutput) *QueryTraceAbJobListResponseBodyData {
	s.Output = v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetResult(v string) *QueryTraceAbJobListResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetStatus(v string) *QueryTraceAbJobListResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetTraceMediaId(v string) *QueryTraceAbJobListResponseBodyData {
	s.TraceMediaId = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetUserData(v string) *QueryTraceAbJobListResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) SetUserId(v int64) *QueryTraceAbJobListResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTraceAbJobListResponseBodyDataInput struct {
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryTraceAbJobListResponseBodyDataInput) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListResponseBodyDataInput) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListResponseBodyDataInput) GetMedia() *string {
	return s.Media
}

func (s *QueryTraceAbJobListResponseBodyDataInput) GetType() *string {
	return s.Type
}

func (s *QueryTraceAbJobListResponseBodyDataInput) SetMedia(v string) *QueryTraceAbJobListResponseBodyDataInput {
	s.Media = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyDataInput) SetType(v string) *QueryTraceAbJobListResponseBodyDataInput {
	s.Type = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyDataInput) Validate() error {
	return dara.Validate(s)
}

type QueryTraceAbJobListResponseBodyDataOutput struct {
	// example:
	//
	// oss://bucket/dir/
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryTraceAbJobListResponseBodyDataOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListResponseBodyDataOutput) GetMedia() *string {
	return s.Media
}

func (s *QueryTraceAbJobListResponseBodyDataOutput) GetType() *string {
	return s.Type
}

func (s *QueryTraceAbJobListResponseBodyDataOutput) SetMedia(v string) *QueryTraceAbJobListResponseBodyDataOutput {
	s.Media = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyDataOutput) SetType(v string) *QueryTraceAbJobListResponseBodyDataOutput {
	s.Type = &v
	return s
}

func (s *QueryTraceAbJobListResponseBodyDataOutput) Validate() error {
	return dara.Validate(s)
}
