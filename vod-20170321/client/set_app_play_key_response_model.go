// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppPlayKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAppPlayKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAppPlayKeyResponse
	GetStatusCode() *int32
	SetBody(v *SetAppPlayKeyResponseBody) *SetAppPlayKeyResponse
	GetBody() *SetAppPlayKeyResponseBody
}

type SetAppPlayKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAppPlayKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAppPlayKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAppPlayKeyResponse) GoString() string {
	return s.String()
}

func (s *SetAppPlayKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAppPlayKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAppPlayKeyResponse) GetBody() *SetAppPlayKeyResponseBody {
	return s.Body
}

func (s *SetAppPlayKeyResponse) SetHeaders(v map[string]*string) *SetAppPlayKeyResponse {
	s.Headers = v
	return s
}

func (s *SetAppPlayKeyResponse) SetStatusCode(v int32) *SetAppPlayKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAppPlayKeyResponse) SetBody(v *SetAppPlayKeyResponseBody) *SetAppPlayKeyResponse {
	s.Body = v
	return s
}

func (s *SetAppPlayKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
