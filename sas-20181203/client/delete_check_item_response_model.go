// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCheckItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCheckItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCheckItemResponseBody) *DeleteCheckItemResponse
	GetBody() *DeleteCheckItemResponseBody
}

type DeleteCheckItemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCheckItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteCheckItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCheckItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCheckItemResponse) GetBody() *DeleteCheckItemResponseBody {
	return s.Body
}

func (s *DeleteCheckItemResponse) SetHeaders(v map[string]*string) *DeleteCheckItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteCheckItemResponse) SetStatusCode(v int32) *DeleteCheckItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCheckItemResponse) SetBody(v *DeleteCheckItemResponseBody) *DeleteCheckItemResponse {
	s.Body = v
	return s
}

func (s *DeleteCheckItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
