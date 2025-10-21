// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelInstanceResponseBody) *DeleteModelInstanceResponse
	GetBody() *DeleteModelInstanceResponseBody
}

type DeleteModelInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelInstanceResponse) GetBody() *DeleteModelInstanceResponseBody {
	return s.Body
}

func (s *DeleteModelInstanceResponse) SetHeaders(v map[string]*string) *DeleteModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelInstanceResponse) SetStatusCode(v int32) *DeleteModelInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelInstanceResponse) SetBody(v *DeleteModelInstanceResponseBody) *DeleteModelInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteModelInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
