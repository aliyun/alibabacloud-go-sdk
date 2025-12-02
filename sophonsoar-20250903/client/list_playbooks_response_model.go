// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaybooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlaybooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlaybooksResponse
	GetStatusCode() *int32
	SetBody(v *ListPlaybooksResponseBody) *ListPlaybooksResponse
	GetBody() *ListPlaybooksResponseBody
}

type ListPlaybooksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlaybooksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlaybooksResponse) GoString() string {
	return s.String()
}

func (s *ListPlaybooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlaybooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlaybooksResponse) GetBody() *ListPlaybooksResponseBody {
	return s.Body
}

func (s *ListPlaybooksResponse) SetHeaders(v map[string]*string) *ListPlaybooksResponse {
	s.Headers = v
	return s
}

func (s *ListPlaybooksResponse) SetStatusCode(v int32) *ListPlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlaybooksResponse) SetBody(v *ListPlaybooksResponseBody) *ListPlaybooksResponse {
	s.Body = v
	return s
}

func (s *ListPlaybooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
