// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBodyPostureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *BodyPostureRequest
	GetImageURL() *string
}

type BodyPostureRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/BodyPosture/BodyPosture4.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BodyPostureRequest) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *BodyPostureRequest) SetImageURL(v string) *BodyPostureRequest {
	s.ImageURL = &v
	return s
}

func (s *BodyPostureRequest) Validate() error {
	return dara.Validate(s)
}
