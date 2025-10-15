// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordHistoryConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPasswordHistoryConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPasswordHistoryConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetPasswordHistoryConfigurationResponseBody) *GetPasswordHistoryConfigurationResponse
	GetBody() *GetPasswordHistoryConfigurationResponseBody
}

type GetPasswordHistoryConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordHistoryConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPasswordHistoryConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPasswordHistoryConfigurationResponse) GetBody() *GetPasswordHistoryConfigurationResponseBody {
	return s.Body
}

func (s *GetPasswordHistoryConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordHistoryConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordHistoryConfigurationResponse) SetStatusCode(v int32) *GetPasswordHistoryConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponse) SetBody(v *GetPasswordHistoryConfigurationResponseBody) *GetPasswordHistoryConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetPasswordHistoryConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
