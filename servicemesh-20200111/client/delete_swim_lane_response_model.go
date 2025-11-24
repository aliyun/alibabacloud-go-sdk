// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSwimLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSwimLaneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSwimLaneResponseBody) *DeleteSwimLaneResponse
	GetBody() *DeleteSwimLaneResponseBody
}

type DeleteSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSwimLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSwimLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSwimLaneResponse) GetBody() *DeleteSwimLaneResponseBody {
	return s.Body
}

func (s *DeleteSwimLaneResponse) SetHeaders(v map[string]*string) *DeleteSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimLaneResponse) SetStatusCode(v int32) *DeleteSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimLaneResponse) SetBody(v *DeleteSwimLaneResponseBody) *DeleteSwimLaneResponse {
	s.Body = v
	return s
}

func (s *DeleteSwimLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
