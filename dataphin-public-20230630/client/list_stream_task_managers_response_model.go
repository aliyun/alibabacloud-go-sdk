// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamTaskManagersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamTaskManagersResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamTaskManagersResponseBody) *ListStreamTaskManagersResponse
	GetBody() *ListStreamTaskManagersResponseBody
}

type ListStreamTaskManagersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamTaskManagersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamTaskManagersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagersResponse) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamTaskManagersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamTaskManagersResponse) GetBody() *ListStreamTaskManagersResponseBody {
	return s.Body
}

func (s *ListStreamTaskManagersResponse) SetHeaders(v map[string]*string) *ListStreamTaskManagersResponse {
	s.Headers = v
	return s
}

func (s *ListStreamTaskManagersResponse) SetStatusCode(v int32) *ListStreamTaskManagersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamTaskManagersResponse) SetBody(v *ListStreamTaskManagersResponseBody) *ListStreamTaskManagersResponse {
	s.Body = v
	return s
}

func (s *ListStreamTaskManagersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
