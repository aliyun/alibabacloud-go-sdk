// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeVideoCharacterAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoURLObject(v io.Reader) *RecognizeVideoCharacterAdvanceRequest
	GetVideoURLObject() io.Reader
}

type RecognizeVideoCharacterAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/RecognizeVideoCharacter/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s RecognizeVideoCharacterAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *RecognizeVideoCharacterAdvanceRequest) SetVideoURLObject(v io.Reader) *RecognizeVideoCharacterAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *RecognizeVideoCharacterAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
