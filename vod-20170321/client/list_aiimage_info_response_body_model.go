// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIImageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIImageInfoList(v []*ListAIImageInfoResponseBodyAIImageInfoList) *ListAIImageInfoResponseBody
	GetAIImageInfoList() []*ListAIImageInfoResponseBodyAIImageInfoList
	SetRequestId(v string) *ListAIImageInfoResponseBody
	GetRequestId() *string
}

type ListAIImageInfoResponseBody struct {
	// The image files that are uploaded for AI processing.
	AIImageInfoList []*ListAIImageInfoResponseBodyAIImageInfoList `json:"AIImageInfoList,omitempty" xml:"AIImageInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D73420CD-D221-9122-5B8FA995A511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIImageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponseBody) GetAIImageInfoList() []*ListAIImageInfoResponseBodyAIImageInfoList {
	return s.AIImageInfoList
}

func (s *ListAIImageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIImageInfoResponseBody) SetAIImageInfoList(v []*ListAIImageInfoResponseBodyAIImageInfoList) *ListAIImageInfoResponseBody {
	s.AIImageInfoList = v
	return s
}

func (s *ListAIImageInfoResponseBody) SetRequestId(v string) *ListAIImageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIImageInfoResponseBody) Validate() error {
	if s.AIImageInfoList != nil {
		for _, item := range s.AIImageInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIImageInfoResponseBodyAIImageInfoList struct {
	// The ID of the image information.
	//
	// example:
	//
	// b89a6aabf1b6197ebd6fe6cf29****
	AIImageInfoId *string `json:"AIImageInfoId,omitempty" xml:"AIImageInfoId,omitempty"`
	// The time when the file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-15T03:30:03Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the image file.
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The format of the image. Valid values: **gif*	- and **png**.
	//
	// example:
	//
	// gif
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The ID of the image AI processing job.
	//
	// example:
	//
	// cf08a2c6e11ee1711b738b9067****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The score of the image.
	//
	// example:
	//
	// 5.035636554444242
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The data version ID.
	//
	// example:
	//
	// b547f3f0e199c3b457369f3cf****
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 357a8748c5789d2726e6436aa****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListAIImageInfoResponseBodyAIImageInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAIImageInfoResponseBodyAIImageInfoList) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetAIImageInfoId() *string {
	return s.AIImageInfoId
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetFileURL() *string {
	return s.FileURL
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetFormat() *string {
	return s.Format
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetJobId() *string {
	return s.JobId
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetScore() *string {
	return s.Score
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) GetVideoId() *string {
	return s.VideoId
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetAIImageInfoId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.AIImageInfoId = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetCreationTime(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetFileURL(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.FileURL = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetFormat(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Format = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetJobId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.JobId = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetScore(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Score = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetVersion(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Version = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetVideoId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.VideoId = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) Validate() error {
	return dara.Validate(s)
}
