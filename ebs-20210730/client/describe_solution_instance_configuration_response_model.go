// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSolutionInstanceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSolutionInstanceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSolutionInstanceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSolutionInstanceConfigurationResponseBody) *DescribeSolutionInstanceConfigurationResponse
	GetBody() *DescribeSolutionInstanceConfigurationResponseBody
}

type DescribeSolutionInstanceConfigurationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSolutionInstanceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSolutionInstanceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSolutionInstanceConfigurationResponse) GetBody() *DescribeSolutionInstanceConfigurationResponseBody {
	return s.Body
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetHeaders(v map[string]*string) *DescribeSolutionInstanceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetStatusCode(v int32) *DescribeSolutionInstanceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponse) SetBody(v *DescribeSolutionInstanceConfigurationResponseBody) *DescribeSolutionInstanceConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
