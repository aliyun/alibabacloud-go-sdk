// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSourceResponse
	GetStatusCode() *int32
	SetBody(v *SaveSourceResponseBody) *SaveSourceResponse
	GetBody() *SaveSourceResponseBody
}

type SaveSourceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSourceResponse) GoString() string {
	return s.String()
}

func (s *SaveSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSourceResponse) GetBody() *SaveSourceResponseBody {
	return s.Body
}

func (s *SaveSourceResponse) SetHeaders(v map[string]*string) *SaveSourceResponse {
	s.Headers = v
	return s
}

func (s *SaveSourceResponse) SetStatusCode(v int32) *SaveSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSourceResponse) SetBody(v *SaveSourceResponseBody) *SaveSourceResponse {
	s.Body = v
	return s
}

func (s *SaveSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
