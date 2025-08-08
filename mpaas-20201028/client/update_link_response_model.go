// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLinkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLinkResponseBody) *UpdateLinkResponse
	GetBody() *UpdateLinkResponseBody
}

type UpdateLinkResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLinkResponse) GetBody() *UpdateLinkResponseBody {
	return s.Body
}

func (s *UpdateLinkResponse) SetHeaders(v map[string]*string) *UpdateLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateLinkResponse) SetStatusCode(v int32) *UpdateLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLinkResponse) SetBody(v *UpdateLinkResponseBody) *UpdateLinkResponse {
	s.Body = v
	return s
}

func (s *UpdateLinkResponse) Validate() error {
	return dara.Validate(s)
}
