// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmountInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnmountInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnmountInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *UnmountInstanceSDGResponseBody) *UnmountInstanceSDGResponse
	GetBody() *UnmountInstanceSDGResponseBody
}

type UnmountInstanceSDGResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnmountInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnmountInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnmountInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnmountInstanceSDGResponse) GetBody() *UnmountInstanceSDGResponseBody {
	return s.Body
}

func (s *UnmountInstanceSDGResponse) SetHeaders(v map[string]*string) *UnmountInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *UnmountInstanceSDGResponse) SetStatusCode(v int32) *UnmountInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *UnmountInstanceSDGResponse) SetBody(v *UnmountInstanceSDGResponseBody) *UnmountInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *UnmountInstanceSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
