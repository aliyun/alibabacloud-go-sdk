// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizedDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizedDictResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizedDictResponseBody) *DeleteCustomizedDictResponse
	GetBody() *DeleteCustomizedDictResponseBody
}

type DeleteCustomizedDictResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizedDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizedDictResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedDictResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizedDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizedDictResponse) GetBody() *DeleteCustomizedDictResponseBody {
	return s.Body
}

func (s *DeleteCustomizedDictResponse) SetHeaders(v map[string]*string) *DeleteCustomizedDictResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizedDictResponse) SetStatusCode(v int32) *DeleteCustomizedDictResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizedDictResponse) SetBody(v *DeleteCustomizedDictResponseBody) *DeleteCustomizedDictResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizedDictResponse) Validate() error {
	return dara.Validate(s)
}
