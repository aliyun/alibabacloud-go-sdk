// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiDestinationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiDestinationResponseBody) *DeleteApiDestinationResponse
	GetBody() *DeleteApiDestinationResponseBody
}

type DeleteApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiDestinationResponse) GetBody() *DeleteApiDestinationResponseBody {
	return s.Body
}

func (s *DeleteApiDestinationResponse) SetHeaders(v map[string]*string) *DeleteApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiDestinationResponse) SetStatusCode(v int32) *DeleteApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiDestinationResponse) SetBody(v *DeleteApiDestinationResponseBody) *DeleteApiDestinationResponse {
	s.Body = v
	return s
}

func (s *DeleteApiDestinationResponse) Validate() error {
	return dara.Validate(s)
}
