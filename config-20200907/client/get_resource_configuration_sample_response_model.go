// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceConfigurationSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceConfigurationSampleResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceConfigurationSampleResponseBody) *GetResourceConfigurationSampleResponse
	GetBody() *GetResourceConfigurationSampleResponseBody
}

type GetResourceConfigurationSampleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceConfigurationSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceConfigurationSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationSampleResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceConfigurationSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceConfigurationSampleResponse) GetBody() *GetResourceConfigurationSampleResponseBody {
	return s.Body
}

func (s *GetResourceConfigurationSampleResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationSampleResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationSampleResponse) SetStatusCode(v int32) *GetResourceConfigurationSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationSampleResponse) SetBody(v *GetResourceConfigurationSampleResponseBody) *GetResourceConfigurationSampleResponse {
	s.Body = v
	return s
}

func (s *GetResourceConfigurationSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
