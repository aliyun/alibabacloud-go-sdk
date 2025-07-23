// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePCACertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePCACertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePCACertResponse
	GetStatusCode() *int32
	SetBody(v *DeletePCACertResponseBody) *DeletePCACertResponse
	GetBody() *DeletePCACertResponseBody
}

type DeletePCACertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePCACertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePCACertResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePCACertResponse) GoString() string {
	return s.String()
}

func (s *DeletePCACertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePCACertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePCACertResponse) GetBody() *DeletePCACertResponseBody {
	return s.Body
}

func (s *DeletePCACertResponse) SetHeaders(v map[string]*string) *DeletePCACertResponse {
	s.Headers = v
	return s
}

func (s *DeletePCACertResponse) SetStatusCode(v int32) *DeletePCACertResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePCACertResponse) SetBody(v *DeletePCACertResponseBody) *DeletePCACertResponse {
	s.Body = v
	return s
}

func (s *DeletePCACertResponse) Validate() error {
	return dara.Validate(s)
}
