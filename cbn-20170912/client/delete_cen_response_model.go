// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenResponseBody) *DeleteCenResponse
	GetBody() *DeleteCenResponseBody
}

type DeleteCenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenResponse) GetBody() *DeleteCenResponseBody {
	return s.Body
}

func (s *DeleteCenResponse) SetHeaders(v map[string]*string) *DeleteCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenResponse) SetStatusCode(v int32) *DeleteCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenResponse) SetBody(v *DeleteCenResponseBody) *DeleteCenResponse {
	s.Body = v
	return s
}

func (s *DeleteCenResponse) Validate() error {
	return dara.Validate(s)
}
