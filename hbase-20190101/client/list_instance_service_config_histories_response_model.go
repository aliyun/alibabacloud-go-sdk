// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceServiceConfigHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceServiceConfigHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceServiceConfigHistoriesResponseBody) *ListInstanceServiceConfigHistoriesResponse
	GetBody() *ListInstanceServiceConfigHistoriesResponseBody
}

type ListInstanceServiceConfigHistoriesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceServiceConfigHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceServiceConfigHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceServiceConfigHistoriesResponse) GetBody() *ListInstanceServiceConfigHistoriesResponseBody {
	return s.Body
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetStatusCode(v int32) *ListInstanceServiceConfigHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetBody(v *ListInstanceServiceConfigHistoriesResponseBody) *ListInstanceServiceConfigHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
