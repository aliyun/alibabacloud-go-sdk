// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCollectorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCollectorResponseBody) *DeleteCollectorResponse
	GetBody() *DeleteCollectorResponseBody
}

type DeleteCollectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCollectorResponse) GetBody() *DeleteCollectorResponseBody {
	return s.Body
}

func (s *DeleteCollectorResponse) SetHeaders(v map[string]*string) *DeleteCollectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectorResponse) SetStatusCode(v int32) *DeleteCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectorResponse) SetBody(v *DeleteCollectorResponseBody) *DeleteCollectorResponse {
	s.Body = v
	return s
}

func (s *DeleteCollectorResponse) Validate() error {
	return dara.Validate(s)
}
