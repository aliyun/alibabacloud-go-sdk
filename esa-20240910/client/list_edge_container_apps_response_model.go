// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeContainerAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeContainerAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeContainerAppsResponseBody) *ListEdgeContainerAppsResponse
	GetBody() *ListEdgeContainerAppsResponseBody
}

type ListEdgeContainerAppsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeContainerAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeContainerAppsResponse) GetBody() *ListEdgeContainerAppsResponseBody {
	return s.Body
}

func (s *ListEdgeContainerAppsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerAppsResponse) SetStatusCode(v int32) *ListEdgeContainerAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerAppsResponse) SetBody(v *ListEdgeContainerAppsResponseBody) *ListEdgeContainerAppsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeContainerAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
