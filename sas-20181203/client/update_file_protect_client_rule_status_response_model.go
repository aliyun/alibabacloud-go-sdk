// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectClientRuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectClientRuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectClientRuleStatusResponseBody) *UpdateFileProtectClientRuleStatusResponse
	GetBody() *UpdateFileProtectClientRuleStatusResponseBody
}

type UpdateFileProtectClientRuleStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectClientRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectClientRuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectClientRuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectClientRuleStatusResponse) GetBody() *UpdateFileProtectClientRuleStatusResponseBody {
	return s.Body
}

func (s *UpdateFileProtectClientRuleStatusResponse) SetHeaders(v map[string]*string) *UpdateFileProtectClientRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectClientRuleStatusResponse) SetStatusCode(v int32) *UpdateFileProtectClientRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusResponse) SetBody(v *UpdateFileProtectClientRuleStatusResponseBody) *UpdateFileProtectClientRuleStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectClientRuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
