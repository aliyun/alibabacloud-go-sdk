// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifySiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifySiteResponse
	GetStatusCode() *int32
	SetBody(v *VerifySiteResponseBody) *VerifySiteResponse
	GetBody() *VerifySiteResponseBody
}

type VerifySiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySiteResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifySiteResponse) GoString() string {
	return s.String()
}

func (s *VerifySiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifySiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifySiteResponse) GetBody() *VerifySiteResponseBody {
	return s.Body
}

func (s *VerifySiteResponse) SetHeaders(v map[string]*string) *VerifySiteResponse {
	s.Headers = v
	return s
}

func (s *VerifySiteResponse) SetStatusCode(v int32) *VerifySiteResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySiteResponse) SetBody(v *VerifySiteResponseBody) *VerifySiteResponse {
	s.Body = v
	return s
}

func (s *VerifySiteResponse) Validate() error {
	return dara.Validate(s)
}
