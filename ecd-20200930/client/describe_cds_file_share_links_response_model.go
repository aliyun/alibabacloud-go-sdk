// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdsFileShareLinksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdsFileShareLinksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdsFileShareLinksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdsFileShareLinksResponseBody) *DescribeCdsFileShareLinksResponse
	GetBody() *DescribeCdsFileShareLinksResponseBody
}

type DescribeCdsFileShareLinksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdsFileShareLinksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdsFileShareLinksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdsFileShareLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdsFileShareLinksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdsFileShareLinksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdsFileShareLinksResponse) GetBody() *DescribeCdsFileShareLinksResponseBody {
	return s.Body
}

func (s *DescribeCdsFileShareLinksResponse) SetHeaders(v map[string]*string) *DescribeCdsFileShareLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdsFileShareLinksResponse) SetStatusCode(v int32) *DescribeCdsFileShareLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponse) SetBody(v *DescribeCdsFileShareLinksResponseBody) *DescribeCdsFileShareLinksResponse {
	s.Body = v
	return s
}

func (s *DescribeCdsFileShareLinksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
