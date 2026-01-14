// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpSetsResponseBody) *ListIpSetsResponse
	GetBody() *ListIpSetsResponseBody
}

type ListIpSetsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpSetsResponse) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpSetsResponse) GetBody() *ListIpSetsResponseBody {
	return s.Body
}

func (s *ListIpSetsResponse) SetHeaders(v map[string]*string) *ListIpSetsResponse {
	s.Headers = v
	return s
}

func (s *ListIpSetsResponse) SetStatusCode(v int32) *ListIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpSetsResponse) SetBody(v *ListIpSetsResponseBody) *ListIpSetsResponse {
	s.Body = v
	return s
}

func (s *ListIpSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
