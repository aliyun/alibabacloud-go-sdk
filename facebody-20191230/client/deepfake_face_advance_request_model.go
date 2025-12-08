// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDeepfakeFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DeepfakeFaceAdvanceRequestTasks) *DeepfakeFaceAdvanceRequest
	GetTasks() []*DeepfakeFaceAdvanceRequestTasks
}

type DeepfakeFaceAdvanceRequest struct {
	// This parameter is required.
	Tasks []*DeepfakeFaceAdvanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DeepfakeFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceAdvanceRequest) GetTasks() []*DeepfakeFaceAdvanceRequestTasks {
	return s.Tasks
}

func (s *DeepfakeFaceAdvanceRequest) SetTasks(v []*DeepfakeFaceAdvanceRequestTasks) *DeepfakeFaceAdvanceRequest {
	s.Tasks = v
	return s
}

func (s *DeepfakeFaceAdvanceRequest) Validate() error {
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

type DeepfakeFaceAdvanceRequestTasks struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2****
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DeepfakeFace/DeepfakeFace1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DeepfakeFaceAdvanceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceAdvanceRequestTasks) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceAdvanceRequestTasks) GetImageData() *string {
	return s.ImageData
}

func (s *DeepfakeFaceAdvanceRequestTasks) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DeepfakeFaceAdvanceRequestTasks) SetImageData(v string) *DeepfakeFaceAdvanceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DeepfakeFaceAdvanceRequestTasks) SetImageURLObject(v io.Reader) *DeepfakeFaceAdvanceRequestTasks {
	s.ImageURLObject = v
	return s
}

func (s *DeepfakeFaceAdvanceRequestTasks) Validate() error {
	return dara.Validate(s)
}
