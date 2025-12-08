// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectInfraredLivingFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DetectInfraredLivingFaceAdvanceRequestTasks) *DetectInfraredLivingFaceAdvanceRequest
	GetTasks() []*DetectInfraredLivingFaceAdvanceRequestTasks
}

type DetectInfraredLivingFaceAdvanceRequest struct {
	// This parameter is required.
	Tasks []*DetectInfraredLivingFaceAdvanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectInfraredLivingFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceAdvanceRequest) GetTasks() []*DetectInfraredLivingFaceAdvanceRequestTasks {
	return s.Tasks
}

func (s *DetectInfraredLivingFaceAdvanceRequest) SetTasks(v []*DetectInfraredLivingFaceAdvanceRequestTasks) *DetectInfraredLivingFaceAdvanceRequest {
	s.Tasks = v
	return s
}

func (s *DetectInfraredLivingFaceAdvanceRequest) Validate() error {
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

type DetectInfraredLivingFaceAdvanceRequestTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectInfraredLivingFace/DetectInfraredLivingFace.jpeg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectInfraredLivingFaceAdvanceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceAdvanceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceAdvanceRequestTasks) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectInfraredLivingFaceAdvanceRequestTasks) SetImageURLObject(v io.Reader) *DetectInfraredLivingFaceAdvanceRequestTasks {
	s.ImageURLObject = v
	return s
}

func (s *DetectInfraredLivingFaceAdvanceRequestTasks) Validate() error {
	return dara.Validate(s)
}
