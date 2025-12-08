// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectLivingFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DetectLivingFaceRequestTasks) *DetectLivingFaceRequest
	GetTasks() []*DetectLivingFaceRequestTasks
}

type DetectLivingFaceRequest struct {
	// This parameter is required.
	Tasks []*DetectLivingFaceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequest) GetTasks() []*DetectLivingFaceRequestTasks {
	return s.Tasks
}

func (s *DetectLivingFaceRequest) SetTasks(v []*DetectLivingFaceRequestTasks) *DetectLivingFaceRequest {
	s.Tasks = v
	return s
}

func (s *DetectLivingFaceRequest) Validate() error {
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

type DetectLivingFaceRequestTasks struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectLivingFace/DetectLivingFace4.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectLivingFaceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequestTasks) GetImageData() *string {
	return s.ImageData
}

func (s *DetectLivingFaceRequestTasks) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectLivingFaceRequestTasks) SetImageData(v string) *DetectLivingFaceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DetectLivingFaceRequestTasks) SetImageURL(v string) *DetectLivingFaceRequestTasks {
	s.ImageURL = &v
	return s
}

func (s *DetectLivingFaceRequestTasks) Validate() error {
	return dara.Validate(s)
}
