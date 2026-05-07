// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCdnDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *GenerateCdnDiagnoseRequest
	GetUrl() *string
}

type GenerateCdnDiagnoseRequest struct {
	// example:
	//
	// http://www.example.com/xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateCdnDiagnoseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCdnDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCdnDiagnoseRequest) GetUrl() *string {
	return s.Url
}

func (s *GenerateCdnDiagnoseRequest) SetUrl(v string) *GenerateCdnDiagnoseRequest {
	s.Url = &v
	return s
}

func (s *GenerateCdnDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
