// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DebugPlaybookResponseBody) *DebugPlaybookResponse
	GetBody() *DebugPlaybookResponseBody
}

type DebugPlaybookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugPlaybookResponse) GetBody() *DebugPlaybookResponseBody {
	return s.Body
}

func (s *DebugPlaybookResponse) SetHeaders(v map[string]*string) *DebugPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DebugPlaybookResponse) SetStatusCode(v int32) *DebugPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugPlaybookResponse) SetBody(v *DebugPlaybookResponseBody) *DebugPlaybookResponse {
	s.Body = v
	return s
}

func (s *DebugPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
