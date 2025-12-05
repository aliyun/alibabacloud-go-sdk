// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceWhiteListResponseBody) *ConfigInstanceWhiteListResponse
	GetBody() *ConfigInstanceWhiteListResponseBody
}

type ConfigInstanceWhiteListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceWhiteListResponse) GetBody() *ConfigInstanceWhiteListResponseBody {
	return s.Body
}

func (s *ConfigInstanceWhiteListResponse) SetHeaders(v map[string]*string) *ConfigInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceWhiteListResponse) SetStatusCode(v int32) *ConfigInstanceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceWhiteListResponse) SetBody(v *ConfigInstanceWhiteListResponseBody) *ConfigInstanceWhiteListResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
