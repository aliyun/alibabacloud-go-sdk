// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *VerifyPlaybookResponseBody) *VerifyPlaybookResponse
	GetBody() *VerifyPlaybookResponseBody
}

type VerifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyPlaybookResponse) GetBody() *VerifyPlaybookResponseBody {
	return s.Body
}

func (s *VerifyPlaybookResponse) SetHeaders(v map[string]*string) *VerifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *VerifyPlaybookResponse) SetStatusCode(v int32) *VerifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPlaybookResponse) SetBody(v *VerifyPlaybookResponseBody) *VerifyPlaybookResponse {
	s.Body = v
	return s
}

func (s *VerifyPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
