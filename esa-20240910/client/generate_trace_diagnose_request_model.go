// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTraceDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GenerateTraceDiagnoseRequest
	GetSource() *string
	SetUrl(v string) *GenerateTraceDiagnoseRequest
	GetUrl() *string
}

type GenerateTraceDiagnoseRequest struct {
	// The source of the request.
	//
	// example:
	//
	// ai
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The URL to diagnose.
	//
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

func (s *GenerateTraceDiagnoseRequest) GetSource() *string {
	return s.Source
}

func (s *GenerateTraceDiagnoseRequest) GetUrl() *string {
	return s.Url
}

func (s *GenerateTraceDiagnoseRequest) SetSource(v string) *GenerateTraceDiagnoseRequest {
	s.Source = &v
	return s
}

func (s *GenerateTraceDiagnoseRequest) SetUrl(v string) *GenerateTraceDiagnoseRequest {
	s.Url = &v
	return s
}

func (s *GenerateTraceDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
