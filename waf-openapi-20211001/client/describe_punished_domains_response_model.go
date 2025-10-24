// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePunishedDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePunishedDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePunishedDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePunishedDomainsResponseBody) *DescribePunishedDomainsResponse
	GetBody() *DescribePunishedDomainsResponseBody
}

type DescribePunishedDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePunishedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePunishedDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePunishedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePunishedDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePunishedDomainsResponse) GetBody() *DescribePunishedDomainsResponseBody {
	return s.Body
}

func (s *DescribePunishedDomainsResponse) SetHeaders(v map[string]*string) *DescribePunishedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribePunishedDomainsResponse) SetStatusCode(v int32) *DescribePunishedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePunishedDomainsResponse) SetBody(v *DescribePunishedDomainsResponseBody) *DescribePunishedDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribePunishedDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
