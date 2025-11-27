// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePGHbaConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePGHbaConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePGHbaConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribePGHbaConfigResponseBody) *DescribePGHbaConfigResponse
	GetBody() *DescribePGHbaConfigResponseBody
}

type DescribePGHbaConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePGHbaConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePGHbaConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePGHbaConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePGHbaConfigResponse) GetBody() *DescribePGHbaConfigResponseBody {
	return s.Body
}

func (s *DescribePGHbaConfigResponse) SetHeaders(v map[string]*string) *DescribePGHbaConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePGHbaConfigResponse) SetStatusCode(v int32) *DescribePGHbaConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePGHbaConfigResponse) SetBody(v *DescribePGHbaConfigResponseBody) *DescribePGHbaConfigResponse {
	s.Body = v
	return s
}

func (s *DescribePGHbaConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
