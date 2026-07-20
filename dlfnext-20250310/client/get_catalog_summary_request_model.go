// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetCatalogSummaryRequest
	GetDate() *string
}

type GetCatalogSummaryRequest struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
}

func (s GetCatalogSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryRequest) GetDate() *string {
	return s.Date
}

func (s *GetCatalogSummaryRequest) SetDate(v string) *GetCatalogSummaryRequest {
	s.Date = &v
	return s
}

func (s *GetCatalogSummaryRequest) Validate() error {
	return dara.Validate(s)
}
