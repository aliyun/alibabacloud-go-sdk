// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListRegionsRequest
	GetAcceptLanguage() *string
}

type ListRegionsRequest struct {
	// Language type. The range of values is as follows:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// - ja: Japanese
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
