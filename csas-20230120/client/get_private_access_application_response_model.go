// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetPrivateAccessApplicationResponseBody) *GetPrivateAccessApplicationResponse
	GetBody() *GetPrivateAccessApplicationResponseBody
}

type GetPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrivateAccessApplicationResponse) GetBody() *GetPrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *GetPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *GetPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateAccessApplicationResponse) SetStatusCode(v int32) *GetPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateAccessApplicationResponse) SetBody(v *GetPrivateAccessApplicationResponseBody) *GetPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *GetPrivateAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
