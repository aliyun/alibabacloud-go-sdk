// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuperResolveVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuperResolveVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuperResolveVideoResponse
	GetStatusCode() *int32
	SetBody(v *SuperResolveVideoResponseBody) *SuperResolveVideoResponse
	GetBody() *SuperResolveVideoResponseBody
}

type SuperResolveVideoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuperResolveVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuperResolveVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s SuperResolveVideoResponse) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuperResolveVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuperResolveVideoResponse) GetBody() *SuperResolveVideoResponseBody {
	return s.Body
}

func (s *SuperResolveVideoResponse) SetHeaders(v map[string]*string) *SuperResolveVideoResponse {
	s.Headers = v
	return s
}

func (s *SuperResolveVideoResponse) SetStatusCode(v int32) *SuperResolveVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *SuperResolveVideoResponse) SetBody(v *SuperResolveVideoResponseBody) *SuperResolveVideoResponse {
	s.Body = v
	return s
}

func (s *SuperResolveVideoResponse) Validate() error {
	return dara.Validate(s)
}
