// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *StartTensorboardResponseBody) *StartTensorboardResponse
	GetBody() *StartTensorboardResponseBody
}

type StartTensorboardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTensorboardResponse) GetBody() *StartTensorboardResponseBody {
	return s.Body
}

func (s *StartTensorboardResponse) SetHeaders(v map[string]*string) *StartTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StartTensorboardResponse) SetStatusCode(v int32) *StartTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTensorboardResponse) SetBody(v *StartTensorboardResponseBody) *StartTensorboardResponse {
	s.Body = v
	return s
}

func (s *StartTensorboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
