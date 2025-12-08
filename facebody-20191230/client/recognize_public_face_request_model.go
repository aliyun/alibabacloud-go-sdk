// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePublicFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTask(v []*RecognizePublicFaceRequestTask) *RecognizePublicFaceRequest
	GetTask() []*RecognizePublicFaceRequestTask
}

type RecognizePublicFaceRequest struct {
	// 1
	Task []*RecognizePublicFaceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequest) GetTask() []*RecognizePublicFaceRequestTask {
	return s.Task
}

func (s *RecognizePublicFaceRequest) SetTask(v []*RecognizePublicFaceRequestTask) *RecognizePublicFaceRequest {
	s.Task = v
	return s
}

func (s *RecognizePublicFaceRequest) Validate() error {
	if s.Task != nil {
		for _, item := range s.Task {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizePublicFaceRequestTask struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// https://viapi-oss.oss-cn-shanghai.aliyuncs.com/doc/facebody/xxx.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePublicFaceRequestTask) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceRequestTask) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequestTask) GetImageData() *string {
	return s.ImageData
}

func (s *RecognizePublicFaceRequestTask) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizePublicFaceRequestTask) SetImageData(v string) *RecognizePublicFaceRequestTask {
	s.ImageData = &v
	return s
}

func (s *RecognizePublicFaceRequestTask) SetImageURL(v string) *RecognizePublicFaceRequestTask {
	s.ImageURL = &v
	return s
}

func (s *RecognizePublicFaceRequestTask) Validate() error {
	return dara.Validate(s)
}
