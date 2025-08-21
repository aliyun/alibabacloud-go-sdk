// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectedClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConnectedClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConnectedClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListConnectedClustersResponseBody) *ListConnectedClustersResponse
	GetBody() *ListConnectedClustersResponseBody
}

type ListConnectedClustersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectedClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConnectedClustersResponse) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConnectedClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConnectedClustersResponse) GetBody() *ListConnectedClustersResponseBody {
	return s.Body
}

func (s *ListConnectedClustersResponse) SetHeaders(v map[string]*string) *ListConnectedClustersResponse {
	s.Headers = v
	return s
}

func (s *ListConnectedClustersResponse) SetStatusCode(v int32) *ListConnectedClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectedClustersResponse) SetBody(v *ListConnectedClustersResponseBody) *ListConnectedClustersResponse {
	s.Body = v
	return s
}

func (s *ListConnectedClustersResponse) Validate() error {
	return dara.Validate(s)
}
