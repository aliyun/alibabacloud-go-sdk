// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficControlTaskResponseBody) *DeleteTrafficControlTaskResponse
	GetBody() *DeleteTrafficControlTaskResponseBody
}

type DeleteTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficControlTaskResponse) GetBody() *DeleteTrafficControlTaskResponseBody {
	return s.Body
}

func (s *DeleteTrafficControlTaskResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlTaskResponse) SetStatusCode(v int32) *DeleteTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlTaskResponse) SetBody(v *DeleteTrafficControlTaskResponseBody) *DeleteTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficControlTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
