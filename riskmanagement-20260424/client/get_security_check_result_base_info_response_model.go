// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckResultBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityCheckResultBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityCheckResultBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityCheckResultBaseInfoResponseBody) *GetSecurityCheckResultBaseInfoResponse
	GetBody() *GetSecurityCheckResultBaseInfoResponseBody
}

type GetSecurityCheckResultBaseInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityCheckResultBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityCheckResultBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckResultBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckResultBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityCheckResultBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityCheckResultBaseInfoResponse) GetBody() *GetSecurityCheckResultBaseInfoResponseBody {
	return s.Body
}

func (s *GetSecurityCheckResultBaseInfoResponse) SetHeaders(v map[string]*string) *GetSecurityCheckResultBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponse) SetStatusCode(v int32) *GetSecurityCheckResultBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponse) SetBody(v *GetSecurityCheckResultBaseInfoResponseBody) *GetSecurityCheckResultBaseInfoResponse {
	s.Body = v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
