// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamPushJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStreamPushJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStreamPushJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryStreamPushJobResponseBody) *QueryStreamPushJobResponse
	GetBody() *QueryStreamPushJobResponseBody
}

type QueryStreamPushJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamPushJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStreamPushJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStreamPushJobResponse) GetBody() *QueryStreamPushJobResponseBody {
	return s.Body
}

func (s *QueryStreamPushJobResponse) SetHeaders(v map[string]*string) *QueryStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamPushJobResponse) SetStatusCode(v int32) *QueryStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamPushJobResponse) SetBody(v *QueryStreamPushJobResponseBody) *QueryStreamPushJobResponse {
	s.Body = v
	return s
}

func (s *QueryStreamPushJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
