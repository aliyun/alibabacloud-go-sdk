// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResumeFunctionInstanceResponseBody) *ResumeFunctionInstanceResponse
	GetBody() *ResumeFunctionInstanceResponseBody
}

type ResumeFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeFunctionInstanceResponse) GetBody() *ResumeFunctionInstanceResponseBody {
	return s.Body
}

func (s *ResumeFunctionInstanceResponse) SetHeaders(v map[string]*string) *ResumeFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeFunctionInstanceResponse) SetStatusCode(v int32) *ResumeFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeFunctionInstanceResponse) SetBody(v *ResumeFunctionInstanceResponseBody) *ResumeFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *ResumeFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
