// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokePushTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokePushTaskResponse
	GetStatusCode() *int32
	SetBody(v *RevokePushTaskResponseBody) *RevokePushTaskResponse
	GetBody() *RevokePushTaskResponseBody
}

type RevokePushTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokePushTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePushTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokePushTaskResponse) GoString() string {
	return s.String()
}

func (s *RevokePushTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokePushTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokePushTaskResponse) GetBody() *RevokePushTaskResponseBody {
	return s.Body
}

func (s *RevokePushTaskResponse) SetHeaders(v map[string]*string) *RevokePushTaskResponse {
	s.Headers = v
	return s
}

func (s *RevokePushTaskResponse) SetStatusCode(v int32) *RevokePushTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePushTaskResponse) SetBody(v *RevokePushTaskResponseBody) *RevokePushTaskResponse {
	s.Body = v
	return s
}

func (s *RevokePushTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
