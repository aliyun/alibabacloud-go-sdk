// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFullNatEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFullNatEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFullNatEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListFullNatEntriesResponseBody) *ListFullNatEntriesResponse
	GetBody() *ListFullNatEntriesResponseBody
}

type ListFullNatEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFullNatEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFullNatEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFullNatEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListFullNatEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFullNatEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFullNatEntriesResponse) GetBody() *ListFullNatEntriesResponseBody {
	return s.Body
}

func (s *ListFullNatEntriesResponse) SetHeaders(v map[string]*string) *ListFullNatEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListFullNatEntriesResponse) SetStatusCode(v int32) *ListFullNatEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFullNatEntriesResponse) SetBody(v *ListFullNatEntriesResponseBody) *ListFullNatEntriesResponse {
	s.Body = v
	return s
}

func (s *ListFullNatEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
