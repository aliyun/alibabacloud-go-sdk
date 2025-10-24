// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageCopyrightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *SubmitImageCopyrightRequest
	GetMessage() *string
	SetOutput(v string) *SubmitImageCopyrightRequest
	GetOutput() *string
	SetParams(v string) *SubmitImageCopyrightRequest
	GetParams() *string
}

type SubmitImageCopyrightRequest struct {
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// {"Bucket":"abc-test","Location":"oss-cn-shanghai","Object":"out.jpeg"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// {"width":2999, "height":2999, "afa": 3, "type":1, "version":0}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s SubmitImageCopyrightRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageCopyrightRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightRequest) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageCopyrightRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitImageCopyrightRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitImageCopyrightRequest) SetMessage(v string) *SubmitImageCopyrightRequest {
	s.Message = &v
	return s
}

func (s *SubmitImageCopyrightRequest) SetOutput(v string) *SubmitImageCopyrightRequest {
	s.Output = &v
	return s
}

func (s *SubmitImageCopyrightRequest) SetParams(v string) *SubmitImageCopyrightRequest {
	s.Params = &v
	return s
}

func (s *SubmitImageCopyrightRequest) Validate() error {
	return dara.Validate(s)
}
