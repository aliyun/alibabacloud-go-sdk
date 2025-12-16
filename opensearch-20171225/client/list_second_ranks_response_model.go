// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecondRanksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecondRanksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecondRanksResponse
	GetStatusCode() *int32
	SetBody(v *ListSecondRanksResponseBody) *ListSecondRanksResponse
	GetBody() *ListSecondRanksResponseBody
}

type ListSecondRanksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecondRanksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecondRanksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecondRanksResponse) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecondRanksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecondRanksResponse) GetBody() *ListSecondRanksResponseBody {
	return s.Body
}

func (s *ListSecondRanksResponse) SetHeaders(v map[string]*string) *ListSecondRanksResponse {
	s.Headers = v
	return s
}

func (s *ListSecondRanksResponse) SetStatusCode(v int32) *ListSecondRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecondRanksResponse) SetBody(v *ListSecondRanksResponseBody) *ListSecondRanksResponse {
	s.Body = v
	return s
}

func (s *ListSecondRanksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
