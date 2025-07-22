// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSaveImage(v bool) *StopInstanceRequest
	GetSaveImage() *bool
}

type StopInstanceRequest struct {
	// example:
	//
	// false
	SaveImage *bool `json:"SaveImage,omitempty" xml:"SaveImage,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetSaveImage() *bool {
	return s.SaveImage
}

func (s *StopInstanceRequest) SetSaveImage(v bool) *StopInstanceRequest {
	s.SaveImage = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
