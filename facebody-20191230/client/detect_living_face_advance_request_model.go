// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectLivingFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DetectLivingFaceAdvanceRequestTasks) *DetectLivingFaceAdvanceRequest
	GetTasks() []*DetectLivingFaceAdvanceRequestTasks
}

type DetectLivingFaceAdvanceRequest struct {
	// This parameter is required.
	Tasks []*DetectLivingFaceAdvanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceAdvanceRequest) GetTasks() []*DetectLivingFaceAdvanceRequestTasks {
	return s.Tasks
}

func (s *DetectLivingFaceAdvanceRequest) SetTasks(v []*DetectLivingFaceAdvanceRequestTasks) *DetectLivingFaceAdvanceRequest {
	s.Tasks = v
	return s
}

func (s *DetectLivingFaceAdvanceRequest) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetectLivingFaceAdvanceRequestTasks struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectLivingFace/DetectLivingFace4.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectLivingFaceAdvanceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceAdvanceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceAdvanceRequestTasks) GetImageData() *string {
	return s.ImageData
}

func (s *DetectLivingFaceAdvanceRequestTasks) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectLivingFaceAdvanceRequestTasks) SetImageData(v string) *DetectLivingFaceAdvanceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DetectLivingFaceAdvanceRequestTasks) SetImageURLObject(v io.Reader) *DetectLivingFaceAdvanceRequestTasks {
	s.ImageURLObject = v
	return s
}

func (s *DetectLivingFaceAdvanceRequestTasks) Validate() error {
	return dara.Validate(s)
}
