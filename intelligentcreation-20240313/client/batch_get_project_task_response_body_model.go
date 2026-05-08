// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetProjectTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetProjectTaskResponseBody
	GetRequestId() *string
	SetResultList(v []*BatchGetProjectTaskResponseBodyResultList) *BatchGetProjectTaskResponseBody
	GetResultList() []*BatchGetProjectTaskResponseBodyResultList
}

type BatchGetProjectTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 11
	RequestId  *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultList []*BatchGetProjectTaskResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
}

func (s BatchGetProjectTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetProjectTaskResponseBody) GetResultList() []*BatchGetProjectTaskResponseBodyResultList {
	return s.ResultList
}

func (s *BatchGetProjectTaskResponseBody) SetRequestId(v string) *BatchGetProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetProjectTaskResponseBody) SetResultList(v []*BatchGetProjectTaskResponseBodyResultList) *BatchGetProjectTaskResponseBody {
	s.ResultList = v
	return s
}

func (s *BatchGetProjectTaskResponseBody) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetProjectTaskResponseBodyResultList struct {
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 11
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// http
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	// example:
	//
	// 1000
	VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
	// example:
	//
	// http
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s BatchGetProjectTaskResponseBodyResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetProjectTaskResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetStatus() *string {
	return s.Status
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetVideoDownloadUrl() *string {
	return s.VideoDownloadUrl
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetVideoDuration() *int32 {
	return s.VideoDuration
}

func (s *BatchGetProjectTaskResponseBodyResultList) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetErrorMsg(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetStatus(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetTaskId(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoDownloadUrl(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoDownloadUrl = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoDuration(v int32) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoDuration = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoUrl(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoUrl = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) Validate() error {
	return dara.Validate(s)
}
