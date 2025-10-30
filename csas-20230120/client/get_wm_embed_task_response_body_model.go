// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmEmbedTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetWmEmbedTaskResponseBodyData) *GetWmEmbedTaskResponseBody
	GetData() *GetWmEmbedTaskResponseBodyData
	SetRequestId(v string) *GetWmEmbedTaskResponseBody
	GetRequestId() *string
}

type GetWmEmbedTaskResponseBody struct {
	Data *GetWmEmbedTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWmEmbedTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWmEmbedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponseBody) GetData() *GetWmEmbedTaskResponseBodyData {
	return s.Data
}

func (s *GetWmEmbedTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWmEmbedTaskResponseBody) SetData(v *GetWmEmbedTaskResponseBodyData) *GetWmEmbedTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetWmEmbedTaskResponseBody) SetRequestId(v string) *GetWmEmbedTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWmEmbedTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWmEmbedTaskResponseBodyData struct {
	// example:
	//
	// https://example.com/embed-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 171859****
	FileUrlExp *string `json:"FileUrlExp,omitempty" xml:"FileUrlExp,omitempty"`
	// example:
	//
	// embed-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8****
	OutFileHashMd5 *string `json:"OutFileHashMd5,omitempty" xml:"OutFileHashMd5,omitempty"`
	// example:
	//
	// 123**
	OutFileSize *int64 `json:"OutFileSize,omitempty" xml:"OutFileSize,omitempty"`
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjUxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetWmEmbedTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWmEmbedTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetWmEmbedTaskResponseBodyData) GetFileUrlExp() *string {
	return s.FileUrlExp
}

func (s *GetWmEmbedTaskResponseBodyData) GetFilename() *string {
	return s.Filename
}

func (s *GetWmEmbedTaskResponseBodyData) GetOutFileHashMd5() *string {
	return s.OutFileHashMd5
}

func (s *GetWmEmbedTaskResponseBodyData) GetOutFileSize() *int64 {
	return s.OutFileSize
}

func (s *GetWmEmbedTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWmEmbedTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetWmEmbedTaskResponseBodyData) SetFileUrl(v string) *GetWmEmbedTaskResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetFileUrlExp(v string) *GetWmEmbedTaskResponseBodyData {
	s.FileUrlExp = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetFilename(v string) *GetWmEmbedTaskResponseBodyData {
	s.Filename = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetOutFileHashMd5(v string) *GetWmEmbedTaskResponseBodyData {
	s.OutFileHashMd5 = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetOutFileSize(v int64) *GetWmEmbedTaskResponseBodyData {
	s.OutFileSize = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetTaskId(v string) *GetWmEmbedTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetTaskStatus(v string) *GetWmEmbedTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
