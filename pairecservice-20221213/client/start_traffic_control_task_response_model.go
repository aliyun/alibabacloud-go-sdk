// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartTrafficControlTaskResponseBody) *StartTrafficControlTaskResponse
	GetBody() *StartTrafficControlTaskResponseBody
}

type StartTrafficControlTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTrafficControlTaskResponse) GetBody() *StartTrafficControlTaskResponseBody {
	return s.Body
}

func (s *StartTrafficControlTaskResponse) SetHeaders(v map[string]*string) *StartTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTrafficControlTaskResponse) SetStatusCode(v int32) *StartTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTrafficControlTaskResponse) SetBody(v *StartTrafficControlTaskResponseBody) *StartTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *StartTrafficControlTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
