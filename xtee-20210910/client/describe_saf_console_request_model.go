// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafConsoleRequest
	GetLang() *string
	SetContent(v string) *DescribeSafConsoleRequest
	GetContent() *string
	SetService(v string) *DescribeSafConsoleRequest
	GetService() *string
}

type DescribeSafConsoleRequest struct {
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s DescribeSafConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafConsoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafConsoleRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafConsoleRequest) GetContent() *string {
	return s.Content
}

func (s *DescribeSafConsoleRequest) GetService() *string {
	return s.Service
}

func (s *DescribeSafConsoleRequest) SetLang(v string) *DescribeSafConsoleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafConsoleRequest) SetContent(v string) *DescribeSafConsoleRequest {
	s.Content = &v
	return s
}

func (s *DescribeSafConsoleRequest) SetService(v string) *DescribeSafConsoleRequest {
	s.Service = &v
	return s
}

func (s *DescribeSafConsoleRequest) Validate() error {
	return dara.Validate(s)
}
