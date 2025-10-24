// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryTraceAbJobResponseBodyData) *QueryTraceAbJobResponseBody
	GetData() []*QueryTraceAbJobResponseBodyData
	SetMessage(v string) *QueryTraceAbJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTraceAbJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryTraceAbJobResponseBody
	GetStatusCode() *int64
}

type QueryTraceAbJobResponseBody struct {
	Data []*QueryTraceAbJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 338CA33A-AE83-5DF4-B6F2-C6D3ED8143F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceAbJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobResponseBody) GetData() []*QueryTraceAbJobResponseBodyData {
	return s.Data
}

func (s *QueryTraceAbJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTraceAbJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTraceAbJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryTraceAbJobResponseBody) SetData(v []*QueryTraceAbJobResponseBodyData) *QueryTraceAbJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceAbJobResponseBody) SetMessage(v string) *QueryTraceAbJobResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceAbJobResponseBody) SetRequestId(v string) *QueryTraceAbJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTraceAbJobResponseBody) SetStatusCode(v int64) *QueryTraceAbJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceAbJobResponseBody) Validate() error {
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

type QueryTraceAbJobResponseBodyData struct {
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
	// 1627357325
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {"Bucket":"ivison-test","Location":"oss-cn-shanghai","Object":"test.mp4"}
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
	// 437bd2b516ffda105d07b12a9a82****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
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
	// 13466****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceAbJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobResponseBodyData) GetCallback() *string {
	return s.Callback
}

func (s *QueryTraceAbJobResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTraceAbJobResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTraceAbJobResponseBodyData) GetInput() *string {
	return s.Input
}

func (s *QueryTraceAbJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceAbJobResponseBodyData) GetLevel() *int64 {
	return s.Level
}

func (s *QueryTraceAbJobResponseBodyData) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryTraceAbJobResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *QueryTraceAbJobResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *QueryTraceAbJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryTraceAbJobResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryTraceAbJobResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryTraceAbJobResponseBodyData) SetCallback(v string) *QueryTraceAbJobResponseBodyData {
	s.Callback = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetGmtCreate(v int64) *QueryTraceAbJobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetGmtModified(v int64) *QueryTraceAbJobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetInput(v string) *QueryTraceAbJobResponseBodyData {
	s.Input = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetJobId(v string) *QueryTraceAbJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetLevel(v int64) *QueryTraceAbJobResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetMediaId(v string) *QueryTraceAbJobResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetOutput(v string) *QueryTraceAbJobResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetResult(v string) *QueryTraceAbJobResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetStatus(v string) *QueryTraceAbJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetUserData(v string) *QueryTraceAbJobResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) SetUserId(v int64) *QueryTraceAbJobResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryTraceAbJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
