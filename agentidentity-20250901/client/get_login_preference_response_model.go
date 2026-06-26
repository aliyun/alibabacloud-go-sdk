// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginPreferenceResponseBody) *GetLoginPreferenceResponse
	GetBody() *GetLoginPreferenceResponseBody
}

type GetLoginPreferenceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginPreferenceResponse) GetBody() *GetLoginPreferenceResponseBody {
	return s.Body
}

func (s *GetLoginPreferenceResponse) SetHeaders(v map[string]*string) *GetLoginPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetLoginPreferenceResponse) SetStatusCode(v int32) *GetLoginPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginPreferenceResponse) SetBody(v *GetLoginPreferenceResponseBody) *GetLoginPreferenceResponse {
	s.Body = v
	return s
}

func (s *GetLoginPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
