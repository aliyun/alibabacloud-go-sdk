// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogSummaryTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogSummaryTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogSummaryTrendResponse
	GetStatusCode() *int32
	SetBody(v *CatalogSummaryTrend) *GetCatalogSummaryTrendResponse
	GetBody() *CatalogSummaryTrend
}

type GetCatalogSummaryTrendResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogSummaryTrend `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogSummaryTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogSummaryTrendResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogSummaryTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogSummaryTrendResponse) GetBody() *CatalogSummaryTrend {
	return s.Body
}

func (s *GetCatalogSummaryTrendResponse) SetHeaders(v map[string]*string) *GetCatalogSummaryTrendResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogSummaryTrendResponse) SetStatusCode(v int32) *GetCatalogSummaryTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogSummaryTrendResponse) SetBody(v *CatalogSummaryTrend) *GetCatalogSummaryTrendResponse {
	s.Body = v
	return s
}

func (s *GetCatalogSummaryTrendResponse) Validate() error {
	return dara.Validate(s)
}
