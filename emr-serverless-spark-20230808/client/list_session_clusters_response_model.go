// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionClustersResponseBody) *ListSessionClustersResponse
	GetBody() *ListSessionClustersResponseBody
}

type ListSessionClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponse) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionClustersResponse) GetBody() *ListSessionClustersResponseBody {
	return s.Body
}

func (s *ListSessionClustersResponse) SetHeaders(v map[string]*string) *ListSessionClustersResponse {
	s.Headers = v
	return s
}

func (s *ListSessionClustersResponse) SetStatusCode(v int32) *ListSessionClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionClustersResponse) SetBody(v *ListSessionClustersResponseBody) *ListSessionClustersResponse {
	s.Body = v
	return s
}

func (s *ListSessionClustersResponse) Validate() error {
	return dara.Validate(s)
}
