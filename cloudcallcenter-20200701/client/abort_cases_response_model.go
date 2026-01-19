// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortCasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortCasesResponse
	GetStatusCode() *int32
	SetBody(v *AbortCasesResponseBody) *AbortCasesResponse
	GetBody() *AbortCasesResponseBody
}

type AbortCasesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortCasesResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortCasesResponse) GoString() string {
	return s.String()
}

func (s *AbortCasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortCasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortCasesResponse) GetBody() *AbortCasesResponseBody {
	return s.Body
}

func (s *AbortCasesResponse) SetHeaders(v map[string]*string) *AbortCasesResponse {
	s.Headers = v
	return s
}

func (s *AbortCasesResponse) SetStatusCode(v int32) *AbortCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortCasesResponse) SetBody(v *AbortCasesResponseBody) *AbortCasesResponse {
	s.Body = v
	return s
}

func (s *AbortCasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
