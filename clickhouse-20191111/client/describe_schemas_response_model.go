// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSchemasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse
	GetBody() *DescribeSchemasResponseBody
}

type DescribeSchemasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSchemasResponse) GetBody() *DescribeSchemasResponseBody {
	return s.Body
}

func (s *DescribeSchemasResponse) SetHeaders(v map[string]*string) *DescribeSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchemasResponse) SetStatusCode(v int32) *DescribeSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSchemasResponse) SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse {
	s.Body = v
	return s
}

func (s *DescribeSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
