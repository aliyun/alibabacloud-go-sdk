// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdInsertionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdInsertionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdInsertionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdInsertionResponseBody) *DeleteAdInsertionResponse
	GetBody() *DeleteAdInsertionResponseBody
}

type DeleteAdInsertionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdInsertionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdInsertionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdInsertionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdInsertionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdInsertionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdInsertionResponse) GetBody() *DeleteAdInsertionResponseBody {
	return s.Body
}

func (s *DeleteAdInsertionResponse) SetHeaders(v map[string]*string) *DeleteAdInsertionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdInsertionResponse) SetStatusCode(v int32) *DeleteAdInsertionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdInsertionResponse) SetBody(v *DeleteAdInsertionResponseBody) *DeleteAdInsertionResponse {
	s.Body = v
	return s
}

func (s *DeleteAdInsertionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
