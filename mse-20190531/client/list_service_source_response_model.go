// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceSourceResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceSourceResponseBody) *ListServiceSourceResponse
	GetBody() *ListServiceSourceResponseBody
}

type ListServiceSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *ListServiceSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceSourceResponse) GetBody() *ListServiceSourceResponseBody {
	return s.Body
}

func (s *ListServiceSourceResponse) SetHeaders(v map[string]*string) *ListServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *ListServiceSourceResponse) SetStatusCode(v int32) *ListServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceSourceResponse) SetBody(v *ListServiceSourceResponseBody) *ListServiceSourceResponse {
	s.Body = v
	return s
}

func (s *ListServiceSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
