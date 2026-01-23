// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityLevelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityLevelResponseBody) *DeleteSecurityLevelResponse
	GetBody() *DeleteSecurityLevelResponseBody
}

type DeleteSecurityLevelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityLevelResponse) GetBody() *DeleteSecurityLevelResponseBody {
	return s.Body
}

func (s *DeleteSecurityLevelResponse) SetHeaders(v map[string]*string) *DeleteSecurityLevelResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityLevelResponse) SetStatusCode(v int32) *DeleteSecurityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityLevelResponse) SetBody(v *DeleteSecurityLevelResponseBody) *DeleteSecurityLevelResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
