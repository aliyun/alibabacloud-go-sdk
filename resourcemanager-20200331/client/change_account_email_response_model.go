// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAccountEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAccountEmailResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAccountEmailResponseBody) *ChangeAccountEmailResponse
	GetBody() *ChangeAccountEmailResponseBody
}

type ChangeAccountEmailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAccountEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAccountEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAccountEmailResponse) GetBody() *ChangeAccountEmailResponseBody {
	return s.Body
}

func (s *ChangeAccountEmailResponse) SetHeaders(v map[string]*string) *ChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountEmailResponse) SetStatusCode(v int32) *ChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountEmailResponse) SetBody(v *ChangeAccountEmailResponseBody) *ChangeAccountEmailResponse {
	s.Body = v
	return s
}

func (s *ChangeAccountEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
