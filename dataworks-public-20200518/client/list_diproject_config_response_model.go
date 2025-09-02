// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIProjectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIProjectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIProjectConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListDIProjectConfigResponseBody) *ListDIProjectConfigResponse
	GetBody() *ListDIProjectConfigResponseBody
}

type ListDIProjectConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIProjectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIProjectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIProjectConfigResponse) GoString() string {
	return s.String()
}

func (s *ListDIProjectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIProjectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIProjectConfigResponse) GetBody() *ListDIProjectConfigResponseBody {
	return s.Body
}

func (s *ListDIProjectConfigResponse) SetHeaders(v map[string]*string) *ListDIProjectConfigResponse {
	s.Headers = v
	return s
}

func (s *ListDIProjectConfigResponse) SetStatusCode(v int32) *ListDIProjectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIProjectConfigResponse) SetBody(v *ListDIProjectConfigResponseBody) *ListDIProjectConfigResponse {
	s.Body = v
	return s
}

func (s *ListDIProjectConfigResponse) Validate() error {
	return dara.Validate(s)
}
