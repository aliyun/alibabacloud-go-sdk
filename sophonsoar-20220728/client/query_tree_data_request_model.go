// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTreeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryTreeDataRequest
	GetLang() *string
}

type QueryTreeDataRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryTreeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTreeDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTreeDataRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTreeDataRequest) SetLang(v string) *QueryTreeDataRequest {
	s.Lang = &v
	return s
}

func (s *QueryTreeDataRequest) Validate() error {
	return dara.Validate(s)
}
