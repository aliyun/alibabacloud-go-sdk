// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPedestrianDetectAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *PedestrianDetectAttributeRequest
	GetImageURL() *string
}

type PedestrianDetectAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/PedestrianDetectAttribute/PedestrianDetectAttribute1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s PedestrianDetectAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *PedestrianDetectAttributeRequest) SetImageURL(v string) *PedestrianDetectAttributeRequest {
	s.ImageURL = &v
	return s
}

func (s *PedestrianDetectAttributeRequest) Validate() error {
	return dara.Validate(s)
}
