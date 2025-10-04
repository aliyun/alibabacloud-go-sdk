// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWuyingServerResponseBody) *DeleteWuyingServerResponse
	GetBody() *DeleteWuyingServerResponseBody
}

type DeleteWuyingServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWuyingServerResponse) GetBody() *DeleteWuyingServerResponseBody {
	return s.Body
}

func (s *DeleteWuyingServerResponse) SetHeaders(v map[string]*string) *DeleteWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteWuyingServerResponse) SetStatusCode(v int32) *DeleteWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWuyingServerResponse) SetBody(v *DeleteWuyingServerResponseBody) *DeleteWuyingServerResponse {
	s.Body = v
	return s
}

func (s *DeleteWuyingServerResponse) Validate() error {
	return dara.Validate(s)
}
