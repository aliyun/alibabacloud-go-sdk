// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectPedestrianAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectPedestrianAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectPedestrianAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectPedestrian/DetectPedestrian8.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectPedestrianAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectPedestrianAdvanceRequest) SetImageURLObject(v io.Reader) *DetectPedestrianAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectPedestrianAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
