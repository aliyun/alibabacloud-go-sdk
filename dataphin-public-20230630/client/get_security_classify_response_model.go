// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityClassifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityClassifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityClassifyResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityClassifyResponseBody) *GetSecurityClassifyResponse
	GetBody() *GetSecurityClassifyResponseBody
}

type GetSecurityClassifyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityClassifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityClassifyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityClassifyResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityClassifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityClassifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityClassifyResponse) GetBody() *GetSecurityClassifyResponseBody {
	return s.Body
}

func (s *GetSecurityClassifyResponse) SetHeaders(v map[string]*string) *GetSecurityClassifyResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityClassifyResponse) SetStatusCode(v int32) *GetSecurityClassifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityClassifyResponse) SetBody(v *GetSecurityClassifyResponseBody) *GetSecurityClassifyResponse {
	s.Body = v
	return s
}

func (s *GetSecurityClassifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
