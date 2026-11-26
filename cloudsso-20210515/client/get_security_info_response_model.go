// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityInfoResponseBody) *GetSecurityInfoResponse
	GetBody() *GetSecurityInfoResponseBody
}

type GetSecurityInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityInfoResponse) GetBody() *GetSecurityInfoResponseBody {
	return s.Body
}

func (s *GetSecurityInfoResponse) SetHeaders(v map[string]*string) *GetSecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityInfoResponse) SetStatusCode(v int32) *GetSecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityInfoResponse) SetBody(v *GetSecurityInfoResponseBody) *GetSecurityInfoResponse {
	s.Body = v
	return s
}

func (s *GetSecurityInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
