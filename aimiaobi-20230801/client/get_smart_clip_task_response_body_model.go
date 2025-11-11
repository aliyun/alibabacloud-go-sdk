// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartClipTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmartClipTaskResponseBody
	GetCode() *string
	SetData(v *GetSmartClipTaskResponseBodyData) *GetSmartClipTaskResponseBody
	GetData() *GetSmartClipTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetSmartClipTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSmartClipTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmartClipTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSmartClipTaskResponseBody
	GetSuccess() *bool
}

type GetSmartClipTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSmartClipTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSmartClipTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmartClipTaskResponseBody) GetData() *GetSmartClipTaskResponseBodyData {
	return s.Data
}

func (s *GetSmartClipTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSmartClipTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmartClipTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartClipTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSmartClipTaskResponseBody) SetCode(v string) *GetSmartClipTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmartClipTaskResponseBody) SetData(v *GetSmartClipTaskResponseBodyData) *GetSmartClipTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetSmartClipTaskResponseBody) SetHttpStatusCode(v int32) *GetSmartClipTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSmartClipTaskResponseBody) SetMessage(v string) *GetSmartClipTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmartClipTaskResponseBody) SetRequestId(v string) *GetSmartClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartClipTaskResponseBody) SetSuccess(v bool) *GetSmartClipTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetSmartClipTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmartClipTaskResponseBodyData struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// RUNNING
	Status  *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubJobs []*GetSmartClipTaskResponseBodyDataSubJobs `json:"SubJobs,omitempty" xml:"SubJobs,omitempty" type:"Repeated"`
}

func (s GetSmartClipTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSmartClipTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSmartClipTaskResponseBodyData) GetSubJobs() []*GetSmartClipTaskResponseBodyDataSubJobs {
	return s.SubJobs
}

func (s *GetSmartClipTaskResponseBodyData) SetErrorMessage(v string) *GetSmartClipTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyData) SetStatus(v string) *GetSmartClipTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyData) SetSubJobs(v []*GetSmartClipTaskResponseBodyDataSubJobs) *GetSmartClipTaskResponseBodyData {
	s.SubJobs = v
	return s
}

func (s *GetSmartClipTaskResponseBodyData) Validate() error {
	if s.SubJobs != nil {
		for _, item := range s.SubJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmartClipTaskResponseBodyDataSubJobs struct {
	// example:
	//
	// x\"x\"x\"x
	ErrorMessage *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileAttr     *GetSmartClipTaskResponseBodyDataSubJobsFileAttr `json:"FileAttr,omitempty" xml:"FileAttr,omitempty" type:"Struct"`
	// example:
	//
	// oss://default/bucket-name/path-xxx/xxx-1.mp4
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xxxxx
	SubJobId *string `json:"SubJobId,omitempty" xml:"SubJobId,omitempty"`
}

func (s GetSmartClipTaskResponseBodyDataSubJobs) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskResponseBodyDataSubJobs) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) GetFileAttr() *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	return s.FileAttr
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) GetFileKey() *string {
	return s.FileKey
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) GetStatus() *string {
	return s.Status
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) GetSubJobId() *string {
	return s.SubJobId
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) SetErrorMessage(v string) *GetSmartClipTaskResponseBodyDataSubJobs {
	s.ErrorMessage = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) SetFileAttr(v *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) *GetSmartClipTaskResponseBodyDataSubJobs {
	s.FileAttr = v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) SetFileKey(v string) *GetSmartClipTaskResponseBodyDataSubJobs {
	s.FileKey = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) SetStatus(v string) *GetSmartClipTaskResponseBodyDataSubJobs {
	s.Status = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) SetSubJobId(v string) *GetSmartClipTaskResponseBodyDataSubJobs {
	s.SubJobId = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobs) Validate() error {
	if s.FileAttr != nil {
		if err := s.FileAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmartClipTaskResponseBodyDataSubJobsFileAttr struct {
	// example:
	//
	// 120
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 290804
	FileLength *string `json:"FileLength,omitempty" xml:"FileLength,omitempty"`
	// example:
	//
	// 2024-12-12.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://www.example.com/tmp.mp4
	TmpUrl *string `json:"TmpUrl,omitempty" xml:"TmpUrl,omitempty"`
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetSmartClipTaskResponseBodyDataSubJobsFileAttr) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetDuration() *float64 {
	return s.Duration
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetFileLength() *string {
	return s.FileLength
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetFileName() *string {
	return s.FileName
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetHeight() *int32 {
	return s.Height
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetTmpUrl() *string {
	return s.TmpUrl
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) GetWidth() *int32 {
	return s.Width
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetDuration(v float64) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.Duration = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetFileLength(v string) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.FileLength = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetFileName(v string) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.FileName = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetHeight(v int32) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.Height = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetTmpUrl(v string) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.TmpUrl = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) SetWidth(v int32) *GetSmartClipTaskResponseBodyDataSubJobsFileAttr {
	s.Width = &v
	return s
}

func (s *GetSmartClipTaskResponseBodyDataSubJobsFileAttr) Validate() error {
	return dara.Validate(s)
}
