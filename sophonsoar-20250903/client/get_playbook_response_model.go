// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *GetPlaybookResponseBody) *GetPlaybookResponse
	GetBody() *GetPlaybookResponseBody
}

type GetPlaybookResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlaybookResponse) GoString() string {
	return s.String()
}

func (s *GetPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlaybookResponse) GetBody() *GetPlaybookResponseBody {
	return s.Body
}

func (s *GetPlaybookResponse) SetHeaders(v map[string]*string) *GetPlaybookResponse {
	s.Headers = v
	return s
}

func (s *GetPlaybookResponse) SetStatusCode(v int32) *GetPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlaybookResponse) SetBody(v *GetPlaybookResponseBody) *GetPlaybookResponse {
	s.Body = v
	return s
}

func (s *GetPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
