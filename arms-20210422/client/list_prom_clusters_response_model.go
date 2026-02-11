// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListPromClustersResponseBody) *ListPromClustersResponse
	GetBody() *ListPromClustersResponseBody
}

type ListPromClustersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromClustersResponse) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromClustersResponse) GetBody() *ListPromClustersResponseBody {
	return s.Body
}

func (s *ListPromClustersResponse) SetHeaders(v map[string]*string) *ListPromClustersResponse {
	s.Headers = v
	return s
}

func (s *ListPromClustersResponse) SetStatusCode(v int32) *ListPromClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromClustersResponse) SetBody(v *ListPromClustersResponseBody) *ListPromClustersResponse {
	s.Body = v
	return s
}

func (s *ListPromClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
