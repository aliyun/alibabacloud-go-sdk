// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryCopyrightJobListResponseBodyData) *QueryCopyrightJobListResponseBody
	GetData() []*QueryCopyrightJobListResponseBodyData
	SetMessage(v string) *QueryCopyrightJobListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCopyrightJobListResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryCopyrightJobListResponseBody
	GetStatusCode() *int64
}

type QueryCopyrightJobListResponseBody struct {
	Data []*QueryCopyrightJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******36-3C1E-4417-BDB2-1E034F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListResponseBody) GetData() []*QueryCopyrightJobListResponseBodyData {
	return s.Data
}

func (s *QueryCopyrightJobListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCopyrightJobListResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryCopyrightJobListResponseBody) SetData(v []*QueryCopyrightJobListResponseBodyData) *QueryCopyrightJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightJobListResponseBody) SetMessage(v string) *QueryCopyrightJobListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCopyrightJobListResponseBody) SetRequestId(v string) *QueryCopyrightJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCopyrightJobListResponseBody) SetStatusCode(v int64) *QueryCopyrightJobListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightJobListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCopyrightJobListResponseBodyData struct {
	// example:
	//
	// 1627357322
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1627357322
	GmtModified *int64                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Input       *QueryCopyrightJobListResponseBodyDataInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// bfb786c639894f4d80648792021****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// test
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *QueryCopyrightJobListResponseBodyDataOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
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
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// example:
	//
	// 1346693***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCopyrightJobListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryCopyrightJobListResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryCopyrightJobListResponseBodyData) GetInput() *QueryCopyrightJobListResponseBodyDataInput {
	return s.Input
}

func (s *QueryCopyrightJobListResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryCopyrightJobListResponseBodyData) GetLevel() *int64 {
	return s.Level
}

func (s *QueryCopyrightJobListResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightJobListResponseBodyData) GetOutput() *QueryCopyrightJobListResponseBodyDataOutput {
	return s.Output
}

func (s *QueryCopyrightJobListResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *QueryCopyrightJobListResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryCopyrightJobListResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryCopyrightJobListResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryCopyrightJobListResponseBodyData) SetGmtCreate(v int64) *QueryCopyrightJobListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetGmtModified(v int64) *QueryCopyrightJobListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetInput(v *QueryCopyrightJobListResponseBodyDataInput) *QueryCopyrightJobListResponseBodyData {
	s.Input = v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetJobId(v string) *QueryCopyrightJobListResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetLevel(v int64) *QueryCopyrightJobListResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetMessage(v string) *QueryCopyrightJobListResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetOutput(v *QueryCopyrightJobListResponseBodyDataOutput) *QueryCopyrightJobListResponseBodyData {
	s.Output = v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetResult(v string) *QueryCopyrightJobListResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetStatus(v string) *QueryCopyrightJobListResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetUserData(v string) *QueryCopyrightJobListResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) SetUserId(v int64) *QueryCopyrightJobListResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryCopyrightJobListResponseBodyDataInput struct {
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryCopyrightJobListResponseBodyDataInput) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListResponseBodyDataInput) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListResponseBodyDataInput) GetMedia() *string {
	return s.Media
}

func (s *QueryCopyrightJobListResponseBodyDataInput) GetType() *string {
	return s.Type
}

func (s *QueryCopyrightJobListResponseBodyDataInput) SetMedia(v string) *QueryCopyrightJobListResponseBodyDataInput {
	s.Media = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyDataInput) SetType(v string) *QueryCopyrightJobListResponseBodyDataInput {
	s.Type = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyDataInput) Validate() error {
	return dara.Validate(s)
}

type QueryCopyrightJobListResponseBodyDataOutput struct {
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryCopyrightJobListResponseBodyDataOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListResponseBodyDataOutput) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListResponseBodyDataOutput) GetMedia() *string {
	return s.Media
}

func (s *QueryCopyrightJobListResponseBodyDataOutput) GetType() *string {
	return s.Type
}

func (s *QueryCopyrightJobListResponseBodyDataOutput) SetMedia(v string) *QueryCopyrightJobListResponseBodyDataOutput {
	s.Media = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyDataOutput) SetType(v string) *QueryCopyrightJobListResponseBodyDataOutput {
	s.Type = &v
	return s
}

func (s *QueryCopyrightJobListResponseBodyDataOutput) Validate() error {
	return dara.Validate(s)
}
