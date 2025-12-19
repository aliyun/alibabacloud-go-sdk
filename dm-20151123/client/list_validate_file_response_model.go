// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValidateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListValidateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListValidateFileResponse
	GetStatusCode() *int32
	SetBody(v *ListValidateFileResponseBody) *ListValidateFileResponse
	GetBody() *ListValidateFileResponseBody
}

type ListValidateFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListValidateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListValidateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListValidateFileResponse) GoString() string {
	return s.String()
}

func (s *ListValidateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListValidateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListValidateFileResponse) GetBody() *ListValidateFileResponseBody {
	return s.Body
}

func (s *ListValidateFileResponse) SetHeaders(v map[string]*string) *ListValidateFileResponse {
	s.Headers = v
	return s
}

func (s *ListValidateFileResponse) SetStatusCode(v int32) *ListValidateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListValidateFileResponse) SetBody(v *ListValidateFileResponseBody) *ListValidateFileResponse {
	s.Body = v
	return s
}

func (s *ListValidateFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
