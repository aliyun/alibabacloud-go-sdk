// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamPushJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamPushJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamPushJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamPushJobResponseBody) *CreateStreamPushJobResponse
	GetBody() *CreateStreamPushJobResponseBody
}

type CreateStreamPushJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamPushJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamPushJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamPushJobResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamPushJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamPushJobResponse) GetBody() *CreateStreamPushJobResponseBody {
	return s.Body
}

func (s *CreateStreamPushJobResponse) SetHeaders(v map[string]*string) *CreateStreamPushJobResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamPushJobResponse) SetStatusCode(v int32) *CreateStreamPushJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamPushJobResponse) SetBody(v *CreateStreamPushJobResponseBody) *CreateStreamPushJobResponse {
	s.Body = v
	return s
}

func (s *CreateStreamPushJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
