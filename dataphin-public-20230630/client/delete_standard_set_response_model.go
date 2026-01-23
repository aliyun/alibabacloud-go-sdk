// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardSetResponseBody) *DeleteStandardSetResponse
	GetBody() *DeleteStandardSetResponseBody
}

type DeleteStandardSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardSetResponse) GetBody() *DeleteStandardSetResponseBody {
	return s.Body
}

func (s *DeleteStandardSetResponse) SetHeaders(v map[string]*string) *DeleteStandardSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardSetResponse) SetStatusCode(v int32) *DeleteStandardSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardSetResponse) SetBody(v *DeleteStandardSetResponseBody) *DeleteStandardSetResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
