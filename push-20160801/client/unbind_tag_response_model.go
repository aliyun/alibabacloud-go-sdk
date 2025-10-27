// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindTagResponse
	GetStatusCode() *int32
	SetBody(v *UnbindTagResponseBody) *UnbindTagResponse
	GetBody() *UnbindTagResponseBody
}

type UnbindTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindTagResponse) GoString() string {
	return s.String()
}

func (s *UnbindTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindTagResponse) GetBody() *UnbindTagResponseBody {
	return s.Body
}

func (s *UnbindTagResponse) SetHeaders(v map[string]*string) *UnbindTagResponse {
	s.Headers = v
	return s
}

func (s *UnbindTagResponse) SetStatusCode(v int32) *UnbindTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindTagResponse) SetBody(v *UnbindTagResponseBody) *UnbindTagResponse {
	s.Body = v
	return s
}

func (s *UnbindTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
