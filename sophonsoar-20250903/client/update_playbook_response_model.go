// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePlaybookResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePlaybookResponseBody) *UpdatePlaybookResponse
	GetBody() *UpdatePlaybookResponseBody
}

type UpdatePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePlaybookResponse) GoString() string {
	return s.String()
}

func (s *UpdatePlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePlaybookResponse) GetBody() *UpdatePlaybookResponseBody {
	return s.Body
}

func (s *UpdatePlaybookResponse) SetHeaders(v map[string]*string) *UpdatePlaybookResponse {
	s.Headers = v
	return s
}

func (s *UpdatePlaybookResponse) SetStatusCode(v int32) *UpdatePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePlaybookResponse) SetBody(v *UpdatePlaybookResponseBody) *UpdatePlaybookResponse {
	s.Body = v
	return s
}

func (s *UpdatePlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
