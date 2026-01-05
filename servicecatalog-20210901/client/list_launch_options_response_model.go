// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaunchOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLaunchOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLaunchOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLaunchOptionsResponseBody) *ListLaunchOptionsResponse
	GetBody() *ListLaunchOptionsResponseBody
}

type ListLaunchOptionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLaunchOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLaunchOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLaunchOptionsResponse) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLaunchOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLaunchOptionsResponse) GetBody() *ListLaunchOptionsResponseBody {
	return s.Body
}

func (s *ListLaunchOptionsResponse) SetHeaders(v map[string]*string) *ListLaunchOptionsResponse {
	s.Headers = v
	return s
}

func (s *ListLaunchOptionsResponse) SetStatusCode(v int32) *ListLaunchOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLaunchOptionsResponse) SetBody(v *ListLaunchOptionsResponseBody) *ListLaunchOptionsResponse {
	s.Body = v
	return s
}

func (s *ListLaunchOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
