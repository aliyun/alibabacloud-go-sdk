// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFirstRanksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFirstRanksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFirstRanksResponse
	GetStatusCode() *int32
	SetBody(v *ListFirstRanksResponseBody) *ListFirstRanksResponse
	GetBody() *ListFirstRanksResponseBody
}

type ListFirstRanksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFirstRanksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFirstRanksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFirstRanksResponse) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFirstRanksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFirstRanksResponse) GetBody() *ListFirstRanksResponseBody {
	return s.Body
}

func (s *ListFirstRanksResponse) SetHeaders(v map[string]*string) *ListFirstRanksResponse {
	s.Headers = v
	return s
}

func (s *ListFirstRanksResponse) SetStatusCode(v int32) *ListFirstRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFirstRanksResponse) SetBody(v *ListFirstRanksResponseBody) *ListFirstRanksResponse {
	s.Body = v
	return s
}

func (s *ListFirstRanksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
