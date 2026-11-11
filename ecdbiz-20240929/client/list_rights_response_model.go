// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRightsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRightsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRightsResponse
	GetStatusCode() *int32
	SetBody(v *ListRightsResponseBody) *ListRightsResponse
	GetBody() *ListRightsResponseBody
}

type ListRightsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRightsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRightsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRightsResponse) GoString() string {
	return s.String()
}

func (s *ListRightsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRightsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRightsResponse) GetBody() *ListRightsResponseBody {
	return s.Body
}

func (s *ListRightsResponse) SetHeaders(v map[string]*string) *ListRightsResponse {
	s.Headers = v
	return s
}

func (s *ListRightsResponse) SetStatusCode(v int32) *ListRightsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRightsResponse) SetBody(v *ListRightsResponseBody) *ListRightsResponse {
	s.Body = v
	return s
}

func (s *ListRightsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
