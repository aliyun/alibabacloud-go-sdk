// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *UnbindSaseUserTagResponseBody) *UnbindSaseUserTagResponse
	GetBody() *UnbindSaseUserTagResponseBody
}

type UnbindSaseUserTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *UnbindSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSaseUserTagResponse) GetBody() *UnbindSaseUserTagResponseBody {
	return s.Body
}

func (s *UnbindSaseUserTagResponse) SetHeaders(v map[string]*string) *UnbindSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *UnbindSaseUserTagResponse) SetStatusCode(v int32) *UnbindSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSaseUserTagResponse) SetBody(v *UnbindSaseUserTagResponseBody) *UnbindSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *UnbindSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
