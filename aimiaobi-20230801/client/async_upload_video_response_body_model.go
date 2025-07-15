// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncUploadVideoResponseBody
	GetCode() *string
	SetData(v *AsyncUploadVideoResponseBodyData) *AsyncUploadVideoResponseBody
	GetData() *AsyncUploadVideoResponseBodyData
	SetRequestId(v string) *AsyncUploadVideoResponseBody
	GetRequestId() *string
}

type AsyncUploadVideoResponseBody struct {
	// example:
	//
	// successful
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsyncUploadVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AsyncUploadVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncUploadVideoResponseBody) GetData() *AsyncUploadVideoResponseBodyData {
	return s.Data
}

func (s *AsyncUploadVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncUploadVideoResponseBody) SetCode(v string) *AsyncUploadVideoResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncUploadVideoResponseBody) SetData(v *AsyncUploadVideoResponseBodyData) *AsyncUploadVideoResponseBody {
	s.Data = v
	return s
}

func (s *AsyncUploadVideoResponseBody) SetRequestId(v string) *AsyncUploadVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncUploadVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type AsyncUploadVideoResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId     *string                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VideoInfos []*AsyncUploadVideoResponseBodyDataVideoInfos `json:"VideoInfos,omitempty" xml:"VideoInfos,omitempty" type:"Repeated"`
}

func (s AsyncUploadVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncUploadVideoResponseBodyData) GetVideoInfos() []*AsyncUploadVideoResponseBodyDataVideoInfos {
	return s.VideoInfos
}

func (s *AsyncUploadVideoResponseBodyData) SetTaskId(v string) *AsyncUploadVideoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncUploadVideoResponseBodyData) SetVideoInfos(v []*AsyncUploadVideoResponseBodyDataVideoInfos) *AsyncUploadVideoResponseBodyData {
	s.VideoInfos = v
	return s
}

func (s *AsyncUploadVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AsyncUploadVideoResponseBodyDataVideoInfos struct {
	VideoExtraInfo *string `json:"VideoExtraInfo,omitempty" xml:"VideoExtraInfo,omitempty"`
	// example:
	//
	// 60616fad41b171f0bb4b4531948c0102
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// example:
	//
	// 123.mp4
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	// example:
	//
	// http://123.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncUploadVideoResponseBodyDataVideoInfos) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoResponseBodyDataVideoInfos) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) GetVideoId() *string {
	return s.VideoId
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) SetVideoExtraInfo(v string) *AsyncUploadVideoResponseBodyDataVideoInfos {
	s.VideoExtraInfo = &v
	return s
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) SetVideoId(v string) *AsyncUploadVideoResponseBodyDataVideoInfos {
	s.VideoId = &v
	return s
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) SetVideoName(v string) *AsyncUploadVideoResponseBodyDataVideoInfos {
	s.VideoName = &v
	return s
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) SetVideoUrl(v string) *AsyncUploadVideoResponseBodyDataVideoInfos {
	s.VideoUrl = &v
	return s
}

func (s *AsyncUploadVideoResponseBodyDataVideoInfos) Validate() error {
	return dara.Validate(s)
}
