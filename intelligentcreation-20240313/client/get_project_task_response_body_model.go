// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *GetProjectTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetProjectTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetProjectTaskResponseBody
	GetStatus() *string
	SetVideoDownloadUrl(v string) *GetProjectTaskResponseBody
	GetVideoDownloadUrl() *string
	SetVideoDuration(v int32) *GetProjectTaskResponseBody
	GetVideoDuration() *int32
	SetVideoUrl(v string) *GetProjectTaskResponseBody
	GetVideoUrl() *string
}

type GetProjectTaskResponseBody struct {
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 5389BE87-571B-573C-90ED-F07C5E68760B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// www.ali.com
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	// example:
	//
	// 111
	VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
	// example:
	//
	// www.ali.com
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s GetProjectTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetProjectTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetProjectTaskResponseBody) GetVideoDownloadUrl() *string {
	return s.VideoDownloadUrl
}

func (s *GetProjectTaskResponseBody) GetVideoDuration() *int32 {
	return s.VideoDuration
}

func (s *GetProjectTaskResponseBody) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GetProjectTaskResponseBody) SetErrorMsg(v string) *GetProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetRequestId(v string) *GetProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetStatus(v string) *GetProjectTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoDownloadUrl(v string) *GetProjectTaskResponseBody {
	s.VideoDownloadUrl = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoDuration(v int32) *GetProjectTaskResponseBody {
	s.VideoDuration = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoUrl(v string) *GetProjectTaskResponseBody {
	s.VideoUrl = &v
	return s
}

func (s *GetProjectTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
