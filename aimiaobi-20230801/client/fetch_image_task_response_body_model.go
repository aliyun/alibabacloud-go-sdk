// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchImageTaskResponseBody
	GetCode() *string
	SetData(v *FetchImageTaskResponseBodyData) *FetchImageTaskResponseBody
	GetData() *FetchImageTaskResponseBodyData
	SetHttpStatusCode(v int32) *FetchImageTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchImageTaskResponseBody
	GetSuccess() *bool
}

type FetchImageTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// DD656AF9-0839-521A-A3D2-F320009F9C87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchImageTaskResponseBody) GetData() *FetchImageTaskResponseBodyData {
	return s.Data
}

func (s *FetchImageTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchImageTaskResponseBody) SetCode(v string) *FetchImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetData(v *FetchImageTaskResponseBodyData) *FetchImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchImageTaskResponseBody) SetHttpStatusCode(v int32) *FetchImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetMessage(v string) *FetchImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetRequestId(v string) *FetchImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetSuccess(v bool) *FetchImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *FetchImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type FetchImageTaskResponseBodyData struct {
	TaskInfoList []*FetchImageTaskResponseBodyDataTaskInfoList `json:"TaskInfoList,omitempty" xml:"TaskInfoList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyData) GetTaskInfoList() []*FetchImageTaskResponseBodyDataTaskInfoList {
	return s.TaskInfoList
}

func (s *FetchImageTaskResponseBodyData) SetTaskInfoList(v []*FetchImageTaskResponseBodyDataTaskInfoList) *FetchImageTaskResponseBodyData {
	s.TaskInfoList = v
	return s
}

func (s *FetchImageTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type FetchImageTaskResponseBodyDataTaskInfoList struct {
	// example:
	//
	// 1
	Id        *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageList []*FetchImageTaskResponseBodyDataTaskInfoListImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// example:
	//
	// net-7eb32699000d4193a3c59fc64ae1e55f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) GetId() *int64 {
	return s.Id
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) GetImageList() []*FetchImageTaskResponseBodyDataTaskInfoListImageList {
	return s.ImageList
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetId(v int64) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.Id = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetImageList(v []*FetchImageTaskResponseBodyDataTaskInfoListImageList) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.ImageList = v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskId(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskId = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskStatus(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskStatus = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) Validate() error {
	return dara.Validate(s)
}

type FetchImageTaskResponseBodyDataTaskInfoListImageList struct {
	// example:
	//
	// NoData
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) GetCode() *string {
	return s.Code
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) GetMessage() *string {
	return s.Message
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) GetUrl() *string {
	return s.Url
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetCode(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetMessage(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetUrl(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Url = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) Validate() error {
	return dara.Validate(s)
}
