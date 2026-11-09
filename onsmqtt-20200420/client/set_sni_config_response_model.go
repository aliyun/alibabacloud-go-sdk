// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSniConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSniConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSniConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetSniConfigResponseBody) *SetSniConfigResponse
	GetBody() *SetSniConfigResponseBody
}

type SetSniConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSniConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSniConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSniConfigResponse) GoString() string {
	return s.String()
}

func (s *SetSniConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSniConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSniConfigResponse) GetBody() *SetSniConfigResponseBody {
	return s.Body
}

func (s *SetSniConfigResponse) SetHeaders(v map[string]*string) *SetSniConfigResponse {
	s.Headers = v
	return s
}

func (s *SetSniConfigResponse) SetStatusCode(v int32) *SetSniConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSniConfigResponse) SetBody(v *SetSniConfigResponseBody) *SetSniConfigResponse {
	s.Body = v
	return s
}

func (s *SetSniConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
