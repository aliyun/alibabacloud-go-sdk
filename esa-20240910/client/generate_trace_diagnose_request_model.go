// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTraceDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *GenerateTraceDiagnoseRequest
	GetUrl() *string
}

type GenerateTraceDiagnoseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://www.example.com/xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateTraceDiagnoseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTraceDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GenerateTraceDiagnoseRequest) GetUrl() *string {
	return s.Url
}

func (s *GenerateTraceDiagnoseRequest) SetUrl(v string) *GenerateTraceDiagnoseRequest {
	s.Url = &v
	return s
}

func (s *GenerateTraceDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
