// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v int32) *RecognizeActionRequest
	GetType() *int32
	SetURLList(v []*RecognizeActionRequestURLList) *RecognizeActionRequest
	GetURLList() []*RecognizeActionRequestURLList
	SetVideoData(v string) *RecognizeActionRequest
	GetVideoData() *string
	SetVideoUrl(v string) *RecognizeActionRequest
	GetVideoUrl() *string
}

type RecognizeActionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type    *int32                           `json:"Type,omitempty" xml:"Type,omitempty"`
	URLList []*RecognizeActionRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
	// example:
	//
	// iVBORw0KGgoAAAANSUhEUgAAAoAAAAHJCAIAAACaEB9NAAEAAElEQVR4nNT9Wb****
	VideoData *string `json:"VideoData,omitempty" xml:"VideoData,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeAction/RecognizeAction-video1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeActionRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequest) GetType() *int32 {
	return s.Type
}

func (s *RecognizeActionRequest) GetURLList() []*RecognizeActionRequestURLList {
	return s.URLList
}

func (s *RecognizeActionRequest) GetVideoData() *string {
	return s.VideoData
}

func (s *RecognizeActionRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RecognizeActionRequest) SetType(v int32) *RecognizeActionRequest {
	s.Type = &v
	return s
}

func (s *RecognizeActionRequest) SetURLList(v []*RecognizeActionRequestURLList) *RecognizeActionRequest {
	s.URLList = v
	return s
}

func (s *RecognizeActionRequest) SetVideoData(v string) *RecognizeActionRequest {
	s.VideoData = &v
	return s
}

func (s *RecognizeActionRequest) SetVideoUrl(v string) *RecognizeActionRequest {
	s.VideoUrl = &v
	return s
}

func (s *RecognizeActionRequest) Validate() error {
	if s.URLList != nil {
		for _, item := range s.URLList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeActionRequestURLList struct {
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeAction/1RecognizeAction1.png
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgQ****
	ImageData *string `json:"imageData,omitempty" xml:"imageData,omitempty"`
}

func (s RecognizeActionRequestURLList) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionRequestURLList) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequestURLList) GetURL() *string {
	return s.URL
}

func (s *RecognizeActionRequestURLList) GetImageData() *string {
	return s.ImageData
}

func (s *RecognizeActionRequestURLList) SetURL(v string) *RecognizeActionRequestURLList {
	s.URL = &v
	return s
}

func (s *RecognizeActionRequestURLList) SetImageData(v string) *RecognizeActionRequestURLList {
	s.ImageData = &v
	return s
}

func (s *RecognizeActionRequestURLList) Validate() error {
	return dara.Validate(s)
}
