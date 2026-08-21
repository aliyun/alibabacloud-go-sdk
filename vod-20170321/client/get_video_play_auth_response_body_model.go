// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlayAuth(v string) *GetVideoPlayAuthResponseBody
	GetPlayAuth() *string
	SetRequestId(v string) *GetVideoPlayAuthResponseBody
	GetRequestId() *string
	SetVideoMeta(v *GetVideoPlayAuthResponseBodyVideoMeta) *GetVideoPlayAuthResponseBody
	GetVideoMeta() *GetVideoPlayAuthResponseBodyVideoMeta
}

type GetVideoPlayAuthResponseBody struct {
	// The playback credential for the audio or video file.
	//
	// example:
	//
	// sstyYuew6789000000xtt7TYUh****
	PlayAuth *string `json:"PlayAuth,omitempty" xml:"PlayAuth,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the audio or video file.
	VideoMeta *GetVideoPlayAuthResponseBodyVideoMeta `json:"VideoMeta,omitempty" xml:"VideoMeta,omitempty" type:"Struct"`
}

func (s GetVideoPlayAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponseBody) GetPlayAuth() *string {
	return s.PlayAuth
}

func (s *GetVideoPlayAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoPlayAuthResponseBody) GetVideoMeta() *GetVideoPlayAuthResponseBodyVideoMeta {
	return s.VideoMeta
}

func (s *GetVideoPlayAuthResponseBody) SetPlayAuth(v string) *GetVideoPlayAuthResponseBody {
	s.PlayAuth = &v
	return s
}

func (s *GetVideoPlayAuthResponseBody) SetRequestId(v string) *GetVideoPlayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoPlayAuthResponseBody) SetVideoMeta(v *GetVideoPlayAuthResponseBodyVideoMeta) *GetVideoPlayAuthResponseBody {
	s.VideoMeta = v
	return s
}

func (s *GetVideoPlayAuthResponseBody) Validate() error {
	if s.VideoMeta != nil {
		if err := s.VideoMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoPlayAuthResponseBodyVideoMeta struct {
	// The thumbnail URL of the audio or video file.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The duration of the audio or video file. Unit: seconds.
	//
	// example:
	//
	// 120.0
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The status of the audio or video file. For valid values and descriptions, see [Status: audio and video status](~~52839#title-vqg-8cz-7p8~~).
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the audio or video file.
	//
	// example:
	//
	// Alibaba Cloud VOD Video Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The audio or video ID.
	//
	// example:
	//
	// dfde02284a5c46622a097adaf44a****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayAuthResponseBodyVideoMeta) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayAuthResponseBodyVideoMeta) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) GetStatus() *string {
	return s.Status
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) GetTitle() *string {
	return s.Title
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetCoverURL(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.CoverURL = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetDuration(v float32) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Duration = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetStatus(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Status = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetTitle(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Title = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetVideoId(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.VideoId = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) Validate() error {
	return dara.Validate(s)
}
