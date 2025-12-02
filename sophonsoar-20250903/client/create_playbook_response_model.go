// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePlaybookResponse
	GetStatusCode() *int32
	SetBody(v *CreatePlaybookResponseBody) *CreatePlaybookResponse
	GetBody() *CreatePlaybookResponseBody
}

type CreatePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePlaybookResponse) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePlaybookResponse) GetBody() *CreatePlaybookResponseBody {
	return s.Body
}

func (s *CreatePlaybookResponse) SetHeaders(v map[string]*string) *CreatePlaybookResponse {
	s.Headers = v
	return s
}

func (s *CreatePlaybookResponse) SetStatusCode(v int32) *CreatePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlaybookResponse) SetBody(v *CreatePlaybookResponseBody) *CreatePlaybookResponse {
	s.Body = v
	return s
}

func (s *CreatePlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
