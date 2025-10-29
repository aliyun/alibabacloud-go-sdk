// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageTokenResponseBody) *GetMessageTokenResponse
	GetBody() *GetMessageTokenResponseBody
}

type GetMessageTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageTokenResponse) GoString() string {
	return s.String()
}

func (s *GetMessageTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageTokenResponse) GetBody() *GetMessageTokenResponseBody {
	return s.Body
}

func (s *GetMessageTokenResponse) SetHeaders(v map[string]*string) *GetMessageTokenResponse {
	s.Headers = v
	return s
}

func (s *GetMessageTokenResponse) SetStatusCode(v int32) *GetMessageTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageTokenResponse) SetBody(v *GetMessageTokenResponseBody) *GetMessageTokenResponse {
	s.Body = v
	return s
}

func (s *GetMessageTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
