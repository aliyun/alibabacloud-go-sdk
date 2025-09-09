// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAccountPasswordResponseBody) *ChangeAccountPasswordResponse
	GetBody() *ChangeAccountPasswordResponseBody
}

type ChangeAccountPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAccountPasswordResponse) GetBody() *ChangeAccountPasswordResponseBody {
	return s.Body
}

func (s *ChangeAccountPasswordResponse) SetHeaders(v map[string]*string) *ChangeAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountPasswordResponse) SetStatusCode(v int32) *ChangeAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountPasswordResponse) SetBody(v *ChangeAccountPasswordResponseBody) *ChangeAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ChangeAccountPasswordResponse) Validate() error {
	return dara.Validate(s)
}
