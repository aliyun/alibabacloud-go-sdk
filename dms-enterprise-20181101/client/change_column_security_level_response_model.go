// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecurityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeColumnSecurityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeColumnSecurityLevelResponse
	GetStatusCode() *int32
	SetBody(v *ChangeColumnSecurityLevelResponseBody) *ChangeColumnSecurityLevelResponse
	GetBody() *ChangeColumnSecurityLevelResponseBody
}

type ChangeColumnSecurityLevelResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeColumnSecurityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeColumnSecurityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecurityLevelResponse) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecurityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeColumnSecurityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeColumnSecurityLevelResponse) GetBody() *ChangeColumnSecurityLevelResponseBody {
	return s.Body
}

func (s *ChangeColumnSecurityLevelResponse) SetHeaders(v map[string]*string) *ChangeColumnSecurityLevelResponse {
	s.Headers = v
	return s
}

func (s *ChangeColumnSecurityLevelResponse) SetStatusCode(v int32) *ChangeColumnSecurityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeColumnSecurityLevelResponse) SetBody(v *ChangeColumnSecurityLevelResponseBody) *ChangeColumnSecurityLevelResponse {
	s.Body = v
	return s
}

func (s *ChangeColumnSecurityLevelResponse) Validate() error {
	return dara.Validate(s)
}
