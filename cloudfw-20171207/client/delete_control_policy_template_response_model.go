// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteControlPolicyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteControlPolicyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteControlPolicyTemplateResponseBody) *DeleteControlPolicyTemplateResponse
	GetBody() *DeleteControlPolicyTemplateResponseBody
}

type DeleteControlPolicyTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteControlPolicyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteControlPolicyTemplateResponse) GetBody() *DeleteControlPolicyTemplateResponseBody {
	return s.Body
}

func (s *DeleteControlPolicyTemplateResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyTemplateResponse) SetStatusCode(v int32) *DeleteControlPolicyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyTemplateResponse) SetBody(v *DeleteControlPolicyTemplateResponseBody) *DeleteControlPolicyTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteControlPolicyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
