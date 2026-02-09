// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetTraceConfigRequest
	GetAcceptLanguage() *string
}

type GetTraceConfigRequest struct {
	// The language in which you want results to be returned. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
}

func (s GetTraceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetTraceConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetTraceConfigRequest) SetAcceptLanguage(v string) *GetTraceConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetTraceConfigRequest) Validate() error {
	return dara.Validate(s)
}
