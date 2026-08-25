// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoginPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoginPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *SetLoginPreferenceResponseBody) *SetLoginPreferenceResponse
	GetBody() *SetLoginPreferenceResponseBody
}

type SetLoginPreferenceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoginPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoginPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoginPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoginPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoginPreferenceResponse) GetBody() *SetLoginPreferenceResponseBody {
	return s.Body
}

func (s *SetLoginPreferenceResponse) SetHeaders(v map[string]*string) *SetLoginPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetLoginPreferenceResponse) SetStatusCode(v int32) *SetLoginPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoginPreferenceResponse) SetBody(v *SetLoginPreferenceResponseBody) *SetLoginPreferenceResponse {
	s.Body = v
	return s
}

func (s *SetLoginPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
