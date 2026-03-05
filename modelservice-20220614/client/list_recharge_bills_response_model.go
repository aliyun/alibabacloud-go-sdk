// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRechargeBillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRechargeBillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRechargeBillsResponse
	GetStatusCode() *int32
	SetBody(v *ListRechargeBillsResponseBody) *ListRechargeBillsResponse
	GetBody() *ListRechargeBillsResponseBody
}

type ListRechargeBillsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRechargeBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRechargeBillsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRechargeBillsResponse) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRechargeBillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRechargeBillsResponse) GetBody() *ListRechargeBillsResponseBody {
	return s.Body
}

func (s *ListRechargeBillsResponse) SetHeaders(v map[string]*string) *ListRechargeBillsResponse {
	s.Headers = v
	return s
}

func (s *ListRechargeBillsResponse) SetStatusCode(v int32) *ListRechargeBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRechargeBillsResponse) SetBody(v *ListRechargeBillsResponseBody) *ListRechargeBillsResponse {
	s.Body = v
	return s
}

func (s *ListRechargeBillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
