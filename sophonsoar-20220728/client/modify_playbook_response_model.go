// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPlaybookResponseBody) *ModifyPlaybookResponse
	GetBody() *ModifyPlaybookResponseBody
}

type ModifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPlaybookResponse) GetBody() *ModifyPlaybookResponseBody {
	return s.Body
}

func (s *ModifyPlaybookResponse) SetHeaders(v map[string]*string) *ModifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookResponse) SetStatusCode(v int32) *ModifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookResponse) SetBody(v *ModifyPlaybookResponseBody) *ModifyPlaybookResponse {
	s.Body = v
	return s
}

func (s *ModifyPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
