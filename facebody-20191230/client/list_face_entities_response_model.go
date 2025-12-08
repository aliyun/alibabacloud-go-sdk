// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFaceEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFaceEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListFaceEntitiesResponseBody) *ListFaceEntitiesResponse
	GetBody() *ListFaceEntitiesResponseBody
}

type ListFaceEntitiesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFaceEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFaceEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFaceEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFaceEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFaceEntitiesResponse) GetBody() *ListFaceEntitiesResponseBody {
	return s.Body
}

func (s *ListFaceEntitiesResponse) SetHeaders(v map[string]*string) *ListFaceEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListFaceEntitiesResponse) SetStatusCode(v int32) *ListFaceEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaceEntitiesResponse) SetBody(v *ListFaceEntitiesResponseBody) *ListFaceEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListFaceEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
