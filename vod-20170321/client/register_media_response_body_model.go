// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedFileURLs(v []*string) *RegisterMediaResponseBody
	GetFailedFileURLs() []*string
	SetRegisteredMediaList(v []*RegisterMediaResponseBodyRegisteredMediaList) *RegisterMediaResponseBody
	GetRegisteredMediaList() []*RegisterMediaResponseBodyRegisteredMediaList
	SetRequestId(v string) *RegisterMediaResponseBody
	GetRequestId() *string
}

type RegisterMediaResponseBody struct {
	// The URLs of the media files that failed to be registered.
	FailedFileURLs []*string `json:"FailedFileURLs,omitempty" xml:"FailedFileURLs,omitempty" type:"Repeated"`
	// The media files that are registered, including newly registered and repeatedly registered media files.
	RegisteredMediaList []*RegisterMediaResponseBodyRegisteredMediaList `json:"RegisteredMediaList,omitempty" xml:"RegisteredMediaList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 14F43C5C-8033-448B-AD04F64E5098****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponseBody) GetFailedFileURLs() []*string {
	return s.FailedFileURLs
}

func (s *RegisterMediaResponseBody) GetRegisteredMediaList() []*RegisterMediaResponseBodyRegisteredMediaList {
	return s.RegisteredMediaList
}

func (s *RegisterMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterMediaResponseBody) SetFailedFileURLs(v []*string) *RegisterMediaResponseBody {
	s.FailedFileURLs = v
	return s
}

func (s *RegisterMediaResponseBody) SetRegisteredMediaList(v []*RegisterMediaResponseBodyRegisteredMediaList) *RegisterMediaResponseBody {
	s.RegisteredMediaList = v
	return s
}

func (s *RegisterMediaResponseBody) SetRequestId(v string) *RegisterMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterMediaResponseBody) Validate() error {
	return dara.Validate(s)
}

type RegisterMediaResponseBodyRegisteredMediaList struct {
	// The URL of the media file.
	//
	// example:
	//
	// http://****.oss-cn-shanghai.aliyuncs.com/vod_sample_01.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the media file that is registered with ApsaraVideo VOD. If the registered media file is an audio or video file, the value of this parameter is the same as that of the VideoId parameter.
	//
	// example:
	//
	// d97af32828084d1896683b1aa38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// Indicates whether the media file is newly registered or repeatedly registered. Valid values:
	//
	// 	- **true**: The media file is newly registered.
	//
	// 	- **false**: The media file is repeatedly registered.
	//
	// example:
	//
	// false
	NewRegister *bool `json:"NewRegister,omitempty" xml:"NewRegister,omitempty"`
}

func (s RegisterMediaResponseBodyRegisteredMediaList) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaResponseBodyRegisteredMediaList) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) GetFileURL() *string {
	return s.FileURL
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) GetNewRegister() *bool {
	return s.NewRegister
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetFileURL(v string) *RegisterMediaResponseBodyRegisteredMediaList {
	s.FileURL = &v
	return s
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetMediaId(v string) *RegisterMediaResponseBodyRegisteredMediaList {
	s.MediaId = &v
	return s
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetNewRegister(v bool) *RegisterMediaResponseBodyRegisteredMediaList {
	s.NewRegister = &v
	return s
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) Validate() error {
	return dara.Validate(s)
}
