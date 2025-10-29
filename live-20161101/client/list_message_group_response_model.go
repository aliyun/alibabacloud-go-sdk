// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageGroupResponseBody) *ListMessageGroupResponse
	GetBody() *ListMessageGroupResponseBody
}

type ListMessageGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *ListMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageGroupResponse) GetBody() *ListMessageGroupResponseBody {
	return s.Body
}

func (s *ListMessageGroupResponse) SetHeaders(v map[string]*string) *ListMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *ListMessageGroupResponse) SetStatusCode(v int32) *ListMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageGroupResponse) SetBody(v *ListMessageGroupResponseBody) *ListMessageGroupResponse {
	s.Body = v
	return s
}

func (s *ListMessageGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
