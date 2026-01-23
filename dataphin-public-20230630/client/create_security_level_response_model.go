// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityLevelResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityLevelResponseBody) *CreateSecurityLevelResponse
	GetBody() *CreateSecurityLevelResponseBody
}

type CreateSecurityLevelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityLevelResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityLevelResponse) GetBody() *CreateSecurityLevelResponseBody {
	return s.Body
}

func (s *CreateSecurityLevelResponse) SetHeaders(v map[string]*string) *CreateSecurityLevelResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityLevelResponse) SetStatusCode(v int32) *CreateSecurityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityLevelResponse) SetBody(v *CreateSecurityLevelResponseBody) *CreateSecurityLevelResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
