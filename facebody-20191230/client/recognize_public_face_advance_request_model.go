// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizePublicFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTask(v []*RecognizePublicFaceAdvanceRequestTask) *RecognizePublicFaceAdvanceRequest
	GetTask() []*RecognizePublicFaceAdvanceRequestTask
}

type RecognizePublicFaceAdvanceRequest struct {
	// 1
	Task []*RecognizePublicFaceAdvanceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceAdvanceRequest) GetTask() []*RecognizePublicFaceAdvanceRequestTask {
	return s.Task
}

func (s *RecognizePublicFaceAdvanceRequest) SetTask(v []*RecognizePublicFaceAdvanceRequestTask) *RecognizePublicFaceAdvanceRequest {
	s.Task = v
	return s
}

func (s *RecognizePublicFaceAdvanceRequest) Validate() error {
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

type RecognizePublicFaceAdvanceRequestTask struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// https://viapi-oss.oss-cn-shanghai.aliyuncs.com/doc/facebody/xxx.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePublicFaceAdvanceRequestTask) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceAdvanceRequestTask) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceAdvanceRequestTask) GetImageData() *string {
	return s.ImageData
}

func (s *RecognizePublicFaceAdvanceRequestTask) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizePublicFaceAdvanceRequestTask) SetImageData(v string) *RecognizePublicFaceAdvanceRequestTask {
	s.ImageData = &v
	return s
}

func (s *RecognizePublicFaceAdvanceRequestTask) SetImageURLObject(v io.Reader) *RecognizePublicFaceAdvanceRequestTask {
	s.ImageURLObject = v
	return s
}

func (s *RecognizePublicFaceAdvanceRequestTask) Validate() error {
	return dara.Validate(s)
}
