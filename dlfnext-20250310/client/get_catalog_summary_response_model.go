// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogSummaryResponse
	GetStatusCode() *int32
	SetBody(v *CatalogSummary) *GetCatalogSummaryResponse
	GetBody() *CatalogSummary
}

type GetCatalogSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogSummary    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogSummaryResponse) GetBody() *CatalogSummary {
	return s.Body
}

func (s *GetCatalogSummaryResponse) SetHeaders(v map[string]*string) *GetCatalogSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogSummaryResponse) SetStatusCode(v int32) *GetCatalogSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogSummaryResponse) SetBody(v *CatalogSummary) *GetCatalogSummaryResponse {
	s.Body = v
	return s
}

func (s *GetCatalogSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
