// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWordGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWordGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWordGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListWordGroupResponseBody) *ListWordGroupResponse
	GetBody() *ListWordGroupResponseBody
}

type ListWordGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWordGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWordGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWordGroupResponse) GoString() string {
	return s.String()
}

func (s *ListWordGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWordGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWordGroupResponse) GetBody() *ListWordGroupResponseBody {
	return s.Body
}

func (s *ListWordGroupResponse) SetHeaders(v map[string]*string) *ListWordGroupResponse {
	s.Headers = v
	return s
}

func (s *ListWordGroupResponse) SetStatusCode(v int32) *ListWordGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWordGroupResponse) SetBody(v *ListWordGroupResponseBody) *ListWordGroupResponse {
	s.Body = v
	return s
}

func (s *ListWordGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
