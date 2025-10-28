// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserConfigsResponse
	GetStatusCode() *int32
	SetBody(v *SetUserConfigsResponseBody) *SetUserConfigsResponse
	GetBody() *SetUserConfigsResponseBody
}

type SetUserConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *SetUserConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserConfigsResponse) GetBody() *SetUserConfigsResponseBody {
	return s.Body
}

func (s *SetUserConfigsResponse) SetHeaders(v map[string]*string) *SetUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *SetUserConfigsResponse) SetStatusCode(v int32) *SetUserConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserConfigsResponse) SetBody(v *SetUserConfigsResponseBody) *SetUserConfigsResponse {
	s.Body = v
	return s
}

func (s *SetUserConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
