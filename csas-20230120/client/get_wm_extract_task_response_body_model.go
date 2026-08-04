// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmExtractTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetWmExtractTaskResponseBodyData) *GetWmExtractTaskResponseBody
	GetData() *GetWmExtractTaskResponseBodyData
	SetRequestId(v string) *GetWmExtractTaskResponseBody
	GetRequestId() *string
}

type GetWmExtractTaskResponseBody struct {
	// The task result.
	Data *GetWmExtractTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWmExtractTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWmExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponseBody) GetData() *GetWmExtractTaskResponseBodyData {
	return s.Data
}

func (s *GetWmExtractTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWmExtractTaskResponseBody) SetData(v *GetWmExtractTaskResponseBodyData) *GetWmExtractTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetWmExtractTaskResponseBody) SetRequestId(v string) *GetWmExtractTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWmExtractTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWmExtractTaskResponseBodyData struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2024-01-01 11:22:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the file that was provided when the extraction task was created.
	//
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// The status of the task. Valid values:
	//
	// - **Running**: The task is running.
	//
	// - **Success**: The task is successful.
	//
	// - **Failed**: The task failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The extracted watermark information as a Base64-encoded string.
	//
	// example:
	//
	// aGVsbG8gc2Fz****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// The size of the watermark information, which was provided when the extraction task was created.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The extracted watermark information in decimal format.
	//
	// example:
	//
	// 123**
	WmInfoUint *int64 `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// The watermark type that was provided when the extraction task was created.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s GetWmExtractTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWmExtractTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWmExtractTaskResponseBodyData) GetFilename() *string {
	return s.Filename
}

func (s *GetWmExtractTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetWmExtractTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWmExtractTaskResponseBodyData) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *GetWmExtractTaskResponseBodyData) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *GetWmExtractTaskResponseBodyData) GetWmInfoUint() *int64 {
	return s.WmInfoUint
}

func (s *GetWmExtractTaskResponseBodyData) GetWmType() *string {
	return s.WmType
}

func (s *GetWmExtractTaskResponseBodyData) SetCreateTime(v string) *GetWmExtractTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetFilename(v string) *GetWmExtractTaskResponseBodyData {
	s.Filename = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetStatus(v string) *GetWmExtractTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetTaskId(v string) *GetWmExtractTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoBytesB64(v string) *GetWmExtractTaskResponseBodyData {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoSize(v int64) *GetWmExtractTaskResponseBodyData {
	s.WmInfoSize = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoUint(v int64) *GetWmExtractTaskResponseBodyData {
	s.WmInfoUint = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmType(v string) *GetWmExtractTaskResponseBodyData {
	s.WmType = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
