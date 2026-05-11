// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTTSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTTSConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTTSConfigResponseBody) *DescribeTTSConfigResponse
	GetBody() *DescribeTTSConfigResponseBody
}

type DescribeTTSConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTTSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTTSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTTSConfigResponse) GetBody() *DescribeTTSConfigResponseBody {
	return s.Body
}

func (s *DescribeTTSConfigResponse) SetHeaders(v map[string]*string) *DescribeTTSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeTTSConfigResponse) SetStatusCode(v int32) *DescribeTTSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTTSConfigResponse) SetBody(v *DescribeTTSConfigResponseBody) *DescribeTTSConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeTTSConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
