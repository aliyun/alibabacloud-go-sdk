// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPropertyValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPropertyValueResponse
	GetStatusCode() *int32
	SetBody(v *ListPropertyValueResponseBody) *ListPropertyValueResponse
	GetBody() *ListPropertyValueResponseBody
}

type ListPropertyValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPropertyValueResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPropertyValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPropertyValueResponse) GetBody() *ListPropertyValueResponseBody {
	return s.Body
}

func (s *ListPropertyValueResponse) SetHeaders(v map[string]*string) *ListPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyValueResponse) SetStatusCode(v int32) *ListPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyValueResponse) SetBody(v *ListPropertyValueResponseBody) *ListPropertyValueResponse {
	s.Body = v
	return s
}

func (s *ListPropertyValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
