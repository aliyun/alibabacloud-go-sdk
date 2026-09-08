// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMarketingPreferencesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadAllMarketingPreferencesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadAllMarketingPreferencesResponse
	GetStatusCode() *int32
	SetBody(v *ReadAllMarketingPreferencesResponseBody) *ReadAllMarketingPreferencesResponse
	GetBody() *ReadAllMarketingPreferencesResponseBody
}

type ReadAllMarketingPreferencesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadAllMarketingPreferencesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadAllMarketingPreferencesResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMarketingPreferencesResponse) GoString() string {
	return s.String()
}

func (s *ReadAllMarketingPreferencesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadAllMarketingPreferencesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadAllMarketingPreferencesResponse) GetBody() *ReadAllMarketingPreferencesResponseBody {
	return s.Body
}

func (s *ReadAllMarketingPreferencesResponse) SetHeaders(v map[string]*string) *ReadAllMarketingPreferencesResponse {
	s.Headers = v
	return s
}

func (s *ReadAllMarketingPreferencesResponse) SetStatusCode(v int32) *ReadAllMarketingPreferencesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponse) SetBody(v *ReadAllMarketingPreferencesResponseBody) *ReadAllMarketingPreferencesResponse {
	s.Body = v
	return s
}

func (s *ReadAllMarketingPreferencesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
