// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeprecatedTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeprecatedTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeprecatedTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDeprecatedTemplatesResponseBody) *ListDeprecatedTemplatesResponse
	GetBody() *ListDeprecatedTemplatesResponseBody
}

type ListDeprecatedTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeprecatedTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeprecatedTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeprecatedTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeprecatedTemplatesResponse) GetBody() *ListDeprecatedTemplatesResponseBody {
	return s.Body
}

func (s *ListDeprecatedTemplatesResponse) SetHeaders(v map[string]*string) *ListDeprecatedTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDeprecatedTemplatesResponse) SetStatusCode(v int32) *ListDeprecatedTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeprecatedTemplatesResponse) SetBody(v *ListDeprecatedTemplatesResponseBody) *ListDeprecatedTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDeprecatedTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
