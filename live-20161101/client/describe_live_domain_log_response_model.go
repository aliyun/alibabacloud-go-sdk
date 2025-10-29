// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainLogResponseBody) *DescribeLiveDomainLogResponse
	GetBody() *DescribeLiveDomainLogResponseBody
}

type DescribeLiveDomainLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainLogResponse) GetBody() *DescribeLiveDomainLogResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainLogResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainLogResponse) SetStatusCode(v int32) *DescribeLiveDomainLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainLogResponse) SetBody(v *DescribeLiveDomainLogResponseBody) *DescribeLiveDomainLogResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
