// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMarketingPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMarketingPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMarketingPreferenceResponseBody) *UpdateMarketingPreferenceResponse
	GetBody() *UpdateMarketingPreferenceResponseBody
}

type UpdateMarketingPreferenceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMarketingPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMarketingPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingPreferenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateMarketingPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMarketingPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMarketingPreferenceResponse) GetBody() *UpdateMarketingPreferenceResponseBody {
	return s.Body
}

func (s *UpdateMarketingPreferenceResponse) SetHeaders(v map[string]*string) *UpdateMarketingPreferenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateMarketingPreferenceResponse) SetStatusCode(v int32) *UpdateMarketingPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMarketingPreferenceResponse) SetBody(v *UpdateMarketingPreferenceResponseBody) *UpdateMarketingPreferenceResponse {
	s.Body = v
	return s
}

func (s *UpdateMarketingPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
