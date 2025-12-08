// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectCelebrityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectCelebrityAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectCelebrityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-oss.oss-cn-shanghai.aliyuncs.com/doc/facebody/xxx.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCelebrityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectCelebrityAdvanceRequest) SetImageURLObject(v io.Reader) *DetectCelebrityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectCelebrityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
