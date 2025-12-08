// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iPedestrianDetectAttributeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *PedestrianDetectAttributeAdvanceRequest
	GetImageURLObject() io.Reader
}

type PedestrianDetectAttributeAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/PedestrianDetectAttribute/PedestrianDetectAttribute1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s PedestrianDetectAttributeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *PedestrianDetectAttributeAdvanceRequest) SetImageURLObject(v io.Reader) *PedestrianDetectAttributeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *PedestrianDetectAttributeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
