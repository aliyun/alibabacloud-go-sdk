// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceDLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceDLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceDLinkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceDLinkResponseBody) *DeleteResourceDLinkResponse
	GetBody() *DeleteResourceDLinkResponseBody
}

type DeleteResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceDLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceDLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceDLinkResponse) GetBody() *DeleteResourceDLinkResponseBody {
	return s.Body
}

func (s *DeleteResourceDLinkResponse) SetHeaders(v map[string]*string) *DeleteResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceDLinkResponse) SetStatusCode(v int32) *DeleteResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceDLinkResponse) SetBody(v *DeleteResourceDLinkResponseBody) *DeleteResourceDLinkResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceDLinkResponse) Validate() error {
	return dara.Validate(s)
}
