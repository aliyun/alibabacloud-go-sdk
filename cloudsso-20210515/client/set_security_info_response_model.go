// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSecurityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSecurityInfoResponse
	GetStatusCode() *int32
	SetBody(v *SetSecurityInfoResponseBody) *SetSecurityInfoResponse
	GetBody() *SetSecurityInfoResponseBody
}

type SetSecurityInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecurityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSecurityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSecurityInfoResponse) GetBody() *SetSecurityInfoResponseBody {
	return s.Body
}

func (s *SetSecurityInfoResponse) SetHeaders(v map[string]*string) *SetSecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *SetSecurityInfoResponse) SetStatusCode(v int32) *SetSecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecurityInfoResponse) SetBody(v *SetSecurityInfoResponseBody) *SetSecurityInfoResponse {
	s.Body = v
	return s
}

func (s *SetSecurityInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
