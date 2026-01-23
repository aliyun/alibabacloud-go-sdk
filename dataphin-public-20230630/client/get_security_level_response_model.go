// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityLevelResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityLevelResponseBody) *GetSecurityLevelResponse
	GetBody() *GetSecurityLevelResponseBody
}

type GetSecurityLevelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityLevelResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityLevelResponse) GetBody() *GetSecurityLevelResponseBody {
	return s.Body
}

func (s *GetSecurityLevelResponse) SetHeaders(v map[string]*string) *GetSecurityLevelResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityLevelResponse) SetStatusCode(v int32) *GetSecurityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityLevelResponse) SetBody(v *GetSecurityLevelResponseBody) *GetSecurityLevelResponse {
	s.Body = v
	return s
}

func (s *GetSecurityLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
