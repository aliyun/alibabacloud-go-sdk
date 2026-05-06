// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutomaticFrequencyControlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAutomaticFrequencyControlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAutomaticFrequencyControlConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetAutomaticFrequencyControlConfigResponseBody) *SetAutomaticFrequencyControlConfigResponse
	GetBody() *SetAutomaticFrequencyControlConfigResponseBody
}

type SetAutomaticFrequencyControlConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAutomaticFrequencyControlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAutomaticFrequencyControlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAutomaticFrequencyControlConfigResponse) GoString() string {
	return s.String()
}

func (s *SetAutomaticFrequencyControlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAutomaticFrequencyControlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAutomaticFrequencyControlConfigResponse) GetBody() *SetAutomaticFrequencyControlConfigResponseBody {
	return s.Body
}

func (s *SetAutomaticFrequencyControlConfigResponse) SetHeaders(v map[string]*string) *SetAutomaticFrequencyControlConfigResponse {
	s.Headers = v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponse) SetStatusCode(v int32) *SetAutomaticFrequencyControlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponse) SetBody(v *SetAutomaticFrequencyControlConfigResponseBody) *SetAutomaticFrequencyControlConfigResponse {
	s.Body = v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
