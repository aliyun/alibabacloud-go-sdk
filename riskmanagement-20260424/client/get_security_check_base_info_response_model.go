// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityCheckBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityCheckBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityCheckBaseInfoResponseBody) *GetSecurityCheckBaseInfoResponse
	GetBody() *GetSecurityCheckBaseInfoResponseBody
}

type GetSecurityCheckBaseInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityCheckBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityCheckBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityCheckBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityCheckBaseInfoResponse) GetBody() *GetSecurityCheckBaseInfoResponseBody {
	return s.Body
}

func (s *GetSecurityCheckBaseInfoResponse) SetHeaders(v map[string]*string) *GetSecurityCheckBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityCheckBaseInfoResponse) SetStatusCode(v int32) *GetSecurityCheckBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponse) SetBody(v *GetSecurityCheckBaseInfoResponseBody) *GetSecurityCheckBaseInfoResponse {
	s.Body = v
	return s
}

func (s *GetSecurityCheckBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
