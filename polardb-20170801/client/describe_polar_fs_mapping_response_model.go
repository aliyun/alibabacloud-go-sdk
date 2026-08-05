// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsMappingResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsMappingResponseBody) *DescribePolarFsMappingResponse
	GetBody() *DescribePolarFsMappingResponseBody
}

type DescribePolarFsMappingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsMappingResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsMappingResponse) GetBody() *DescribePolarFsMappingResponseBody {
	return s.Body
}

func (s *DescribePolarFsMappingResponse) SetHeaders(v map[string]*string) *DescribePolarFsMappingResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsMappingResponse) SetStatusCode(v int32) *DescribePolarFsMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsMappingResponse) SetBody(v *DescribePolarFsMappingResponseBody) *DescribePolarFsMappingResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
