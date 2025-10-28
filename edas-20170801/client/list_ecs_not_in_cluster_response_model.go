// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsNotInClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEcsNotInClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEcsNotInClusterResponse
	GetStatusCode() *int32
	SetBody(v *ListEcsNotInClusterResponseBody) *ListEcsNotInClusterResponse
	GetBody() *ListEcsNotInClusterResponseBody
}

type ListEcsNotInClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEcsNotInClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEcsNotInClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEcsNotInClusterResponse) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEcsNotInClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEcsNotInClusterResponse) GetBody() *ListEcsNotInClusterResponseBody {
	return s.Body
}

func (s *ListEcsNotInClusterResponse) SetHeaders(v map[string]*string) *ListEcsNotInClusterResponse {
	s.Headers = v
	return s
}

func (s *ListEcsNotInClusterResponse) SetStatusCode(v int32) *ListEcsNotInClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsNotInClusterResponse) SetBody(v *ListEcsNotInClusterResponseBody) *ListEcsNotInClusterResponse {
	s.Body = v
	return s
}

func (s *ListEcsNotInClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
