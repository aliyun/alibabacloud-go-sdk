// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrTemplatesResponseBody) *DescribeCrTemplatesResponse
	GetBody() *DescribeCrTemplatesResponseBody
}

type DescribeCrTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrTemplatesResponse) GetBody() *DescribeCrTemplatesResponseBody {
	return s.Body
}

func (s *DescribeCrTemplatesResponse) SetHeaders(v map[string]*string) *DescribeCrTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrTemplatesResponse) SetStatusCode(v int32) *DescribeCrTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrTemplatesResponse) SetBody(v *DescribeCrTemplatesResponseBody) *DescribeCrTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeCrTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
