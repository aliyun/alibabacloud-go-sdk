// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseFeatureSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetMseFeatureSwitchRequest
	GetAcceptLanguage() *string
}

type GetMseFeatureSwitchRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s GetMseFeatureSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMseFeatureSwitchRequest) GoString() string {
	return s.String()
}

func (s *GetMseFeatureSwitchRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetMseFeatureSwitchRequest) SetAcceptLanguage(v string) *GetMseFeatureSwitchRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetMseFeatureSwitchRequest) Validate() error {
	return dara.Validate(s)
}
