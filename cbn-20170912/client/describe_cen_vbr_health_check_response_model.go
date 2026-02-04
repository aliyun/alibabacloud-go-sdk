// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenVbrHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenVbrHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenVbrHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenVbrHealthCheckResponseBody) *DescribeCenVbrHealthCheckResponse
	GetBody() *DescribeCenVbrHealthCheckResponseBody
}

type DescribeCenVbrHealthCheckResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenVbrHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenVbrHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenVbrHealthCheckResponse) GetBody() *DescribeCenVbrHealthCheckResponseBody {
	return s.Body
}

func (s *DescribeCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *DescribeCenVbrHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenVbrHealthCheckResponse) SetStatusCode(v int32) *DescribeCenVbrHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponse) SetBody(v *DescribeCenVbrHealthCheckResponseBody) *DescribeCenVbrHealthCheckResponse {
	s.Body = v
	return s
}

func (s *DescribeCenVbrHealthCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
