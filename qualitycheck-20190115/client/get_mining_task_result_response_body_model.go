// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiningTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMiningTaskResultResponseBody
	GetCode() *string
	SetData(v *GetMiningTaskResultResponseBodyData) *GetMiningTaskResultResponseBody
	GetData() *GetMiningTaskResultResponseBodyData
	SetMessage(v string) *GetMiningTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMiningTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetMiningTaskResultResponseBody
	GetSuccess() *string
}

type GetMiningTaskResultResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetMiningTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMiningTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMiningTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMiningTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMiningTaskResultResponseBody) GetData() *GetMiningTaskResultResponseBodyData {
	return s.Data
}

func (s *GetMiningTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMiningTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMiningTaskResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetMiningTaskResultResponseBody) SetCode(v string) *GetMiningTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetMiningTaskResultResponseBody) SetData(v *GetMiningTaskResultResponseBodyData) *GetMiningTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetMiningTaskResultResponseBody) SetMessage(v string) *GetMiningTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetMiningTaskResultResponseBody) SetRequestId(v string) *GetMiningTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMiningTaskResultResponseBody) SetSuccess(v string) *GetMiningTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetMiningTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMiningTaskResultResponseBodyData struct {
	// example:
	//
	// 123.22.com/manger/static/login-back.jpg
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 123.22.com/manger/static/login-back.md
	FilePathMd *string `json:"FilePathMd,omitempty" xml:"FilePathMd,omitempty"`
	// example:
	//
	// 20201231de3d34ec-40fa-4a55-8d27-76ea*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// finish
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetMiningTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMiningTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMiningTaskResultResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetMiningTaskResultResponseBodyData) GetFilePathMd() *string {
	return s.FilePathMd
}

func (s *GetMiningTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMiningTaskResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetMiningTaskResultResponseBodyData) SetFilePath(v string) *GetMiningTaskResultResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetMiningTaskResultResponseBodyData) SetFilePathMd(v string) *GetMiningTaskResultResponseBodyData {
	s.FilePathMd = &v
	return s
}

func (s *GetMiningTaskResultResponseBodyData) SetTaskId(v string) *GetMiningTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetMiningTaskResultResponseBodyData) SetTaskStatus(v string) *GetMiningTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetMiningTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
