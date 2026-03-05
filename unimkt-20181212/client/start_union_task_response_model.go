// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUnionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartUnionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartUnionTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartUnionTaskResponseBody) *StartUnionTaskResponse
	GetBody() *StartUnionTaskResponseBody
}

type StartUnionTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartUnionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartUnionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartUnionTaskResponse) GoString() string {
	return s.String()
}

func (s *StartUnionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartUnionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartUnionTaskResponse) GetBody() *StartUnionTaskResponseBody {
	return s.Body
}

func (s *StartUnionTaskResponse) SetHeaders(v map[string]*string) *StartUnionTaskResponse {
	s.Headers = v
	return s
}

func (s *StartUnionTaskResponse) SetStatusCode(v int32) *StartUnionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartUnionTaskResponse) SetBody(v *StartUnionTaskResponseBody) *StartUnionTaskResponse {
	s.Body = v
	return s
}

func (s *StartUnionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
