// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryCopyrightJobResponseBodyData) *QueryCopyrightJobResponseBody
	GetData() []*QueryCopyrightJobResponseBodyData
	SetMessage(v string) *QueryCopyrightJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCopyrightJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryCopyrightJobResponseBody
	GetStatusCode() *int64
}

type QueryCopyrightJobResponseBody struct {
	Data []*QueryCopyrightJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5CA6E020-4102-4FFF-AA56-5ED7ECD811A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobResponseBody) GetData() []*QueryCopyrightJobResponseBodyData {
	return s.Data
}

func (s *QueryCopyrightJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCopyrightJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryCopyrightJobResponseBody) SetData(v []*QueryCopyrightJobResponseBodyData) *QueryCopyrightJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightJobResponseBody) SetMessage(v string) *QueryCopyrightJobResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCopyrightJobResponseBody) SetRequestId(v string) *QueryCopyrightJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCopyrightJobResponseBody) SetStatusCode(v int64) *QueryCopyrightJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightJobResponseBody) Validate() error {
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

type QueryCopyrightJobResponseBodyData struct {
	// example:
	//
	// http://callbacktest.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// 1627357322
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1627357328
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {"Bucket":"ivison-test","Location":"oss-cn-shanghai","Object":"gambling.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// bfb786c639894f4d80648792021e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 平头哥半导体(上海)
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// {"Bucket":"ivison-test","Location":"oss-cn-shanghai","Object":"out.mp4"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
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

func (s QueryCopyrightJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobResponseBodyData) GetCallback() *string {
	return s.Callback
}

func (s *QueryCopyrightJobResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryCopyrightJobResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryCopyrightJobResponseBodyData) GetInput() *string {
	return s.Input
}

func (s *QueryCopyrightJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryCopyrightJobResponseBodyData) GetLevel() *int64 {
	return s.Level
}

func (s *QueryCopyrightJobResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *QueryCopyrightJobResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *QueryCopyrightJobResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *QueryCopyrightJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryCopyrightJobResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryCopyrightJobResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryCopyrightJobResponseBodyData) SetCallback(v string) *QueryCopyrightJobResponseBodyData {
	s.Callback = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetGmtCreate(v int64) *QueryCopyrightJobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetGmtModified(v int64) *QueryCopyrightJobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetInput(v string) *QueryCopyrightJobResponseBodyData {
	s.Input = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetJobId(v string) *QueryCopyrightJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetLevel(v int64) *QueryCopyrightJobResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetMessage(v string) *QueryCopyrightJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetOutput(v string) *QueryCopyrightJobResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetResult(v string) *QueryCopyrightJobResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetStatus(v string) *QueryCopyrightJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetUserData(v string) *QueryCopyrightJobResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) SetUserId(v int64) *QueryCopyrightJobResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryCopyrightJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
