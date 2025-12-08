// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeActionAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v int32) *RecognizeActionAdvanceRequest
	GetType() *int32
	SetURLList(v []*RecognizeActionAdvanceRequestURLList) *RecognizeActionAdvanceRequest
	GetURLList() []*RecognizeActionAdvanceRequestURLList
	SetVideoData(v string) *RecognizeActionAdvanceRequest
	GetVideoData() *string
	SetVideoUrlObject(v io.Reader) *RecognizeActionAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type RecognizeActionAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type    *int32                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	URLList []*RecognizeActionAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
	// example:
	//
	// iVBORw0KGgoAAAANSUhEUgAAAoAAAAHJCAIAAACaEB9NAAEAAElEQVR4nNT9Wb****
	VideoData *string `json:"VideoData,omitempty" xml:"VideoData,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeAction/RecognizeAction-video1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeActionAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeActionAdvanceRequest) GetType() *int32 {
	return s.Type
}

func (s *RecognizeActionAdvanceRequest) GetURLList() []*RecognizeActionAdvanceRequestURLList {
	return s.URLList
}

func (s *RecognizeActionAdvanceRequest) GetVideoData() *string {
	return s.VideoData
}

func (s *RecognizeActionAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *RecognizeActionAdvanceRequest) SetType(v int32) *RecognizeActionAdvanceRequest {
	s.Type = &v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetURLList(v []*RecognizeActionAdvanceRequestURLList) *RecognizeActionAdvanceRequest {
	s.URLList = v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetVideoData(v string) *RecognizeActionAdvanceRequest {
	s.VideoData = &v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeActionAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *RecognizeActionAdvanceRequest) Validate() error {
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

type RecognizeActionAdvanceRequestURLList struct {
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeAction/1RecognizeAction1.png
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgQ****
	ImageData *string `json:"imageData,omitempty" xml:"imageData,omitempty"`
}

func (s RecognizeActionAdvanceRequestURLList) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *RecognizeActionAdvanceRequestURLList) GetURLObject() io.Reader {
	return s.URLObject
}

func (s *RecognizeActionAdvanceRequestURLList) GetImageData() *string {
	return s.ImageData
}

func (s *RecognizeActionAdvanceRequestURLList) SetURLObject(v io.Reader) *RecognizeActionAdvanceRequestURLList {
	s.URLObject = v
	return s
}

func (s *RecognizeActionAdvanceRequestURLList) SetImageData(v string) *RecognizeActionAdvanceRequestURLList {
	s.ImageData = &v
	return s
}

func (s *RecognizeActionAdvanceRequestURLList) Validate() error {
	return dara.Validate(s)
}
