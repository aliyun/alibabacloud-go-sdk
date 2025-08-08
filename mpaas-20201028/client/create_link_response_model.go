// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateLinkResponseBody) *CreateLinkResponse
	GetBody() *CreateLinkResponseBody
}

type CreateLinkResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLinkResponse) GetBody() *CreateLinkResponseBody {
	return s.Body
}

func (s *CreateLinkResponse) SetHeaders(v map[string]*string) *CreateLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateLinkResponse) SetStatusCode(v int32) *CreateLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLinkResponse) SetBody(v *CreateLinkResponseBody) *CreateLinkResponse {
	s.Body = v
	return s
}

func (s *CreateLinkResponse) Validate() error {
	return dara.Validate(s)
}
