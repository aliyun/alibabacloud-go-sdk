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
	// The list of file URLs that failed to be registered.
	FailedFileURLs []*string `json:"FailedFileURLs,omitempty" xml:"FailedFileURLs,omitempty" type:"Repeated"`
	// The list of media assets that are successfully registered, including both newly registered files and previously registered files.
	RegisteredMediaList []*RegisterMediaResponseBodyRegisteredMediaList `json:"RegisteredMediaList,omitempty" xml:"RegisteredMediaList,omitempty" type:"Repeated"`
	// The request ID.
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
	if s.RegisteredMediaList != nil {
		for _, item := range s.RegisteredMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RegisterMediaResponseBodyRegisteredMediaList struct {
	// The OSS file URL.
	//
	// example:
	//
	// http://****.oss-cn-shanghai.aliyuncs.com/vod_sample_01.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The VOD media ID. If the registered media file is an audio or video file, this value corresponds to the VideoId in ApsaraVideo VOD.
	//
	// example:
	//
	// d97af32828084d1896683b1aa38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// Indicates whether the media asset is newly registered or repeatedly registered.
	//
	// - **true**: newly registered.
	//
	// - **false**: repeatedly registered.
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
