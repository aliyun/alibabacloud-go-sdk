// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivateAccessApplicationResponseBody) *CreatePrivateAccessApplicationResponse
	GetBody() *CreatePrivateAccessApplicationResponseBody
}

type CreatePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivateAccessApplicationResponse) GetBody() *CreatePrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *CreatePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessApplicationResponse) SetStatusCode(v int32) *CreatePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessApplicationResponse) SetBody(v *CreatePrivateAccessApplicationResponseBody) *CreatePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *CreatePrivateAccessApplicationResponse) Validate() error {
	return dara.Validate(s)
}
