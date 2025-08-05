// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetMmsDataSourceRequest
	GetLang() *string
	SetWithConfig(v bool) *GetMmsDataSourceRequest
	GetWithConfig() *bool
}

type GetMmsDataSourceRequest struct {
	// example:
	//
	// en_US
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// true
	WithConfig *bool `json:"withConfig,omitempty" xml:"withConfig,omitempty"`
}

func (s GetMmsDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *GetMmsDataSourceRequest) GetWithConfig() *bool {
	return s.WithConfig
}

func (s *GetMmsDataSourceRequest) SetLang(v string) *GetMmsDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *GetMmsDataSourceRequest) SetWithConfig(v bool) *GetMmsDataSourceRequest {
	s.WithConfig = &v
	return s
}

func (s *GetMmsDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
