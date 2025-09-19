// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileProtectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileProtectRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileProtectRuleResponseBody) *DeleteFileProtectRuleResponse
	GetBody() *DeleteFileProtectRuleResponseBody
}

type DeleteFileProtectRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileProtectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileProtectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileProtectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileProtectRuleResponse) GetBody() *DeleteFileProtectRuleResponseBody {
	return s.Body
}

func (s *DeleteFileProtectRuleResponse) SetHeaders(v map[string]*string) *DeleteFileProtectRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileProtectRuleResponse) SetStatusCode(v int32) *DeleteFileProtectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileProtectRuleResponse) SetBody(v *DeleteFileProtectRuleResponseBody) *DeleteFileProtectRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteFileProtectRuleResponse) Validate() error {
	return dara.Validate(s)
}
