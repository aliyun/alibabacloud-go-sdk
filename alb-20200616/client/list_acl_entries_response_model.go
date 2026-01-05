// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAclEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAclEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAclEntriesResponseBody) *ListAclEntriesResponse
	GetBody() *ListAclEntriesResponseBody
}

type ListAclEntriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAclEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAclEntriesResponse) GetBody() *ListAclEntriesResponseBody {
	return s.Body
}

func (s *ListAclEntriesResponse) SetHeaders(v map[string]*string) *ListAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListAclEntriesResponse) SetStatusCode(v int32) *ListAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclEntriesResponse) SetBody(v *ListAclEntriesResponseBody) *ListAclEntriesResponse {
	s.Body = v
	return s
}

func (s *ListAclEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
