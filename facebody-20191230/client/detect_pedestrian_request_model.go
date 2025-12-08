// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectPedestrianRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectPedestrianRequest
	GetImageURL() *string
}

type DetectPedestrianRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectPedestrian/DetectPedestrian8.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectPedestrianRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectPedestrianRequest) SetImageURL(v string) *DetectPedestrianRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectPedestrianRequest) Validate() error {
	return dara.Validate(s)
}
