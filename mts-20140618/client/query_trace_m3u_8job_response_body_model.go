// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryTraceM3u8JobResponseBodyData) *QueryTraceM3u8JobResponseBody
	GetData() []*QueryTraceM3u8JobResponseBodyData
	SetMessage(v string) *QueryTraceM3u8JobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTraceM3u8JobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *QueryTraceM3u8JobResponseBody
	GetStatusCode() *int64
}

type QueryTraceM3u8JobResponseBody struct {
	Data []*QueryTraceM3u8JobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5CA6E020-4102-4FFF-AA56-5ED7ECD8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceM3u8JobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobResponseBody) GetData() []*QueryTraceM3u8JobResponseBodyData {
	return s.Data
}

func (s *QueryTraceM3u8JobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTraceM3u8JobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTraceM3u8JobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *QueryTraceM3u8JobResponseBody) SetData(v []*QueryTraceM3u8JobResponseBodyData) *QueryTraceM3u8JobResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceM3u8JobResponseBody) SetMessage(v string) *QueryTraceM3u8JobResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBody) SetRequestId(v string) *QueryTraceM3u8JobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBody) SetStatusCode(v int64) *QueryTraceM3u8JobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBody) Validate() error {
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

type QueryTraceM3u8JobResponseBodyData struct {
	// example:
	//
	// 1627357322
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1627357327
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 平头哥半导体(上海)
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// example:
	//
	// 13466****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceM3u8JobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryTraceM3u8JobResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryTraceM3u8JobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceM3u8JobResponseBodyData) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryTraceM3u8JobResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *QueryTraceM3u8JobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryTraceM3u8JobResponseBodyData) GetTrace() *string {
	return s.Trace
}

func (s *QueryTraceM3u8JobResponseBodyData) GetUserData() *string {
	return s.UserData
}

func (s *QueryTraceM3u8JobResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryTraceM3u8JobResponseBodyData) SetGmtCreate(v int64) *QueryTraceM3u8JobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetGmtModified(v int64) *QueryTraceM3u8JobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetJobId(v string) *QueryTraceM3u8JobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetMediaId(v string) *QueryTraceM3u8JobResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetOutput(v string) *QueryTraceM3u8JobResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetStatus(v string) *QueryTraceM3u8JobResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetTrace(v string) *QueryTraceM3u8JobResponseBodyData {
	s.Trace = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetUserData(v string) *QueryTraceM3u8JobResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) SetUserId(v int64) *QueryTraceM3u8JobResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryTraceM3u8JobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
