// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABTestGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABTestGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListABTestGroupsResponseBody) *ListABTestGroupsResponse
	GetBody() *ListABTestGroupsResponseBody
}

type ListABTestGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABTestGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABTestGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABTestGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABTestGroupsResponse) GetBody() *ListABTestGroupsResponseBody {
	return s.Body
}

func (s *ListABTestGroupsResponse) SetHeaders(v map[string]*string) *ListABTestGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestGroupsResponse) SetStatusCode(v int32) *ListABTestGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestGroupsResponse) SetBody(v *ListABTestGroupsResponseBody) *ListABTestGroupsResponse {
	s.Body = v
	return s
}

func (s *ListABTestGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
