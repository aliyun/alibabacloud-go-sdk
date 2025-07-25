// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePdnsAppKeysRequest
	GetLang() *string
}

type DescribePdnsAppKeysRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePdnsAppKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeysRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsAppKeysRequest) SetLang(v string) *DescribePdnsAppKeysRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsAppKeysRequest) Validate() error {
	return dara.Validate(s)
}
