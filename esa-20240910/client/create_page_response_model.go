// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePageResponse
	GetStatusCode() *int32
	SetBody(v *CreatePageResponseBody) *CreatePageResponse
	GetBody() *CreatePageResponseBody
}

type CreatePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePageResponse) GoString() string {
	return s.String()
}

func (s *CreatePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePageResponse) GetBody() *CreatePageResponseBody {
	return s.Body
}

func (s *CreatePageResponse) SetHeaders(v map[string]*string) *CreatePageResponse {
	s.Headers = v
	return s
}

func (s *CreatePageResponse) SetStatusCode(v int32) *CreatePageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePageResponse) SetBody(v *CreatePageResponseBody) *CreatePageResponse {
	s.Body = v
	return s
}

func (s *CreatePageResponse) Validate() error {
	return dara.Validate(s)
}
