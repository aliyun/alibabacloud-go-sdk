// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DeletePlaybookResponseBody) *DeletePlaybookResponse
	GetBody() *DeletePlaybookResponseBody
}

type DeletePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DeletePlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePlaybookResponse) GetBody() *DeletePlaybookResponseBody {
	return s.Body
}

func (s *DeletePlaybookResponse) SetHeaders(v map[string]*string) *DeletePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DeletePlaybookResponse) SetStatusCode(v int32) *DeletePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlaybookResponse) SetBody(v *DeletePlaybookResponseBody) *DeletePlaybookResponse {
	s.Body = v
	return s
}

func (s *DeletePlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
