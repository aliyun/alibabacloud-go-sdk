// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlbSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QuerySlbSpecRequest
	GetAcceptLanguage() *string
}

type QuerySlbSpecRequest struct {
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

func (s QuerySlbSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySlbSpecRequest) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QuerySlbSpecRequest) SetAcceptLanguage(v string) *QuerySlbSpecRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QuerySlbSpecRequest) Validate() error {
	return dara.Validate(s)
}
