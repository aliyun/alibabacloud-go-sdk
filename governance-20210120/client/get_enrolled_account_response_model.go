// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnrolledAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnrolledAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnrolledAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetEnrolledAccountResponseBody) *GetEnrolledAccountResponse
	GetBody() *GetEnrolledAccountResponseBody
}

type GetEnrolledAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnrolledAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnrolledAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountResponse) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnrolledAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnrolledAccountResponse) GetBody() *GetEnrolledAccountResponseBody {
	return s.Body
}

func (s *GetEnrolledAccountResponse) SetHeaders(v map[string]*string) *GetEnrolledAccountResponse {
	s.Headers = v
	return s
}

func (s *GetEnrolledAccountResponse) SetStatusCode(v int32) *GetEnrolledAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnrolledAccountResponse) SetBody(v *GetEnrolledAccountResponseBody) *GetEnrolledAccountResponse {
	s.Body = v
	return s
}

func (s *GetEnrolledAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
