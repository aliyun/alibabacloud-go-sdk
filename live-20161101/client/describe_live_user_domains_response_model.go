// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUserDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUserDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUserDomainsResponseBody) *DescribeLiveUserDomainsResponse
	GetBody() *DescribeLiveUserDomainsResponseBody
}

type DescribeLiveUserDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUserDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUserDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUserDomainsResponse) GetBody() *DescribeLiveUserDomainsResponseBody {
	return s.Body
}

func (s *DescribeLiveUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeLiveUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetStatusCode(v int32) *DescribeLiveUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetBody(v *DescribeLiveUserDomainsResponseBody) *DescribeLiveUserDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUserDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
