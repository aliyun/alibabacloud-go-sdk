// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlFixPriceDomainListUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListDate(v string) *GetIntlFixPriceDomainListUrlRequest
	GetListDate() *string
}

type GetIntlFixPriceDomainListUrlRequest struct {
	// The date when the list is exported.
	//
	// example:
	//
	// 20240809
	ListDate *string `json:"ListDate,omitempty" xml:"ListDate,omitempty"`
}

func (s GetIntlFixPriceDomainListUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntlFixPriceDomainListUrlRequest) GoString() string {
	return s.String()
}

func (s *GetIntlFixPriceDomainListUrlRequest) GetListDate() *string {
	return s.ListDate
}

func (s *GetIntlFixPriceDomainListUrlRequest) SetListDate(v string) *GetIntlFixPriceDomainListUrlRequest {
	s.ListDate = &v
	return s
}

func (s *GetIntlFixPriceDomainListUrlRequest) Validate() error {
	return dara.Validate(s)
}
