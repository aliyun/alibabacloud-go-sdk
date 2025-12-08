// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectCelebrityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectCelebrityRequest
	GetImageURL() *string
}

type DetectCelebrityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-oss.oss-cn-shanghai.aliyuncs.com/doc/facebody/xxx.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCelebrityRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectCelebrityRequest) SetImageURL(v string) *DetectCelebrityRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectCelebrityRequest) Validate() error {
	return dara.Validate(s)
}
