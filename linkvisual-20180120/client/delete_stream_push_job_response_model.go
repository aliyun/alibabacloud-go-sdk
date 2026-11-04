// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamPushJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamPushJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamPushJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamPushJobResponseBody) *DeleteStreamPushJobResponse
	GetBody() *DeleteStreamPushJobResponseBody
}

type DeleteStreamPushJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamPushJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamPushJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamPushJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamPushJobResponse) GetBody() *DeleteStreamPushJobResponseBody {
	return s.Body
}

func (s *DeleteStreamPushJobResponse) SetHeaders(v map[string]*string) *DeleteStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamPushJobResponse) SetStatusCode(v int32) *DeleteStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamPushJobResponse) SetBody(v *DeleteStreamPushJobResponseBody) *DeleteStreamPushJobResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamPushJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
