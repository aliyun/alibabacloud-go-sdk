// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCharacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoURL(v string) *RecognizeVideoCharacterRequest
	GetVideoURL() *string
}

type RecognizeVideoCharacterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/RecognizeVideoCharacter/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s RecognizeVideoCharacterRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *RecognizeVideoCharacterRequest) SetVideoURL(v string) *RecognizeVideoCharacterRequest {
	s.VideoURL = &v
	return s
}

func (s *RecognizeVideoCharacterRequest) Validate() error {
	return dara.Validate(s)
}
