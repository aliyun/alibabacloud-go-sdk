// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoDisposeEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoDisposeEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoDisposeEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoDisposeEntitiesResponseBody) *ListAutoDisposeEntitiesResponse
	GetBody() *ListAutoDisposeEntitiesResponseBody
}

type ListAutoDisposeEntitiesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoDisposeEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoDisposeEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoDisposeEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoDisposeEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoDisposeEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoDisposeEntitiesResponse) GetBody() *ListAutoDisposeEntitiesResponseBody {
	return s.Body
}

func (s *ListAutoDisposeEntitiesResponse) SetHeaders(v map[string]*string) *ListAutoDisposeEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoDisposeEntitiesResponse) SetStatusCode(v int32) *ListAutoDisposeEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponse) SetBody(v *ListAutoDisposeEntitiesResponseBody) *ListAutoDisposeEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListAutoDisposeEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
