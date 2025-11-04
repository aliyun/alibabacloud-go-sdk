// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduledTaskResponseBody) *DeleteScheduledTaskResponse
	GetBody() *DeleteScheduledTaskResponseBody
}

type DeleteScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduledTaskResponse) GetBody() *DeleteScheduledTaskResponseBody {
	return s.Body
}

func (s *DeleteScheduledTaskResponse) SetHeaders(v map[string]*string) *DeleteScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledTaskResponse) SetStatusCode(v int32) *DeleteScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledTaskResponse) SetBody(v *DeleteScheduledTaskResponseBody) *DeleteScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
