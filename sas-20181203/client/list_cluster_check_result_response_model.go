// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterCheckResultResponseBody) *ListClusterCheckResultResponse
	GetBody() *ListClusterCheckResultResponseBody
}

type ListClusterCheckResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultResponse) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterCheckResultResponse) GetBody() *ListClusterCheckResultResponseBody {
	return s.Body
}

func (s *ListClusterCheckResultResponse) SetHeaders(v map[string]*string) *ListClusterCheckResultResponse {
	s.Headers = v
	return s
}

func (s *ListClusterCheckResultResponse) SetStatusCode(v int32) *ListClusterCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterCheckResultResponse) SetBody(v *ListClusterCheckResultResponseBody) *ListClusterCheckResultResponse {
	s.Body = v
	return s
}

func (s *ListClusterCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
