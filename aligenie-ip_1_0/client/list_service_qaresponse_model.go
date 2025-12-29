// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceQAResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceQAResponseBody) *ListServiceQAResponse
	GetBody() *ListServiceQAResponseBody
}

type ListServiceQAResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceQAResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAResponse) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceQAResponse) GetBody() *ListServiceQAResponseBody {
	return s.Body
}

func (s *ListServiceQAResponse) SetHeaders(v map[string]*string) *ListServiceQAResponse {
	s.Headers = v
	return s
}

func (s *ListServiceQAResponse) SetStatusCode(v int32) *ListServiceQAResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceQAResponse) SetBody(v *ListServiceQAResponseBody) *ListServiceQAResponse {
	s.Body = v
	return s
}

func (s *ListServiceQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
