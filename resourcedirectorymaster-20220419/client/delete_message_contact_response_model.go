// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageContactResponseBody) *DeleteMessageContactResponse
	GetBody() *DeleteMessageContactResponseBody
}

type DeleteMessageContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageContactResponse) GetBody() *DeleteMessageContactResponseBody {
	return s.Body
}

func (s *DeleteMessageContactResponse) SetHeaders(v map[string]*string) *DeleteMessageContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageContactResponse) SetStatusCode(v int32) *DeleteMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageContactResponse) SetBody(v *DeleteMessageContactResponseBody) *DeleteMessageContactResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
