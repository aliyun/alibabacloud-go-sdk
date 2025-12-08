// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTasks(v []*DeepfakeFaceRequestTasks) *DeepfakeFaceRequest
	GetTasks() []*DeepfakeFaceRequestTasks
}

type DeepfakeFaceRequest struct {
	// This parameter is required.
	Tasks []*DeepfakeFaceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DeepfakeFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceRequest) GetTasks() []*DeepfakeFaceRequestTasks {
	return s.Tasks
}

func (s *DeepfakeFaceRequest) SetTasks(v []*DeepfakeFaceRequestTasks) *DeepfakeFaceRequest {
	s.Tasks = v
	return s
}

func (s *DeepfakeFaceRequest) Validate() error {
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

type DeepfakeFaceRequestTasks struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2****
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DeepfakeFace/DeepfakeFace1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DeepfakeFaceRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceRequestTasks) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceRequestTasks) GetImageData() *string {
	return s.ImageData
}

func (s *DeepfakeFaceRequestTasks) GetImageURL() *string {
	return s.ImageURL
}

func (s *DeepfakeFaceRequestTasks) SetImageData(v string) *DeepfakeFaceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DeepfakeFaceRequestTasks) SetImageURL(v string) *DeepfakeFaceRequestTasks {
	s.ImageURL = &v
	return s
}

func (s *DeepfakeFaceRequestTasks) Validate() error {
	return dara.Validate(s)
}
