// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAbacAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAbacAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAbacAuthorizationResponseBody) *DeleteAbacAuthorizationResponse
	GetBody() *DeleteAbacAuthorizationResponseBody
}

type DeleteAbacAuthorizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAbacAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAbacAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAbacAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAbacAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAbacAuthorizationResponse) GetBody() *DeleteAbacAuthorizationResponseBody {
	return s.Body
}

func (s *DeleteAbacAuthorizationResponse) SetHeaders(v map[string]*string) *DeleteAbacAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAbacAuthorizationResponse) SetStatusCode(v int32) *DeleteAbacAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAbacAuthorizationResponse) SetBody(v *DeleteAbacAuthorizationResponseBody) *DeleteAbacAuthorizationResponse {
	s.Body = v
	return s
}

func (s *DeleteAbacAuthorizationResponse) Validate() error {
	return dara.Validate(s)
}
