// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutomaticFrequencyControlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutomaticFrequencyControlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutomaticFrequencyControlConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAutomaticFrequencyControlConfigResponseBody) *GetAutomaticFrequencyControlConfigResponse
	GetBody() *GetAutomaticFrequencyControlConfigResponseBody
}

type GetAutomaticFrequencyControlConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutomaticFrequencyControlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutomaticFrequencyControlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutomaticFrequencyControlConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutomaticFrequencyControlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutomaticFrequencyControlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutomaticFrequencyControlConfigResponse) GetBody() *GetAutomaticFrequencyControlConfigResponseBody {
	return s.Body
}

func (s *GetAutomaticFrequencyControlConfigResponse) SetHeaders(v map[string]*string) *GetAutomaticFrequencyControlConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponse) SetStatusCode(v int32) *GetAutomaticFrequencyControlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponse) SetBody(v *GetAutomaticFrequencyControlConfigResponseBody) *GetAutomaticFrequencyControlConfigResponse {
	s.Body = v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
