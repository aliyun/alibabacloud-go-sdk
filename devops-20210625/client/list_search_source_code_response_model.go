// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchSourceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchSourceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchSourceCodeResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchSourceCodeResponseBody) *ListSearchSourceCodeResponse
	GetBody() *ListSearchSourceCodeResponseBody
}

type ListSearchSourceCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchSourceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchSourceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeResponse) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchSourceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchSourceCodeResponse) GetBody() *ListSearchSourceCodeResponseBody {
	return s.Body
}

func (s *ListSearchSourceCodeResponse) SetHeaders(v map[string]*string) *ListSearchSourceCodeResponse {
	s.Headers = v
	return s
}

func (s *ListSearchSourceCodeResponse) SetStatusCode(v int32) *ListSearchSourceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchSourceCodeResponse) SetBody(v *ListSearchSourceCodeResponseBody) *ListSearchSourceCodeResponse {
	s.Body = v
	return s
}

func (s *ListSearchSourceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
