// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackImageResponse
	GetStatusCode() *int32
	SetBody(v *RollbackImageResponseBody) *RollbackImageResponse
	GetBody() *RollbackImageResponseBody
}

type RollbackImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackImageResponse) GoString() string {
	return s.String()
}

func (s *RollbackImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackImageResponse) GetBody() *RollbackImageResponseBody {
	return s.Body
}

func (s *RollbackImageResponse) SetHeaders(v map[string]*string) *RollbackImageResponse {
	s.Headers = v
	return s
}

func (s *RollbackImageResponse) SetStatusCode(v int32) *RollbackImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackImageResponse) SetBody(v *RollbackImageResponseBody) *RollbackImageResponse {
	s.Body = v
	return s
}

func (s *RollbackImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
