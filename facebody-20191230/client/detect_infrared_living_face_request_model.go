// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectInfraredLivingFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DetectInfraredLivingFaceRequestTasks) *DetectInfraredLivingFaceRequest
	GetTasks() []*DetectInfraredLivingFaceRequestTasks
}

type DetectInfraredLivingFaceRequest struct {
	// This parameter is required.
	Tasks []*DetectInfraredLivingFaceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectInfraredLivingFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceRequest) GetTasks() []*DetectInfraredLivingFaceRequestTasks {
	return s.Tasks
}

func (s *DetectInfraredLivingFaceRequest) SetTasks(v []*DetectInfraredLivingFaceRequestTasks) *DetectInfraredLivingFaceRequest {
	s.Tasks = v
	return s
}

func (s *DetectInfraredLivingFaceRequest) Validate() error {
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

type DetectInfraredLivingFaceRequestTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectInfraredLivingFace/DetectInfraredLivingFace.jpeg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectInfraredLivingFaceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceRequestTasks) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectInfraredLivingFaceRequestTasks) SetImageURL(v string) *DetectInfraredLivingFaceRequestTasks {
	s.ImageURL = &v
	return s
}

func (s *DetectInfraredLivingFaceRequestTasks) Validate() error {
	return dara.Validate(s)
}
