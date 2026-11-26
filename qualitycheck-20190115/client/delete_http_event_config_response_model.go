// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpEventConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpEventConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpEventConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpEventConfigResponseBody) *DeleteHttpEventConfigResponse
	GetBody() *DeleteHttpEventConfigResponseBody
}

type DeleteHttpEventConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpEventConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpEventConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpEventConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpEventConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpEventConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpEventConfigResponse) GetBody() *DeleteHttpEventConfigResponseBody {
	return s.Body
}

func (s *DeleteHttpEventConfigResponse) SetHeaders(v map[string]*string) *DeleteHttpEventConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpEventConfigResponse) SetStatusCode(v int32) *DeleteHttpEventConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpEventConfigResponse) SetBody(v *DeleteHttpEventConfigResponseBody) *DeleteHttpEventConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpEventConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
