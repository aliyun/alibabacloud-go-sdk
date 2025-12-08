// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBodyPostureAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *BodyPostureAdvanceRequest
	GetImageURLObject() io.Reader
}

type BodyPostureAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/BodyPosture/BodyPosture4.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BodyPostureAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *BodyPostureAdvanceRequest) SetImageURLObject(v io.Reader) *BodyPostureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *BodyPostureAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
