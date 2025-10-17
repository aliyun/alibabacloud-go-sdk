// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocOcrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocOcrResponse
	GetStatusCode() *int32
	SetBody(v *DocOcrResponseBody) *DocOcrResponse
	GetBody() *DocOcrResponseBody
}

type DocOcrResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrResponse) String() string {
	return dara.Prettify(s)
}

func (s DocOcrResponse) GoString() string {
	return s.String()
}

func (s *DocOcrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocOcrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocOcrResponse) GetBody() *DocOcrResponseBody {
	return s.Body
}

func (s *DocOcrResponse) SetHeaders(v map[string]*string) *DocOcrResponse {
	s.Headers = v
	return s
}

func (s *DocOcrResponse) SetStatusCode(v int32) *DocOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *DocOcrResponse) SetBody(v *DocOcrResponseBody) *DocOcrResponse {
	s.Body = v
	return s
}

func (s *DocOcrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
