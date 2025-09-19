// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetHoneypotStatisticsRequest
	GetLang() *string
}

type GetHoneypotStatisticsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetHoneypotStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotStatisticsRequest) SetLang(v string) *GetHoneypotStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
