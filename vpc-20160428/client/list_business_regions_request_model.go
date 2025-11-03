// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListBusinessRegionsRequest
	GetAcceptLanguage() *string
}

type ListBusinessRegionsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListBusinessRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListBusinessRegionsRequest) SetAcceptLanguage(v string) *ListBusinessRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListBusinessRegionsRequest) Validate() error {
	return dara.Validate(s)
}
