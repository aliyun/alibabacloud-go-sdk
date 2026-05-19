// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectClientRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileProtectClientRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileProtectClientRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileProtectClientRuleResponseBody) *DeleteFileProtectClientRuleResponse
	GetBody() *DeleteFileProtectClientRuleResponseBody
}

type DeleteFileProtectClientRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileProtectClientRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileProtectClientRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectClientRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectClientRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileProtectClientRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileProtectClientRuleResponse) GetBody() *DeleteFileProtectClientRuleResponseBody {
	return s.Body
}

func (s *DeleteFileProtectClientRuleResponse) SetHeaders(v map[string]*string) *DeleteFileProtectClientRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileProtectClientRuleResponse) SetStatusCode(v int32) *DeleteFileProtectClientRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileProtectClientRuleResponse) SetBody(v *DeleteFileProtectClientRuleResponseBody) *DeleteFileProtectClientRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteFileProtectClientRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
