// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMarketingPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMarketingPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMarketingPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *ReadMarketingPreferenceResponseBody) *ReadMarketingPreferenceResponse
	GetBody() *ReadMarketingPreferenceResponseBody
}

type ReadMarketingPreferenceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMarketingPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMarketingPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMarketingPreferenceResponse) GoString() string {
	return s.String()
}

func (s *ReadMarketingPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMarketingPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMarketingPreferenceResponse) GetBody() *ReadMarketingPreferenceResponseBody {
	return s.Body
}

func (s *ReadMarketingPreferenceResponse) SetHeaders(v map[string]*string) *ReadMarketingPreferenceResponse {
	s.Headers = v
	return s
}

func (s *ReadMarketingPreferenceResponse) SetStatusCode(v int32) *ReadMarketingPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMarketingPreferenceResponse) SetBody(v *ReadMarketingPreferenceResponseBody) *ReadMarketingPreferenceResponse {
	s.Body = v
	return s
}

func (s *ReadMarketingPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
