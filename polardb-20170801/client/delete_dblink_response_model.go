// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBLinkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBLinkResponseBody) *DeleteDBLinkResponse
	GetBody() *DeleteDBLinkResponseBody
}

type DeleteDBLinkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBLinkResponse) GetBody() *DeleteDBLinkResponseBody {
	return s.Body
}

func (s *DeleteDBLinkResponse) SetHeaders(v map[string]*string) *DeleteDBLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBLinkResponse) SetStatusCode(v int32) *DeleteDBLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBLinkResponse) SetBody(v *DeleteDBLinkResponseBody) *DeleteDBLinkResponse {
	s.Body = v
	return s
}

func (s *DeleteDBLinkResponse) Validate() error {
	return dara.Validate(s)
}
