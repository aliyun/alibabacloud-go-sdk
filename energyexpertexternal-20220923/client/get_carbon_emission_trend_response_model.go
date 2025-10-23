// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCarbonEmissionTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCarbonEmissionTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCarbonEmissionTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetCarbonEmissionTrendResponseBody) *GetCarbonEmissionTrendResponse
	GetBody() *GetCarbonEmissionTrendResponseBody
}

type GetCarbonEmissionTrendResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCarbonEmissionTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCarbonEmissionTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendResponse) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCarbonEmissionTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCarbonEmissionTrendResponse) GetBody() *GetCarbonEmissionTrendResponseBody {
	return s.Body
}

func (s *GetCarbonEmissionTrendResponse) SetHeaders(v map[string]*string) *GetCarbonEmissionTrendResponse {
	s.Headers = v
	return s
}

func (s *GetCarbonEmissionTrendResponse) SetStatusCode(v int32) *GetCarbonEmissionTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCarbonEmissionTrendResponse) SetBody(v *GetCarbonEmissionTrendResponseBody) *GetCarbonEmissionTrendResponse {
	s.Body = v
	return s
}

func (s *GetCarbonEmissionTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
